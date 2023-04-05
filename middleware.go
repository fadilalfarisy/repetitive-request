package main

import (
	"net/http"
)

const USERNAME = "batman"
const PASSWORD = "secret"

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
