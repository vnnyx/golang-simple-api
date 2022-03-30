package exception

type NotFoundError struct {
	Error string
}

func NewErrorNotFound(error string) NotFoundError {
	return NotFoundError{
		Error: error,
	}
}
