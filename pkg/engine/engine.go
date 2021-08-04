package engine

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/decadevvv/miniwf/pkg/core"
	"github.com/decadevvv/miniwf/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type Engine struct {
	Debug      bool
	Registry   Registry
	ContextMap map[string]interface{}
}

var DefaultEngine = Engine{
	Debug:      true,
	Registry:   DefaultRegistry,
	ContextMap: map[string]interface{}{},
}

func (e *Engine) SetDebug(debug bool) {
	e.Debug = debug
}

func (e *Engine) RunWorkflow(w core.Workflow) error {
	if e.Debug {
		log.Infof("Registered actions: %#v", e.Registry.ListRegisteredActions())
	}
	for index, step := range w.Steps {
		log.Infof("STEP %d: %s", index, step.Name)
		action, ok := e.Registry.GetAction(step.Action)
		if !ok {
			return fmt.Errorf("cannot find action %s", step.Action)
		}
		conf := action.DefaultConf
		err := utils.UnmarshalMapIntoStructWithTemplate(step.Conf, &conf, e.ContextMap, e.Debug)
		if err != nil {
			return fmt.Errorf("failed to unmarshal map into struct: %w", err)
		}
		output, err := action.Run(conf)
		if err != nil {
			return fmt.Errorf("failecd to run action: %w", err)
		}
		_, ok = e.ContextMap[step.Output]
		if ok {
			return fmt.Errorf("output name %s is already occupied in context map", step.Output)
		}
		e.ContextMap[step.Output] = output
	}
	if e.Debug {
		fmt.Println("FINAL CONTEXT MAP:")
		spew.Dump(e.ContextMap)
	}
	return nil
}
