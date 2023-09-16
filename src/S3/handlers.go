package s3

/*
	This file describes a handler for each command in the S3 api, if a command is not here, i haven't coded it lol
*/

import (
	"net/http"
)

func initEndpoints(controller S3Controller) {

}

type commandHandler struct {
	controller S3Controller
}

// createBucket - mb command
type makeBucketHandler struct {
	commandHandler
}

func (h makeBucketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		return
	}
}
