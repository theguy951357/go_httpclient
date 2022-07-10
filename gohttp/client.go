package gohttp

type httpclient struct{}

func New() HttpClient {
	client := &httpclient{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

func (c *httpclient) Get() {}

func (c *httpclient) Post() {}

func (c *httpclient) Put() {}

func (c *httpclient) Patch() {}

func (c *httpclient) Delete() {}
