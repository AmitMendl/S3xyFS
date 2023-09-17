package fs

import (
	"io/fs"
	"os"
	"path"
)

type Object struct {
	object string
	path   string
	acl    string
}

func (c *S3Controller) getObjectPath(objectPath string) string {
	return path.Join(c.root, objectPath)
}

func (c *S3Controller) PutObject(object []byte, objectPath string, acl string) *FSError {

	os.WriteFile(c.getObjectPath(objectPath), object, fs.FileMode(FS_FILEMODE))

	return nil
}
