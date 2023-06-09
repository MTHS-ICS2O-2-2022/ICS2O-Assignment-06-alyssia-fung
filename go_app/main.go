package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getDogImage(urlAddress string) (string, error) {
	parsedURL, err := url.Parse(urlAddress)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Message string `json:"message"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return result.Message, nil
}

func main() {
	urlAddress := "https://dog.ceo/api/breeds/image/random"
	imageURL, err := getDogImage(urlAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Dog Image: %s\n", imageURL)
	fmt.Println("\nDone.")
}
