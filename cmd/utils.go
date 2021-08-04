package cmd

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/decadevvv/miniwf/pkg/utils"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func FlagAdd(cmd *cobra.Command, conf interface{}) {
	flagSet, err := gpflag.Parse(conf)
	utils.PanicOnError("failed to parse flagset from struct", err)
	cmd.Flags().AddFlagSet(flagSet)
}

func FlagParse(cmd *cobra.Command, conf interface{}) error {
	v := viper.New()
	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return fmt.Errorf("failed to bind pflags to viper: %w", err)
	}
	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("failed to unmarshal pflags: %w", err)
	}
	if err := utils.ValidateStruct(conf); err != nil {
		return fmt.Errorf("failed to validate pflags unmarshal result: %w", err)
	}
	dumpStr := spew.Sdump(conf)
	cmd.Println("----------------START PRE EXECUTE-----------------")
	cmd.Printf("flag parse result: \n%s", dumpStr)
	cmd.Println("----------------FINISH PRE EXECUTE----------------")
	return nil
}
