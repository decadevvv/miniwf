package core

type Workflow struct {
	Name  string `yaml:"name" validate:"required"`
	Doc   string `yaml:"doc" validate:"required"`
	Steps []Step `yaml:"steps" validate:"required,dive"`
}

type Step struct {
	Name   string            `yaml:"name" validate:"required"`
	Doc    string            `yaml:"doc" validate:"require"`
	Action string            `yaml:"action" validate:"required"`
	Conf   map[string]string `yaml:"conf" validate:"required"`
	Output string            `yaml:"output" validate:"required"`
}
