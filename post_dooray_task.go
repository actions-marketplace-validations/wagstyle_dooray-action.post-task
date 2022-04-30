package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	projectId := os.Args[0]
	authorizationToken := os.Args[1]

	var jsonStr = []byte(`{
		"parentPostId": null,
		"users": {
			"to": [],
			"cc": []
		},
		"subject": "제목을 입력합니다.",
		"body": {
			"mimeType": "text/html",
			"content": "본문을 입력합니다."
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
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))
}
