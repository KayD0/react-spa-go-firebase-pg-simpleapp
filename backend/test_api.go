package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// APIResponse represents the response from the API
type APIResponse struct {
	Message   string                 `json:"message"`
	Endpoints map[string]interface{} `json:"endpoints"`
}

// TestAPI tests the API by making a request to the root endpoint
func TestAPI() {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:5000"
	}

	// Check if the API is running
	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Printf("Error: Could not connect to the API. Make sure it's running with 'go run main.go'\n")
		fmt.Printf("Error details: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API returned status code %d\n", resp.StatusCode)
		return
	}

	// Parse the response
	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	// Print the API information
	fmt.Printf("\nAPI Status: %s\n\n", apiResp.Message)
	fmt.Println("Available Endpoints:")
	for endpoint, description := range apiResp.Endpoints {
		fmt.Printf("- %s: %s\n", endpoint, description)
	}
}

func main() {
	// Parse command line flags
	flag.Parse()

	// Run the test
	TestAPI()
}
