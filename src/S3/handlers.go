package s3

/*
	This file describes a handler for each command in the S3 api, if a command is not here, i haven't coded it lol
*/

import (
	"net/http"
	"strings"

	s3fs "github.com/AmitMendl/S3xyFS/src/FS"
)

type CommandHandler struct {
	Controller s3fs.S3Controller
}

// createBucket - mb command
type CreateBucketHandler struct {
	*CommandHandler
}

func (h CreateBucketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		CreateBucketConfiguration struct {
			LocationConstraint string
		}
	}

	if r.Method != "PUT" {
		return
	}

	uriParams := r.URL.Query()
	pathParams := strings.Split(r.URL.Path, "/")

	bucket := pathParams[1]
	acl := uriParams.Get(CB_URI_PARAM_ACL)
	// body doesn't matter

	fserr := h.Controller.CreateBucket(bucket, acl)
	if fserr != nil {
		RespBody, _ := GetXML(fserr)
		http.Error(w, RespBody, fserr.Errorcode)
	}

}
