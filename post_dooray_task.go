package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	projectId := os.Getenv("INPUT_PROJECT_ID")
	authorizationToken := os.Getenv("INPUT_AUTHORIZATION_TOKEN")
	releaseName := os.Getenv("INPUT_RELEASE_NAME")
	releaseBody := os.Getenv("INPUT_RELEASE_BODY")

	fmt.Println("projectId: ", projectId)
	fmt.Println("authorizationToken: ", authorizationToken)
	fmt.Println("releaseName: ", releaseName)
	fmt.Println("releaseBody: ", releaseBody)

	var jsonStr = []byte(`{
		"parentPostId": null,
		"users": {
			"to": [],
			"cc": []
		},
		"subject": "` + releaseName + `",
		"body": {
			"mimeType": "text/html",
			"content": "` + releaseBody + `"
		},
		"dueDate": null,
		"dueDateFlag": true,
		"milestoneId": null,
		"tagIds": [],
		"priority": "none"
	}`)

	url := fmt.Sprintf("https://api.dooray.com/project/v1/projects/%s/posts", projectId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("dooray-api %s", authorizationToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Status: ", resp.Status)
	fmt.Println("response Headers: ", resp.Header)
	fmt.Println("response Body: ", string(body))
}
