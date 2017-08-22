package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var mylog = log.New(os.Stderr, "app: ", log.LstdFlags|log.Lshortfile)

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
		content, parseError := ioutil.ReadAll(response.Body)
		if parseError != nil {
			return parseError.Error()
		}
		return string(content)
	}

	return "Could not complete request"
}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		mylog.Fatalln("You need to pass in the url as a parameter")
	}

	url := args[0]

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s %s ", url, ">")
		scanner.Scan()
		text := scanner.Text()
		responseText := executePost(url, text)
		fmt.Println(responseText)
	}
}
