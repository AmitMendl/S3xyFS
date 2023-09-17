package server

import (
	"fmt"
	"net/http"

	s3fs "github.com/AmitMendl/S3xyFS/src/FS"
	s3 "github.com/AmitMendl/S3xyFS/src/S3"
)

func StartServer(config S3rverConfig) {

	// init controller
	controller := s3fs.InitController(config.Root)

	// Start handler
	Handler := s3.S3Handler{
		Controller: *controller,
	}
	http.Handle("/", Handler)

	// TODO: Proper interraction with config
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
