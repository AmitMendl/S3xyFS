package s3

/*
	This file describes a handler for each command in the S3 api, if a command is not here, i haven't coded it lol
*/

import (
	"io"
	"net/http"
	"strings"

	s3fs "github.com/AmitMendl/S3xyFS/src/FS"
)

type CommandHandler struct {
	Controller s3fs.S3Controller
}

// createBucket - mb command

type S3Handler struct {
	*CommandHandler
}

func (h S3Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		CreateBucketConfiguration struct {
			LocationConstraint string
		}
	}

	if r.Method != "PUT" {
		return
	}

}

func (h S3Handler) handlePut(w http.ResponseWriter, r *http.Request) *S3Error {

	uriParams := r.URL.Query()
	pathParams := strings.Split(r.URL.Path, "/")

	acl := uriParams.Get(CB_URI_PARAM_ACL)

	bucket := pathParams[1]

	objectPath := r.URL.Path
	object, _ := io.ReadAll(r.Body)

	var s3err S3Error = nil
	if r.Body == http.NoBody {
		s3err = h.Controller.CreateBucket(bucket, acl)
	} else {
		s3err = h.Controller.PutObject(object, objectPath, acl)
	}

	if s3err != nil {
		RespBody, _ := GetXML(s3err)
		http.Error(w, RespBody, s3err.ErrHttpCode())
	}

	return nil
}
