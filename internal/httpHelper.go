package internal

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func BuildGetRequest(fullUrl, bearerToken string) *http.Request {
	return BuildHttpRequest("GET", fullUrl, bearerToken, nil)
}

func BuildPostRequest(fullUrl, bearerToken string, body io.Reader) *http.Request {
	return BuildHttpRequest("POST", fullUrl, bearerToken, body)
}

func BuildPostRequestNoBody(fullUrl, bearerToken string) *http.Request {
	return BuildPostRequest(fullUrl, bearerToken, nil)
}

func BuildHttpRequest(method, fullUrl, bearerToken string, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, fullUrl, body)

	if bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+bearerToken)
	}

	req.Header.Add("Accept", "application/json")
	if method == "POST" || method == "PUT" && body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		log.Panic("Unable to create new http request: ", err)
	}
	return req
}

func MakeHttpRequest(req *http.Request, verbose bool) (*http.Response, []byte) {

	if verbose {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Panic("Unable dump outgoing http request: ", err)
		}
		log.Println(string(requestDump))
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("Unable to make http call: ", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Unable to read http response: ", err)
	}
	if verbose {
		log.Println(string(responseBody))
	}
	return resp, responseBody
}
