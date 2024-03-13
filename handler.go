package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

// Define command-line flags
var (
    port    = flag.String("port", "5050", "Port for the server to listen on")
    secret  = flag.String("secret", "my_secret_key", "Secret key for HMAC computation")
    command = flag.String("command", "./handler.sh", "Command to execute upon receiving the webhook")
)

// handler function to handle incoming HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)

	// Check for errors while reading request body
	if err != nil {
		http.Error(w, "[Error] Reading request body", http.StatusBadRequest)

		return
	}

	if *secret != "" {
		// Extract X-Hub-Signature-256 header value
		receivedSignature := strings.TrimPrefix(r.Header.Get("X-Hub-Signature-256"), "sha256=")

		// Compute HMAC with SHA-256 using the secret and the request body
		mac := hmac.New(sha256.New, []byte(*secret))
		mac.Write(body)

		// Calculate the expected HMAC signature
		expectedSignature := hex.EncodeToString(mac.Sum(nil))

		// Compare the computed HMAC with the received signature
		if receivedSignature != expectedSignature {
			http.Error(w, "Signature mismatch", http.StatusUnauthorized)

			return
		}

		// If signatures match, print success message
		fmt.Println("[Success] Signature matched successfully")
	}

	// Execute the command
	if *command != "" {
		cmd := exec.Command(*command)
		err := cmd.Run()

		if err != nil {
			fmt.Println("[Error] Executing command:", err)
		}
	}
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Print information about the server setup
	fmt.Printf("[GitHub Handler] Webhook handler is running:\n\n> port: %s\n> secret: %s\n> command: %s\n", *port, *secret, *command)

	// Register the handler function for incoming requests
	http.HandleFunc("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":"+*port, nil)
}