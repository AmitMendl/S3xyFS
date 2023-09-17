package fs

import (
	"log"
	"os"
)

type S3Controller struct {
	root    string
	buckets map[string]Bucket
}

func InitController(root string) *S3Controller {

	controller := &S3Controller{
		root:    root,
		buckets: make(map[string]Bucket),
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			controller.InitBucket(e.Name(), "") // TODO: add ACL implementation for each bucket
		}
	}

	return controller
}
