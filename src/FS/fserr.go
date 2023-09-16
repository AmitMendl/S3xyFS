package fs

type FSerr struct {
	message   string
	errorcode int
}

// compliance with Error interface
func (e FSerr) Error() string {
	return e.message
}
