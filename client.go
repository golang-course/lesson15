package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	//GET на /api/phones

	// resp, err := http.Get("http://localhost:8080/api/phones/1")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	//POST на /contacts
	// message := []byte(`{"Email": "ivanov@example.com", "Message": "All good!"}`)
	// resp, err := http.Post("http://localhost:8080/contacts", "application/json", bytes.NewBuffer(message))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	message := []byte(`{"status": "ok"}`)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/unknown", bytes.NewBuffer(message))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(body))

}
