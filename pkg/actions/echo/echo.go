package echo

import (
	"fmt"
)

type EchoActionConf struct {
	Message string `yaml:"message" validate:"required"`
}

type EchoAction struct {
	o string
}

func NewEchoAction() *EchoAction {
	return &EchoAction{}
}

func (a *EchoAction) Name() string {
	return "echo"
}

func (a *EchoAction) Doc() string {
	return "echo message"
}

func (a *EchoAction) DefaultConf() interface{} {
	return EchoActionConf{
		Message: "hello world!",
	}
}

func (a *EchoAction) Run(conf interface{}) error {
	c, ok := conf.(EchoActionConf)
	if !ok {
		return fmt.Errorf("conf type is not correct")
	}
	fmt.Println(c.Message)
	a.o = c.Message
	return nil
}
