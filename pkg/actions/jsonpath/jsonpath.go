package jsonpath

import (
	"github.com/decadevvv/miniwf/pkg/core"
)

type JSONPathActionConf struct {
	Source   string `yaml:"source" validate:"required"`
	JSONPath string `yaml:"jsonpath" validate:"required"`
}

var JSONPathAction = core.Action{
	Name: "jsonpath",
	Doc:  "use jsonpath to transform or extract struct and JSONs",
	DefaultConf: JSONPathActionConf{
		Source:   "",
		JSONPath: "",
	},
	Run: func(interface{}) (interface{}, error) { panic("not implemented") },
}
