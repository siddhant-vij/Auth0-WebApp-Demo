package utils

import (
	"html/template"
	"os"
)

func GenerateHtml(tpl *template.Template, fileName string, content interface{}) error {
	err := CopyFiles("static", "public")
	if err != nil {
		return err
	}

	file, err := os.Create("public/" + fileName + ".html")
	if err != nil {
		return err
	}

	err = tpl.ExecuteTemplate(file, fileName+".gohtml", content)
	if err != nil {
		return err
	}
	return nil
}
