package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// GetFunFact fetches a fun fact for the given number from the Numbers API.
func GetFunFact(number int) (string, error){

// Construct the URL for the Numbers API
 url := fmt.Sprintf("http://numbersapi.com/%d?json", number)
 
// Make an HTTP GET request
 resp, err := http.Get(url)
 if err != nil {
	log.Println("Error fetching fun fact:",err)
	return "", fmt.Errorf("failed to fetch fun fact: %w", err)
 }
 defer resp.Body.Close()

// Read the response body
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil {
	log.Println("Error reading response body:", err)
	return "", fmt.Errorf("failed to send response: %w", err)
 }

 //Parse the JSON response
 var result map[string]interface{}
 if err:= json.Unmarshal(body, &result); 
 err != nil {
	log.Println("Error decoding JSON response:", err)
	return "", fmt.Errorf("failed to decode JSON:%w", err)
 }

 // Extract the "text" field
 if text, exists := result["text"].(string); 
 exists {
	return text, nil

 }
 // Default message if the "text" field is not found
 return "No fun fact available", nil

}