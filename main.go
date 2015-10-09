package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

type Resp struct {
	Success bool
}

func handlePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	password := r.FormValue("password")
	ret := Resp{}
	ret.Success = false
	if password == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	} else if password == "12345" {
		ret.Success = true

	}
	json.NewEncoder(w).Encode(ret)

}

func main() {
	http.HandleFunc("/", handlePassword)
	http.ListenAndServe(":8000", nil)

}
