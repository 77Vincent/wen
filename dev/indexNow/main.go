package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 10,
	}
	payload = map[string]interface{}{
		"host":        "wenstudy.com",
		"key":         "d1b801009c857fd452e1f9086f9b567e",
		"keyLocation": "https://wenstudy.com/d1b801009c857fd452e1f9086f9b567e.txt",
		"urlList": []string{
			"https://wenstudy.com",
			"https://wenstudy.com/en/",
		},
	}
)

func main() {

	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("failed to marshal body: %s", err)
	}
	req, err := http.NewRequest("POST", "https://api.indexnow.org/IndexNow", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read res body: %s", err)
	}

	if res.StatusCode > 399 {
		log.Fatalf("something wrong: %s", res.Status)
	}

	log.Println("Response:", res.Status)
	log.Println(string(resBody))
}
