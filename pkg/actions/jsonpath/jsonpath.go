package echo

import "fmt"

type JSONPathActionConf struct {
	Source   string `yaml:"source" validate:"required"`
	JSONPath string `yaml:"jsonpath" validate:"required"`
}

type JSONPathAction struct {
	o string
}

func NewJSONPathAction() *JSONPathAction {
	return &JSONPathAction{}
}

func (a *JSONPathAction) Name() string {
	return "jsonpath"
}

func (a *JSONPathAction) Doc() string {
	return "use jsonpath to transform or extract struct and JSONs"
}

func (a *JSONPathAction) DefaultConf() interface{} {
	return JSONPathActionConf{
		Source:   "",
		JSONPath: "",
	}
}

func (a *JSONPathAction) Run(conf interface{}) error {
	c, ok := conf.(JSONPathActionConf)
	if !ok {
		return fmt.Errorf("conf type is not correct")
	}
	fmt.Println(c.Message)
	a.o = c.Message
	return nil
}

func (a *JSONPathAction) Output() interface{} {
	return a.o
}
