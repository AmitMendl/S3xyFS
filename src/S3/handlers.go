package s3

/*
	This file describes a handler for each command in the S3 api, if a command is not here, i haven't coded it lol
*/

import (
	"net/http"

	fs "github.com/AmitMendl/S3xyFS/src/FS"
)

func initEndpoints(controller fs.S3Controller) {

}

type commandHandler struct {
	controller fs.S3Controller
}

// createBucket - mb command
type makeBucketHandler struct {
	commandHandler
}

func (h makeBucketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		return
	}

	uriParams := r.URL.Query()
	bucket := uriParams.Get(CB_URI_PARAM_BUCKET)
	acl := uriParams.Get(CB_URI_PARAM_ACL)

	h.controller.CreateBucket(bucket, acl)
}
