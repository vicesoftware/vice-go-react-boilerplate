package main

type invalidRequest struct {
	message string
}

func (e *invalidRequest) Error() string {
	if e.message != "" {
		return e.message
	}
	return "invalid request"
}

type notFound struct {
	message string
}

func (e *notFound) Error() string {
	if e.message != "" {
		return e.message
	}
	return "not found"
}
