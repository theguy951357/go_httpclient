package main

import (
	"fmt"
	"net/http"

	"github.com/theguy951357/go_httpclient/gohttp"
)

var (
	httpclient = gohttp.New()
)

func getGithubClient() gohttp.HTTPClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)
	return client
}

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
