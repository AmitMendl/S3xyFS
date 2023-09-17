package s3

type S3Response interface {
	ResponseXML() string
	Headers() map[string]string
}
