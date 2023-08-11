package golang_web

import (
	"fmt"
	"net/http"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "name is empty")
	}
}
