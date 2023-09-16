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
	cbHandler := s3.CreateBucketHandler{
		CommandHandler: &s3.CommandHandler{
			Controller: *controller,
		},
	}
	http.Handle("/", cbHandler)

	// TODO: Proper interraction with config
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
