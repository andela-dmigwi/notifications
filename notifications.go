package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func createToken(credentials string) string {
	encoded := b64.StdEncoding.EncodeToString([]byte(credentials))
	return strings.TrimRight(encoded, "=")
}

func makeAPICall(token string) {
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", "https://api.github.com/notifications", nil)

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Add("Content_type", "application/json")
	resp, _ := client.Do(req)

	type Details []struct {
		Subject struct {
			Title string `json:"title"`
		} `json:"subject"`
	}

	var details Details
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &details)
	for _, element := range details {
		fmt.Println(element)
	}

}

func main() {

	var (
		username string
		password string
	)

	fmt.Println("Enter Your github Username:")
	if _, err := fmt.Scanf("%s", &username); err != nil {
		fmt.Printf(" Error: %s\n Occured", err)
		return
	}
	fmt.Println("Enter Your github password:")
	if _, err := fmt.Scanf("%s", &password); err != nil {
		fmt.Printf(" Error: %s\n Occured", err)
		return
	}
	user := fmt.Sprintf("%s:%s", username, password)
	token := createToken(user)
	fmt.Printf("\nToken ::%s", token)
	fmt.Printf("\n")
	makeAPICall(token)
}
