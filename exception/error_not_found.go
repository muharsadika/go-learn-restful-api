package exception

type ErrorNotFound struct {
	Error string `json:"error"`
}

func NewErrorNotFound(error string) ErrorNotFound {
	ErrorNotFound := ErrorNotFound{
		Error: error,
	}

	return ErrorNotFound
}
