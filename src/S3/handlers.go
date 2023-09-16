package s3

/*
	This file describes a handler for each command in the S3 api, if a command is not here, i haven't coded it lol
*/

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	fs "github.com/AmitMendl/S3xyFS/src/FS"
)

type commandHandler struct {
	controller fs.S3Controller
}

// createBucket - mb command
type makeBucketHandler struct {
	commandHandler
}

func (h makeBucketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		CreateBucketConfiguration struct {
			LocationConstraint string
		}
	}

	if r.Method != "PUT" {
		return
	}

	uriParams := r.URL.Query()
	bucket := uriParams.Get(CB_URI_PARAM_BUCKET)
	acl := uriParams.Get(CB_URI_PARAM_ACL)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var body Body
	if err := xml.Unmarshal(reqBody, &body); err != nil {
		// TODO: Something something error 422
		return
	}

	h.controller.CreateBucket(bucket, acl)
}
