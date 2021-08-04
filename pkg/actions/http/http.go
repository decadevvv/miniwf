package http

import (
	"fmt"
	"net/http"

	"github.com/decadevvv/miniwf/pkg/core"
)

type HTTPActionConf struct {
	Operate string `yaml:"operate" validate:"required,oneof=get"`
	URL     string `yaml:"url" validate:"required,url"`
}

var HTTPAction = core.Action{
	Name: "http",
	Doc:  "do http requests",
	DefaultConf: HTTPActionConf{
		Operate: "get",
		URL:     "",
	},
	Run: func(conf interface{}) (interface{}, error) {
		c, ok := conf.(HTTPActionConf)
		if !ok {
			return nil, fmt.Errorf("conf type is not correct")
		}
		switch c.Operate {
		case "get":
			resp, err := http.Get(c.URL)
			if err != nil {
				return nil, err
			}
			return *resp, nil
		default:
			return nil, fmt.Errorf("undefined operate %s", c.Operate)
		}
	},
}
