package main

/**
* This Program that lists all github notifications once
* a user logs in.
**/

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// Error Response
const errResp string = "Your username or password is wrong"

func createToken(credentials string) string {
	matched, err := regexp.MatchString(`(([\w\d\-*_#]+){1}(:){1}([\w\d\-*_#]+){1})`, credentials)
	if err != nil {
		fmt.Println("Error occured while validating your credentials")
	}
	if !matched {
		return errResp
	}
	encoded := b64.StdEncoding.EncodeToString([]byte(credentials))
	return strings.TrimRight(encoded, "=")
}

func makeAPICall(token string) {
	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("GET", "https://api.github.com/notifications", nil)
	if err != nil {
		fmt.Printf("Error Occurred : %s\n", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Add("Content_type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error Occurred : %s\n", err)
	}

	type Details []struct {
		Subject struct {
			Title string `json:"title"`
		} `json:"subject"`
	}
	var details Details
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error Occurred : %s\n", err)
	}

	err = json.Unmarshal(body, &details)
	if err != nil {
		fmt.Printf("Error Occurred : %s\n", err)
	}

	fmt.Print("\nListing Notifications")
	for position, element := range details {
		fmt.Printf("\n%d. %s", (position + 1), element.Subject.Title)
	}

}

func main() {
	var username string
	fmt.Println("Enter Your github Username:")
	if _, err := fmt.Scanf("%s", &username); err != nil {
		fmt.Printf(" Error: %s Occured\n", err)
		return
	}
	fmt.Println("Enter Your github password:")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Printf(" Error: %s Occured\n", err)
		return
	}
	user := fmt.Sprintf("%s:%s", username, string(password))
	token := createToken(user)
	// If err terminate the program and return the error
	if token == errResp {
		fmt.Println(token)
	} else {
		makeAPICall(token)
	}

}
