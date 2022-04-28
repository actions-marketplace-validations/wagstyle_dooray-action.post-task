package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

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

	url := "https://api.dooray.com/project/v1/projects/3202204541701520238/posts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "dooray-api ajjt1imxmtj4:en4Wg43eQmC2LE97dVeU7Q")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
