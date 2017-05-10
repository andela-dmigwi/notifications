package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const correctSampleCredentials string = "username:password"
const incorrectSampleCredentials string = "usernamepassword"

func TestCreateToken_With_Correct_Credentials(t *testing.T) {
	expectedResponse := "dXNlcm5hbWU6cGFzc3dvcmQ"
	if token := createToken(correctSampleCredentials); token != expectedResponse {
		t.Fatalf("Expected : %s, but got : %s", expectedResponse, token)
	}
}

func TestCreateToken_With_Incorrect_Credentials(t *testing.T) {
	expectedResponse := "Your username or password is wrong"
	if token := createToken(incorrectSampleCredentials); token != expectedResponse {
		t.Fatalf("Expected : %s, but got : %s", expectedResponse, token)
	}
}

func TestMakeAPICall(t *testing.T) {
	sampleToken := "TYYQhhgsd8Iyhjhjyuhaau"
	sampleOutput := `[{
						"subject": {
						"title": "New PR merged",
						"latest_comment_url": null,
						"type": "PullRequest"
						}
					}]`
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/notifications",
		httpmock.NewStringResponder(200, sampleOutput))

	// Setup Capture output methods
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	makeAPICall(sampleToken)

	w.Close()
	output, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Contains("1. New PR merged\n", string(output), "Output should contain the notification")

}

func TestMainl(t *testing.T) {
}
