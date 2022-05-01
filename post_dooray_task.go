package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	var (
		project   = os.Getenv("PROJECT_ID")
		token     = os.Getenv("AUTHORIZATION_TOKEN")
		subject   = os.Getenv("SUBJECT")
		content   = os.Getenv("CONTENT")
		recipient = os.Getenv("RECIPIENT")
		tag       = os.Getenv("TAG")
	)

	fmt.Println("project: ", project)
	fmt.Println("token: ", token)
	fmt.Println("subject: ", subject)
	fmt.Println("content: ", content)
	fmt.Println("recipient: ", recipient)
	fmt.Println("tag: ", tag)

	var jsonStr = []byte(`{
		"parentPostId": null,
		"users": {
			"to": [` + recipient + `],
			"cc": []
		},
		"subject": "` + subject + `",
		"body": {
			"mimeType": "text/x-markdown",
			"content": "` + content + `"
		},
		"dueDate": null,
		"dueDateFlag": true,
		"milestoneId": null,
		"tagIds": ["` + tag + `"],
		"priority": "none"
	}`)

	url := fmt.Sprintf("https://api.dooray.com/project/v1/projects/%s/posts", project)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("dooray-api %s", token))

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
