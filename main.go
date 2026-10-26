package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	username := os.Args[1]

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln("Failed generating new request: ", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Request failed: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Unexpected status: ", resp.StatusCode)
	}

	fmt.Println("Response: ", resp)
}
