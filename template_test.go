package golang_web

import (
	"html/template"
	"net/http"
)

func TemplateText(writer http.ResponseWriter, request *http.Request) {
	templateTest := "<html><body>{{.}}></body></html>"
	//t, err := template.New("SIMPLE").Parse(templateTest)
	//if err != nil {
	//	panic(err)
	//}

	t := template.Must(template.New("SIMPLE").Parse(templateTest))
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}
