package http

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type TemplateData struct {
	StringMap map[string]string
	Data      map[string]interface{}
	CSRFToken string
}

func GetFilePaths(module string, tmplName string, includeLayout bool) []string {
	commonLayouts := []string{"./views/layout.html", "./views/partials/header.html", "./views/partials/sidebar.html", "./views/partials/footer.html"}
	var selectedLayouts []string

	selectedLayouts = commonLayouts

	templatePath := strings.Replace("./views/modulename/templatename.html", "modulename", module, 1)
	templatePath = strings.Replace(templatePath, "templatename", tmplName, 1)

	if includeLayout == true {
		return append(selectedLayouts, templatePath)
	}

	return []string{templatePath}
}

func Render(w http.ResponseWriter, module string, tmplName string, includeLayout bool, data interface{}) {
	content, _ := template.ParseFiles(GetFilePaths(module, tmplName, includeLayout)...)
	w.Header().Set("Content-Type", "text/html")
	err := content.Execute(w, data)

	if err != nil {
		fmt.Println("Error executing template: ", tmplName, "error: ", err)
	}
}
