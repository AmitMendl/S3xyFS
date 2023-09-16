package server

import (
	"fmt"
	"net/http"

	s3 "github.com/AmitMendl/S3xyFS/src/S3"
)

func StartServer(config S3rverConfig) {

	// init controller
	

	// Start handler
	http.Handle("/", new(s3.MakeBucketHandler))

	// TODO: Proper interraction with config
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
