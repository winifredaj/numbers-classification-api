package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/winifredaj/number-classification-api/utils"
	"strconv"
)

// NumberClassificationResponse defines the structure for the JSON response.
// It includes various properties of the number being classified.
type NumberClassificationResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

// ClassifyNumberHandler handles HTTP requests to classify a given number.
// It performs various classifications (prime, perfect, etc.), calculates properties, and fetches a fun fact.
// Endpoint: /classify?number={number}
func ClassifyNumberHandler(w http.ResponseWriter, r *http.Request){
	
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract the 'number' query parameter from the request
	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr) // Convert the number from string to integer

  if err != nil{
		// If conversion fails, respond with a 400 Bad Request error
		http.Error(w,fmt.Sprintf(`{"number": "%s","error": Not a number}`,numberStr), http.StatusBadRequest)
		return
	}

// Perform various classifications using utility functions
isPrime := utils.IsPrime(number)
isPerfect := utils.IsPerfect(number)
isArmstrong := utils.IsArmstrong(number)
digitSum := utils.DigitSum(number)
parity := utils.Parity(number)


// Build a list of properties based on the classifications
properties := []string{}
if isArmstrong {
	properties = append(properties, "armstrong") // Add "armstrong" to properties if applicable
}
properties = append(properties, parity) // Add "odd" or "even" to properties

// Fetch a fun fact about the number
funFact, err := utils.GetFunFact(number)
if err != nil{
	// If fetching the fun fact fails, respond with a 500 Internal Server Error
	http.Error(w, `{"error": true, "message": "Failed to fetch fun fact"}`,http.StatusInternalServerError)
	return
}

// Construct the response payload
response := map[string]interface{}{
	"number": number,
	"is_prime": isPrime,
	"is_perfect": isPerfect,
	"properties": properties,
	"digit_sum": digitSum,
	"fun_fact": funFact,

}
// Encode the response as JSON and send it to the client
json.NewEncoder(w).Encode(response)
}
