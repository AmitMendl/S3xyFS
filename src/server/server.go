package server

import (
	"log"
	"net/http"

	s3 "github.com/AmitMendl/S3xyFS/src/S3"
)

type S3rver struct {
	server     http.Server
	controller s3.S3Controller
}

func StartServer() {
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
