package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func main() {
	body := []byte(`{
"host": "wenstudy.com",
"key": "86dbd11a307c45ccb46b89c8d1f45bff",
"keyLocation": "https://wenstudy.com/86dbd11a307c45ccb46b89c8d1f45bff.txt",
"urlList": [
  "https://wenstudy.com",
  "https://wenstudy.com/en/",
  ]}`)

	req, err := http.NewRequest("POST", "https://api.indexnow.org/IndexNow", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("failed to create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

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
