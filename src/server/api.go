package server

import (
	"net/http"
)

func createBucket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" { // check HTTP method
		return // TODO: add error or something
	}

	
}
