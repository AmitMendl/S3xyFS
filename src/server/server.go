package server

import (
	"fmt"
	"net/http"
)

func StartServer(config S3rverConfig) {
	// TODO: Proper interraction with config
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
