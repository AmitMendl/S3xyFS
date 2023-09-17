package fs

type FSerr struct {
	Message   string
	Errorcode int
}

// compliance with Error interface
func (e FSerr) Error() string {
	return e.Message
}
