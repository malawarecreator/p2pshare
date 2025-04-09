package main

import (
	"fmt"
	"io"
	"mime" // Required for MIME type detection
	"net/http"
	"os"
	"path/filepath" 
)

var args = os.Args

func main() {
	if len(args) < 3 {
		fmt.Println("Usage: p2pshare <filepath> <authtoken>")
		return
	}

	filename := args[1]
	expectedToken := args[2]

	cleanedFilename := filepath.Clean(filename)

	fileInfo, err := os.Stat(cleanedFilename)
	if os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' not found\n", cleanedFilename)
		return
	}
    if err != nil {
         fmt.Printf("Error: Cannot stat file '%s': %v\n", cleanedFilename, err)
         return
    }
	if fileInfo.IsDir() {
		fmt.Printf("Error: '%s' is a directory, not a file\n", cleanedFilename)
		return
	}


	fmt.Println("Fetching public IP address...")
	resp, err := http.Get("https://api.ipify.org")
	publicIP := "YOUR_PUBLIC_IP (fetch failed)"
	if err != nil {
		fmt.Println("Warning: Could not fetch public IP:", err)
		fmt.Println("External access URL will not be accurate.")
	} else {
		defer resp.Body.Close()
		body, errRead := io.ReadAll(resp.Body)
		if errRead != nil {
			fmt.Println("Warning: Could not read public IP response:", errRead)
		} else {
			publicIP = string(body)
		}
	}


	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		queryAuthToken := r.URL.Query().Get("authtoken")

		if queryAuthToken != expectedToken {
			http.Error(w, "Auth token incorrect or missing", http.StatusUnauthorized)
			fmt.Printf("Rejected download attempt from %s: incorrect token '%s'\n", r.RemoteAddr, queryAuthToken)
			return
		}

		fileExtension := filepath.Ext(cleanedFilename)
		mimeType := mime.TypeByExtension(fileExtension)

		if mimeType == "" {
			mimeType = "application/octet-stream" 
			fmt.Printf("Warning: Could not determine MIME type for extension '%s'. Defaulting to '%s'.\n", fileExtension, mimeType)
		}

		w.Header().Set("Content-Type", mimeType)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(cleanedFilename)))

		fmt.Printf("Serving file '%s' (Type: %s) to %s\n", cleanedFilename, mimeType, r.RemoteAddr)
		http.ServeFile(w, r, cleanedFilename)
	})

	fmt.Println("\n--- Server Starting ---")
	fmt.Printf("Serving file: %s\n", cleanedFilename)
	fmt.Printf("Required auth token: %s\n", expectedToken)
	fmt.Println("Listening on port 8888...")

	fmt.Println("\n--- Access URLs ---")
	fmt.Println("From THIS machine:    http://localhost:8888/download?authtoken=" + expectedToken)
	fmt.Println("From SAME network:    http://<YOUR_LOCAL_IP>:8888/download?authtoken=" + expectedToken + " (Find local IP with 'ipconfig' or 'ifconfig')")
	fmt.Println("From OUTSIDE network: http://" + publicIP + ":8888/download?authtoken=" + expectedToken + " (Requires firewall & router port forwarding for port 8888)")

	fmt.Println("\nServer running. Press Ctrl+C to stop.")
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("\nFatal Error: Could not start server:", err)
		os.Exit(1)
	}
}