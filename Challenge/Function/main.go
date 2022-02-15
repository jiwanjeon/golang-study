/*
- Write a function that gets a URL and returns the value of Content-Type response HTTP header

- The function should return an error if it can't perform a GET request

fun contentType(url string) (string,error)
	resp, err := http.Get("https:/golang.org") --> Use net/http Get function to make an HTTP call
	resp.Header.Get("Content-Length") --> Use resp.Header.Get to get the value of a header --> Use resp.Header.Get to get the value of a header

func contentType(url string) (string,error)
	resp.Body.Close() --> Make sure the response body is closed
*/

package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	fmt.Println("ctype: ", ctype)

	return ctype, nil
}

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
