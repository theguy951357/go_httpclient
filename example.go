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
	getUrls()
}

type User struct {
	FirstName string 'json: "first_name"'
	LastName string 'json: "last_name"'
	
}

func getUrls() {

	headers := make(http.Header)
	//headers.Set("Authorization", "Bearer ABC-123")

	response, err := httpclient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}

func createUser(user User) {
	//requestBody := json.Marshal(user)
	//headers := make(http.Header)
	//headers.Set("Authorization", "Bearer ABC-123")

	response, err := httpclient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}
