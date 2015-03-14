package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	command := fmt.Sprintf("node index.js %v", "chris")
	output, _ := exec.Command("/bin/bash", "-c", command).CombinedOutput()
	w.Write([]byte(string(output)))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
