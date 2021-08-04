package core

type ActionRunFunc func(conf interface{}) (interface{}, error)

type Action struct {
	Name        string
	Doc         string
	DefaultConf interface{}
	Run         ActionRunFunc
}
