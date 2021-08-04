package echo

import (
	"fmt"
	"net/http"
)

type HTTPActionConf struct {
	Operate string `yaml:"operate" validate:"required,oneof=get"`
	URL     string `yaml:"url" validate:"required,url"`
}

type HTTPAction struct {
	o http.Response
}

func NewHTTPAction() *HTTPAction {
	return &HTTPAction{}
}

func (a *HTTPAction) Name() string {
	return "http"
}

func (a *HTTPAction) DefaultConf() interface{} {
	return HTTPActionConf{
		Operate: "get",
		URL:     "",
	}
}

func (a *HTTPAction) Run(conf interface{}) error {
	c, ok := conf.(HTTPActionConf)
	if !ok {
		return fmt.Errorf("conf type is not correct")
	}
	if c.Operate == "get" {
		resp, err := http.Get(c.URL)
		if err != nil {
			return err
		}
		a.o = *resp
	}
	return nil
}

func (a *HTTPAction) Output() interface{} {
	return a.o
}

func (a *HTTPAction) Doc() string {
	return ""
}
