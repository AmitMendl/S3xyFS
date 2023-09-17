package s3

// ACL

const (
	ACL_PRIVATE            string = "private"
	ACL_PUBLIC_READ               = "public-read"
	ACL_PUBLIC_READ_WRITE         = "public-read-write"
	ACL_AUTHENTICATED_READ        = "authenticated-read"
)

// COMMAND ENUM

type COMMAND int64

const (
	CMD_NOT_FOUND COMMAND = iota
	CMD_PUT_CREATE_BUCKET
	CMD_PUT_CREATE_OBJECT
)

// COMMAND CONSTS
const (
	CB_URI_PARAM_ACL string = "x-amz-acl"
)
