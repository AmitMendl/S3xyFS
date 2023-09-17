package fs

type FSError struct {
	Code      string
	Message   string
	Errorcode int
}

// compliance with Error interface
func (e FSError) Error() string {
	return e.Message
}
