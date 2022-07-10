package gohttpclient

import "theguy951357/go_httpclient/gohttp"

func exampleUsage() {
	client := gohttp.New()

	client.Get()
}
