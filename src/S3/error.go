package s3

import (
	"encoding/xml"

	s3fs "github.com/AmitMendl/S3xyFS/src/FS"
)

type S3Error struct {
	Code      string
	Message   string
	Resource  string
	RequestID int
}

func something(fserr s3fs.FSError) (s3err S3Error) {
	return S3Error{
		Code:    fserr.Code,
		Message: fserr.Message,
	}
}

func (s3err *S3Error) GetXML() ([]byte, error) {
	xml, err := xml.Marshal(
		map[string]S3Error{
			"Error": *s3err,
		},
	)
	return xml, err
}
