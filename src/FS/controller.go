package fs

import (
	"path"
)

type S3Controller struct {
	root    string
	buckets map[string]Bucket
}

func (c *S3Controller) InitController(path string) {
	c.root = path
	c.buckets = make(map[string]Bucket)
}

func (c *S3Controller) getBucketPath(bucket string) string {
	return path.Join(c.root, string(bucket))
}
