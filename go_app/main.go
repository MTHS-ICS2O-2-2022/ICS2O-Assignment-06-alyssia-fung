package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getDogImage(URLAddress string) (string, error) {
	resp, err := http.Get(URLAddress)
	if err != nil {
		return "", err
	}

	// this is were i got defer resp.Body.Close() from
	//https://stackoverflow.com/questions/23928983/defer-body-close-after-receiving-response
	defer resp.Body.Close()

	// this is were i got body, err := ioutil.ReadAll(resp.Body) from
	//https://www.educative.io/answers/how-to-read-the-response-body-in-golang
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Message string `json:"message"`
	}

	// this is were i got json.Unmarshal(body, &result) from
	//stackoverflow.com/questions/47693238/how-to-judge-unmarshal-json-interface-type-in-golang
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return result.Message, nil
}

func main() {
	URLAddress := "https://dog.ceo/api/breeds/image/random"
	imageURL, err := getDogImage(URLAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Dog Image: %s\n", imageURL)
	fmt.Println("\nDone.")
}
