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

const errResp string = "Your username or password is wrong"

func createToken(credentials string) string {
	matched, _ := regexp.MatchString(`(([\w\d\-*_#]+){1}(:){1}([\w\d\-*_#]+){1})`, credentials)
	if !matched {
		return errResp
	}
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
	fmt.Print("\nListing Notifications")
	for position, element := range details {
		fmt.Printf("\n%d. %s", (position + 1), element.Subject.Title)
	}

}

func main() {
	var username string
	fmt.Println("Enter Your github Username:")
	if _, err := fmt.Scanf("%s", &username); err != nil {
		fmt.Printf(" Error: %s\n Occured", err)
		return
	}
	fmt.Println("Enter Your github password:")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Printf(" Error: %s\n Occured", err)
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
