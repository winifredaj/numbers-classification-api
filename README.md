<<<<<<< HEAD
# Number Classification API

An API that provides interesting mathematical properties and a fun fact about a given number.

## Features

The Number Classification API provides the following features:

1. Number Classification:

- Checks if a number is prime.
- Checks if a number is perfect.
- Determines whether a number is an Armstrong number.
- Caculates the digit sum of a number.
- Determines the parity (even/odd) of a number.
- Returns a list of applicable properties.

2. Fun Facts:

- Fetches interesting trivia or facts about numbers via an external API (http://numbersapi.com).

3. Cross-Origin Resource Sharing (CORS):
   Ensures compatibility with frontend clients by supporting cross-origin requests.

## API Specification

### Endpoint: `GET /api/classify-number`

#### Query Parameters:

- `number` (required): The number to classify (integer).

#### Response Example (200 OK):

```json
{
  "number": 371,
  "is_prime": false,
  "is_perfect": false,
  "properties": ["armstrong", "odd"],
  "digit_sum": 11,
  "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

## Response Fields:

- number: The input number.
- is_prime: Boolean indicating if the number is a prime number.
- is_perfect: Boolean indicating if the number is a perfect number.
- properties: Array of additional properties (e.g., "armstrong", "even/odd").
- digit_sum: Sum of the digits of the number.
- fun_fact: An interesting fact about the number.

### Code Overview

## Main Components

1. main.go:

- Entry point of the application.
- Sets up the HTTP server and binds the /api/classify-number endpoint to the appropriate handler.

2. handlers Package:

- Contains the ClassifyNumberHandler function, which handles incoming API requests, processes the number, and generates the response.

3. utils Package:

- Contains utility functions for:
- Mathematical computations:
  - IsPrime: Checks if a number is prime.
  - IsPerfect: Checks if a number is perfect.
  - IsArmstrong: Checks if a number is an Armstrong number.
  - DigitSum: Calculates the sum of digits.
  - Parity: Determines whether a number is even or odd.
- CORS Middleware:
  - Ensures cross-origin compatibility.
- Fun Fact Fetching:
  - GetFunFact: Fetches fun trivia about numbers using an external API.
    Project Structure

## Environment Variables

- The server port is configurable via the PORT environment variable in a .env file. Defaults to 8080 if not specified.

## Error Handling

- Invalid input is handled gracefully, returning descriptive error messages to the client.
- If fetching fun facts fails, a default message ("Failed to fetch fun fact") is returned.

Contributions are welcome! Please follow these steps:

Fork the repository.
Create a new branch for your feature/bug fix.
Commit your changes.
Push to your branch and submit a pull request.
License

This project is licensed under the MIT License. See the LICENSE file for details.

This documentation provides an overview of the project, its functionality, and clear instructions for usage, making it easier for developers and users to understand and contribute.
=======
# Numbers-classification-API
>>>>>>> 294204c5acf1812e867f54972fe2bec4fc876263
