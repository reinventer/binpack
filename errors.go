package binpack

// Error represents error in this package
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrCannotFit error then rectangle cannot fit in bin
	ErrCannotFit Error = "rectangle has not fit"
)
