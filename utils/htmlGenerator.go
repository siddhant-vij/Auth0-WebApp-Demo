package utils

import (
	"html/template"
	"os"
)

func GenerateHtml(tpl *template.Template, fileName string, content interface{}) error {
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
