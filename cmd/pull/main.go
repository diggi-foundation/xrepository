package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const url = "https://www.xrepository.de/api/swagger"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Status error: %v", resp.StatusCode)
	}

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, responseBody, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	out, err := os.Create("swagger.json")
	if err != nil {
		fmt.Printf("Error creating swagger.json file: %v", err)
	}
	defer out.Close()

	_, err = out.Write(prettyJSON.Bytes())
	if err != nil {
		fmt.Println("Error writing JSON to swagger.json:", err)
		return
	}

	fmt.Println("Swagger schema has been written to swagger.json")
}
