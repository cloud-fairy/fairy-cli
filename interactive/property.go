package interactive

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type PropertyOption struct {
	label string
	value string
	hint  string ""
}

type PropertyInfo struct {
	property_type    string
	property_name    string
	property_options []PropertyOption
}

func AddOption() (*PropertyOption, error) {
	var p promptui.Prompt
	var err error
	opt := PropertyOption{
		label: "",
		value: "",
	}
	p = promptui.Prompt{Label: "Displayed label"}
	opt.label, err = p.Run()
	if err != nil {
		return nil, err
	}
	p = promptui.Prompt{Label: "Option value"}
	opt.value, err = p.Run()
	if err != nil {
		return nil, err
	}
	p = promptui.Prompt{Label: "Hint display (optional, can be empty)"}
	opt.hint, err = p.Run()
	if err != nil {
		return nil, err
	}
	return &opt, err
}

func AskProperty() (*PropertyInfo, error) {

	result := PropertyInfo{
		property_type:    "",
		property_name:    "",
		property_options: []PropertyOption{},
	}

	p := promptui.Prompt{
		IsConfirm: true,
		Label:     "Add a config property",
	}
	_, err := p.Run()
	if err != nil {
		return nil, err
	}
	pName := promptui.Prompt{Label: "- enter property name (example: user_name)"}
	pType := promptui.Select{
		Label: "- property type",
		Items: []string{"string", "options"},
	}
	result.property_name, _ = pName.Run()
	_, result.property_type, _ = pType.Run()

	if result.property_type == "options" {
		var add_option = true
		for add_option {
			opt, err := AddOption()
			if err == nil {
				result.property_options = append(result.property_options, *opt)
			}
			p := promptui.Prompt{
				Label:     fmt.Sprint("Add another option (currently ", len(result.property_options), "options)"),
				IsConfirm: true,
			}
			_, err = p.Run()
			if err != nil {
				add_option = false
			}
		}
	}
	return &result, nil
}
