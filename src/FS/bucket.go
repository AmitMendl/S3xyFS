package fs

import (
	"io/fs"
	"os"
	"path"
)

// TODO: understand and implement bucket ownership (??)

// This is a bucket
type Bucket struct {
	name string
	path string
	acl  string
}

func initBucket(name string)

func (c *S3Controller) getBucketPath(bucket string) string {
	return path.Join(c.root, string(bucket))
}

func (c *S3Controller) CreateBucket(name string, acl string) *FSerr {
	if _, exists := c.buckets[name]; exists {
		return &FSerr{
			message:   ERR_BUCKET_ALREADY_EXISTS,
			errorcode: 400,
		}
	}

	path := path.Join(c.root, name)

	c.buckets[name] = Bucket{
		name: name,
		acl:  acl,
		path: path,
	}
	os.Mkdir(name, fs.FileMode(FS_FILEMODE))

	return nil
}
