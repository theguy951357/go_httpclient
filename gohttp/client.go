package gohttp

import "fmt"

type httpclient struct{}

func New() HTTPClient {
	client := &httpclient{}
	return client
}

type HTTPClient interface {
	Get()
	Post()
	Pull()
	Patch()
	Delete()
}

func (c *httpclient) Get() {
	fmt.Println("Mind your fucking business you park bench!!!")
}

func (c *httpclient) Post() {}

func (c *httpclient) Pull() {}

func (c *httpclient) Patch() {}

func (c *httpclient) Delete() {}
