package core

type Action interface {
	Name() string
	DefaultConf() interface{}
	Run(conf interface{}) error
	Output() interface{}
	Doc() string
}
