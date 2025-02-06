package utils

import "net/http"

// CORS is a middleware function that adds Cross-Origin Resource Sharing (CORS) headers
// to the HTTP response. It ensures that the API can be accessed from different origins (domains).
// Parameters:
// - next: The next handler function in the middleware chain.
// Returns:
// - A wrapped HTTP handler function that includes CORS headers.

func CORS(next http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//Set the "Access-Control-Allow-Origin" header to allow all origins (*).
		w.Header().Set("Access-Control-Allow-Origin", "*")

		//Set the "Access-Control-Allow-Methods" header to specify allowed HTTP methods.
		w.Header().Set("Access-Control-Allow-Origin", "GET, OPTIONS")

		// Set the "Access-Control-Allow-Headers" header to specify allowed headers in requests.
		w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
		
		// Check if the request method is OPTIONS, which is used for CORS preflight requests.
		if r.Method == http.MethodOptions{
			// For preflight requests, respond with HTTP 200 (OK) and exit without calling the next handler.
			w.WriteHeader(http.StatusOK)
			return
		}
		// If it's not a preflight request, proceed to the next handler in the middleware chain.
		next(w,r)
	}

}