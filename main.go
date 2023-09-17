package main

import (
	server "github.com/AmitMendl/S3xyFS/src/server"
)

func main() {

	config := server.S3rverConfig{
		Port: 8080,
		Root: "/tmp/t3st/",
	}

	server.StartServer(config)
}
