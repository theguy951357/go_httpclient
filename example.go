package gohttpclient

import "github.com/theguy951357/go_httpclient/gohttp"

func basicExample() {
	client := gohttp.New()

	client.Get()
}
