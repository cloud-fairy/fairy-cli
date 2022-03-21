package templates

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
)

type File struct {
	name    string
	content []byte
}

type Template struct {
	name    string
	content []File
}

var templates []Template

//go:embed connector component
var staticFiles embed.FS

func loadTemplates(dir string, index string) {
	template := Template{
		name: index,
	}
	files, err := staticFiles.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var fileData, err = staticFiles.ReadFile(dir + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		template.content = append(template.content, File{
			name:    file.Name(),
			content: fileData,
		})
	}
	templates = append(templates, template)
}

func Load() {
	if len(templates) > 0 {
		return
	}

	files, _ := staticFiles.ReadDir(".")

	for _, file := range files {
		if file.IsDir() == true {
			loadTemplates(file.Name(), file.Name())
		}
	}
}

func Find(template string) (*Template, error) {
	for _, tpl := range templates {
		if tpl.name == template {
			return &tpl, nil
		}
	}
	return nil, errors.New("Could not find template")
}

func Write(template string, out string) {
	tpl, err := Find(template)
	if err != nil {
		log.Fatal("Could not load template")
	}

	os.MkdirAll(out, os.ModePerm)

	for _, file := range tpl.content {
		filename := out + "/" + file.name
		fmt.Println("Writing file", filename)
		fileHandle, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		fileHandle.Write(file.content)
	}
}
