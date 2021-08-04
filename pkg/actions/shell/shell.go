package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/decadevvv/miniwf/pkg/core"
)

type ShellActionConf struct {
	Command string `yaml:"command" validate:"required"`
}

var ShellAction = core.Action{
	Name: "shell",
	Doc:  "run shell commands",
	DefaultConf: ShellActionConf{
		Command: "echo \"hello world\"",
	},
	Run: func(conf interface{}) (interface{}, error) {
		c, ok := conf.(ShellActionConf)
		if !ok {
			return "", fmt.Errorf("conf type is not correct")
		}
		args := strings.Split(c.Command, " ")
		eCmd := exec.Command(args[0], args[1:]...)
		stdout := bytes.NewBuffer([]byte(""))
		stderr := bytes.NewBuffer([]byte(""))
		eCmd.Stdout = stdout
		eCmd.Stderr = stderr
		err := eCmd.Run()
		if err != nil {
			return nil, err
		}
		fmt.Printf("command: %s\n", eCmd.String())
		fmt.Printf("stdout: %s\n", stdout.String())
		fmt.Printf("stderr: %s\n", stderr.String())
		return map[string]string{
			"stdout": stdout.String(),
			"stderr": stderr.String(),
		}, nil
	},
}
