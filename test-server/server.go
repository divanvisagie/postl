package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Command struct {
	text string
}

func handler(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)

	fmt.Printf("got: %s \n", bodyString)

	fmt.Fprintf(w, "Parsed:  %s!", bodyString)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
