package gohttpclient

import (
	"fmt"

	"github.com/theguy951357/go_httpclient/gohttp"
)

func basicExample() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}
