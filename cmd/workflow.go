package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/decadevvv/miniwf/pkg/core"
	"github.com/decadevvv/miniwf/pkg/engine"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

func init() {
	SetupWorkflowCommands(rootCmd)
}

func SetupWorkflowCommands(parent *cobra.Command) {
	workflowCmd := &cobra.Command{
		Use:              "workflow",
		Short:            "manage workflows",
		Long:             "manage workflows",
		TraverseChildren: true,
	}

	initRun(workflowCmd)

	parent.AddCommand(workflowCmd)
}

func initRun(parent *cobra.Command) {
	type WorkflowRunConfiguration struct {
		Definition string `flag:"definition" desc:"path to the workflow definition file" validate:"required,file"`
		Debug      bool   `flag:"debug" desc:"whether or not to run workflow engine in debug mode"`
	}

	defaultConf := WorkflowRunConfiguration{
		Definition: "",
		Debug:      false,
	}

	cmd := &cobra.Command{
		Use:     "run",
		Short:   "run workflow",
		Long:    "run workflow",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := FlagParse(cmd, &defaultConf); err != nil {
				return err
			}
			yamlFile, err := ioutil.ReadFile(defaultConf.Definition)
			if err != nil {
				return fmt.Errorf("failed to read workflow definition file %s: %w", defaultConf.Definition, err)
			}
			workflow := core.Workflow{}
			err = yaml.Unmarshal(yamlFile, &workflow)
			if err != nil {
				return fmt.Errorf("failed to unmarshal workflow definition into workflow struct: %w", err)
			}
			engine.DefaultEngine.SetDebug(defaultConf.Debug)
			err = engine.DefaultEngine.RunWorkflow(workflow)
			if err != nil {
				return fmt.Errorf("failed to run workflow: %w", err)
			}
			return nil
		},
	}

	FlagAdd(cmd, &defaultConf)

	parent.AddCommand(cmd)
}
