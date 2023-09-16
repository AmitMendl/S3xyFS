package server

import (
	"fmt"
	"net/http"

	s3 "github.com/AmitMendl/S3xyFS/src/S3"
)

type S3rver struct {
	controller s3.S3Controller
}

func StartServer(config S3rverConfig) {
	// TODO: Proper interraction with config
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
