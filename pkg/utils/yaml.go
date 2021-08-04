package utils

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

func UnmarshalMapIntoStructWithTemplate(input map[string]string, output interface{}, contextValues map[string]interface{}, debug bool) error {
	if debug {
		fmt.Println("default conf:")
		spew.Dump(output)
	}
	inputTemplatedString, err := MarshalMapWithTemplate(input, contextValues, debug)
	if err != nil {
		return fmt.Errorf("failed to marshal input with template values: %w", err)
	}
	inputTemplatedMap, err := UnmarshalStringIntoMap(inputTemplatedString)
	if err != nil {
		return fmt.Errorf("failed to unmarshal templated input map into generic map: %w", err)
	}
	return UnmarshalMapIntoStruct(inputTemplatedMap, output, debug)
}

func UnmarshalStringIntoMap(input string) (map[string]interface{}, error) {
	output := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(input), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func MarshalMapWithTemplate(input map[string]string, values map[string]interface{}, debug bool) (string, error) {
	inputBytes, err := yaml.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("failed to marshal input into bytes: %w", err)
	}
	if debug {
		fmt.Printf("marshalled input: \n%s\n", string(inputBytes))
	}
	tmpl, err := template.New("conf").Parse(string(inputBytes))
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}
	tmplOutput := bytes.NewBuffer([]byte(""))
	err = tmpl.Execute(tmplOutput, values)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	if debug {
		fmt.Printf("marshalled input (templated): \n%s\n", tmplOutput.String())
	}
	return tmplOutput.String(), nil
}

func UnmarshalMapIntoStruct(input map[string]interface{}, output interface{}, debug bool) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  "yaml",
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return fmt.Errorf("failed to create decoder: %w", err)
	}
	decoder.Decode(input)
	if err != nil {
		return fmt.Errorf("failed to unmarshal map into struct: %w", err)
	}
	if debug {
		fmt.Println("unmarshaled conf:")
		spew.Dump(output)
	}
	return nil
}
