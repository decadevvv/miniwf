package engine

import (
	"fmt"

	"github.com/decadevvv/miniwf/pkg/actions/echo"
	"github.com/decadevvv/miniwf/pkg/core"
)

type Registry struct {
	actionMap map[string]core.Action
}

var DefaultRegistry = Registry{
	actionMap: map[string]core.Action{},
}

func init() {
	DefaultRegistry.Register(echo.NewEchoAction())
}

func (r *Registry) Register(action core.Action) {
	_, ok := r.actionMap[action.Name()]
	if ok {
		panic(fmt.Errorf("action %s already exists", action.Name()))
	}
	r.actionMap[action.Name()] = action
}

func (r *Registry) GetAction(name string) (core.Action, bool) {
	action, ok := r.actionMap[name]
	return action, ok
}
