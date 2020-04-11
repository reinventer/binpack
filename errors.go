package binpack

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrCannotFit Error = "rectangle has not fit"
)
