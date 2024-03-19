package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func palindromeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Empty text", http.StatusBadRequest)
		return
	}

	result := isPalindrome(text)
	fmt.Fprintf(w, "%t", result)
}

var startedAt = time.Now()

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome API")
}

func main() {
	http.HandleFunc("/palindrome", palindromeHandler)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/", Home)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
