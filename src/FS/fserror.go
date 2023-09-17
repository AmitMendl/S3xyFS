package fs

type FSError struct {
	HttpCode  int
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

func (e FSError) ErrHttpCode() int {
	return e.ErrHttpCode()
}
