package fs

// FS config
const (
	FS_FILEMODE int = 0700
)

// ERRORS
const (
	ERR_BUCKET_ALREADY_EXISTS       string = "The requested bucket name is not available. The bucket namespace is shared by all users of the system. Select a different name and try again."
	ERR_BUCKET_ALREADY_OWNED_BY_YOU string = "The bucket you tried to create already exists, and you own it."
)
