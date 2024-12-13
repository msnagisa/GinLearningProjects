package main

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	url := "http://127.0.0.1:8080/hello"
	data := map[string]string{ "msg": "hello world after Post" }
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status", resp.Status)
	fmt.Println("Response body", resp.Body)
}