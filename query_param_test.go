package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(writer, "hello")
	} else {
		fmt.Fprintf(writer, "hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {

	request := httptest.NewRequest("GET", "httpp:localhost:8080/hello?name=batman", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
