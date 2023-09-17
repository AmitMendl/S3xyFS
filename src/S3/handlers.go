package s3

/*
	This file describes a handler for each command in the S3 api, if a command is not here, i haven't coded it lol
*/

import (
	"net/http"
	"strings"

	s3fs "github.com/AmitMendl/S3xyFS/src/FS"
)

type S3Handler struct {
	Controller s3fs.S3Controller
}

func (h S3Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		CreateBucketConfiguration struct {
			LocationConstraint string
		}
	}

	if r.Method == "PUT" {
		h.handlePut(w, r)
	}
}

func (h S3Handler) handlePut(w http.ResponseWriter, r *http.Request) {

	uriParams := r.URL.Query()
	pathParams := strings.Split(r.URL.Path, "/")

	bucket := pathParams[1]
	acl := uriParams.Get(CB_URI_PARAM_ACL)

	fserr := h.Controller.CreateBucket(bucket, acl)
	if fserr != nil {
		RespBody, _ := GetXML(fserr)
		http.Error(w, RespBody, fserr.Errorcode)
	}
}
