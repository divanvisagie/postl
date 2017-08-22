package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func executePost(url string, data string) string {
	var dataBytes = []byte(data)

	var client http.Client
	request, requestCreationError := http.NewRequest("POST", url, bytes.NewBuffer(dataBytes))
	request.Header.Set("Content-Type", "text/plain")
	if requestCreationError != nil {
		return requestCreationError.Error()
	}

	response, requestExecutionError := client.Do(request)
	if requestExecutionError != nil {
		return requestExecutionError.Error()
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		bodyBytes, parseErr := ioutil.ReadAll(response.Body)
		if parseErr != nil {
			return parseErr.Error()
		}
		bodyString := string(bodyBytes)
		return bodyString
	}

	return "Could not complete request"
}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		panic("You need to pass in the url as a parameter")
	}

	url := args[0]

	for {
		var text string
		fmt.Printf("%s %s ", url, ">")
		fmt.Scanln(&text)

		responseText := executePost(url, text)

		fmt.Println(responseText)
	}
}
