package cmd

import (
	"fmt"

	"github.com/decadevvv/miniwf/pkg/engine"
	"github.com/spf13/cobra"
)

func init() {
	SetupActionCommands(rootCmd)
}

func SetupActionCommands(parent *cobra.Command) {
	actionCmd := &cobra.Command{
		Use:              "action",
		Short:            "manage actions",
		Long:             "manage actions",
		TraverseChildren: true,
	}

	initDoc(actionCmd)

	parent.AddCommand(actionCmd)
}

func initDoc(parent *cobra.Command) {
	type ActionDocConfiguration struct {
		Name string `flag:"name" desc:"name of the action to lookup documentation" validate:"required"`
	}

	defaultConf := ActionDocConfiguration{
		Name: "",
	}

	cmd := &cobra.Command{
		Use:     "doc",
		Short:   "action document",
		Long:    "lookup and print documentation of an action",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := FlagParse(cmd, &defaultConf); err != nil {
				return err
			}
			action, ok := engine.DefaultEngine.Registry.GetAction(defaultConf.Name)
			if !ok {
				return fmt.Errorf("cannot find action %s", defaultConf.Name)
			}
			fmt.Printf("Documentation for action %s: \n%s\n", defaultConf.Name, action.Doc)
			return nil
		},
	}

	FlagAdd(cmd, &defaultConf)

	parent.AddCommand(cmd)
}
