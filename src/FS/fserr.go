package fs

type FSError struct {
	Code      string
	Message   string
	Errorcode int
}

// compliance with S3Error interface
func (e FSError) ErrCode() string {
	return e.Code
}

func (e FSError) ErrMessage() string {
	return e.Message
}
