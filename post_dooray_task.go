package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	project := os.Getenv("INPUT_PROJECT")
	token := os.Getenv("INPUT_TOKEN")
	subject := os.Getenv("INPUT_SUBJECT")
	content := os.Getenv("INPUT_CONTENT")

	fmt.Println("project: ", project)
	fmt.Println("token: ", token)
	fmt.Println("subject: ", subject)
	fmt.Println("content: ", content)

	var jsonStr = []byte(`{
		"parentPostId": null,
		"users": {
			"to": [],
			"cc": []
		},
		"subject": "` + subject + `",
		"body": {
			"mimeType": "text/html",
			"content": "` + content + `"
		},
		"dueDate": null,
		"dueDateFlag": true,
		"milestoneId": null,
		"tagIds": [],
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
