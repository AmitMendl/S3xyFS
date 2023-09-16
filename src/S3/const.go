package s3

// ACL

const (
	ACL_PRIVATE            string = "private"
	ACL_PUBLIC_READ               = "public-read"
	ACL_PUBLIC_READ_WRITE         = "public-read-write"
	ACL_AUTHENTICATED_READ        = "authenticated-read"
)

// COMMAND CONSTS

// createBucket
const (
	CB_URI_PARAM_BUCKET string = "BUCKET"
	CB_URI_PARAM_ACL    string = "x-amz-acl"
)
