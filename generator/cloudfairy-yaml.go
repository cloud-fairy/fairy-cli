package generator

import (
	"cloudfairy/interactive"

	"gopkg.in/yaml.v2"
)

type YamlDescriptor struct {
	Type           string
	BuilderRuntime string
	Entry          string

	Internal bool

	Vendor string
	Name   string

	DisplayName string
	Properties  []interactive.PropertyInfo

	From []string
	To   []string
}

func (self YamlDescriptor) AddProperty(p interactive.PropertyInfo) {
	self.Properties = append(self.Properties, p)
}


func CreateDescriptor(
	Vendor string,
	Name string,
	DisplayName string,
	Type string,
) ([]byte, error) {
	
	descriptor := YamlDescriptor{
		Vendor: Vendor,
		Name: Name,
		DisplayName: DisplayName,
		Type: Type,
	}
	result, err := yaml.Marshal(&descriptor)
	return result, err
}
