package s3

import (
	"encoding/xml"
	"log"
)

type S3Error interface {
	ErrCode() string
	ErrMessage() string
}

func GetXML(S3err S3Error) (string, error) {

	type s3err struct {
		XMLName xml.Name `xml:"Error"`
		Code    string   `xml:"Code"`
		Message string   `xml:"Message"`
	}

	xmlBody := s3err{
		Code:    S3err.ErrCode(),
		Message: S3err.ErrMessage(),
	}

	xml, err := xml.Marshal(
		xmlBody,
	)

	if err != nil {
		log.Fatal(err)
	}
	return string(xml), err
}
