package s3

import (
	"errors"
	"os"
	"path"
)

type S3Controller struct {
	root    string
	buckets map[string]Bucket
}

func (c *S3Controller) AddBucket(name string) (err error) {
	if _, isIn := c.buckets[name]; isIn {
		return errors.New("Bucket already exists") // TODO: Check actual error message
	}

	bucket := Bucket{
		name: name,
		path: c.getBucketPath(name),
	}

	if err := os.Mkdir(bucket.path, 0700); err != nil {
		return err
	}

	c.buckets[name] = bucket

	return nil
}

func (c *S3Controller) getBucketPath(bucket string) string {
	return path.Join(c.root, string(bucket))
}
