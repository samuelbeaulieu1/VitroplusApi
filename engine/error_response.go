package engine

import "strings"

type ErrorResponse struct {
	*RequestResponse
}

type FieldErrorResponse struct {
	*ErrorResponse
	Body   []string `json:"body"`
	Fields []string `json:"fields"`
}

type Error interface {
	Error() string
}

func NewError(message string) *ErrorResponse {
	return &ErrorResponse{
		&RequestResponse{
			Title: "Erreur",
			Body:  message,
		},
	}
}

func (err *ErrorResponse) Error() string {
	return err.Body
}

func NewFieldsError(messages []string, fields []string) *FieldErrorResponse {
	err := &ErrorResponse{
		&RequestResponse{
			Title: "Erreur",
		},
	}
	fieldErr := &FieldErrorResponse{
		Fields: fields,
		Body:   messages,
	}
	fieldErr.ErrorResponse = err

	return fieldErr
}

func (err *FieldErrorResponse) Error() string {
	return strings.Join(err.Body, "\n")
}
