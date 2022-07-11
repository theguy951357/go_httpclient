package main

import (
	"fmt"
	"net/http"

	"github.com/theguy951357/go_httpclient/gohttp"
)

var (
	httpclient = gohttp.New()
)

func main() {
	basicExample()
}

func basicExample() {

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := httpclient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}
