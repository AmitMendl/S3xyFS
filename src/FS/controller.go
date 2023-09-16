package fs

type S3Controller struct {
	root    string
	buckets map[string]Bucket
}

func InitController(root string) *S3Controller {
	return &S3Controller{
		root:    root,
		buckets: make(map[string]Bucket),
	}
}
