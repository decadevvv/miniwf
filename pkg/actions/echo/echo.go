package echo

import (
	"fmt"

	"github.com/decadevvv/miniwf/pkg/core"
)

type EchoActionConf struct {
	Message string `yaml:"message" validate:"required"`
}

var EchoAction = core.Action{
	Name: "echo",
	Doc:  "echo message",
	DefaultConf: EchoActionConf{
		Message: "hello world",
	},
	Run: func(conf interface{}) (interface{}, error) {
		c, ok := conf.(EchoActionConf)
		if !ok {
			return "", fmt.Errorf("conf type is not correct")
		}
		fmt.Println(c.Message)
		return c.Message, nil
	},
}
