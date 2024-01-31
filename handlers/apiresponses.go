package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

var (
	MissingField = NewOutput("missing_field", "A required field is missing.")
	NotFound     = NewOutput("not_found", "What you are looking for cannot be found.")
	BadRequest   = NewOutput("bad_request", "The request is invalid.")
	Invalid      = NewOutput("invalid", "The value provided is invalid.")
	Internal     = NewOutput("internal_error", "An internal error occurred.")
)

type Output struct {
	Value string
	Desc  string
}

func NewOutput(value, desc string) *Output {
	return &Output{
		Value: value,
		Desc:  desc,
	}
}

func (o Output) String() string {
	return o.Value
}

func (o Output) ToUpper() string {
	return strings.ToUpper(o.String())
}

func (o Output) Description() string {
	return o.Desc
}

type FieldError struct {
	Field     string `json:"field,omitempty"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message,omitempty"`
}

type Error struct {
	Message     string        `json:"message"`
	Resource    string        `json:"resource"`
	Description string        `json:"description"`
	Errors      []*FieldError `json:"errors,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

func write(response events.APIGatewayProxyResponse, data interface{}) events.APIGatewayProxyResponse {
	switch {
	case response.StatusCode < 200:
		panic(fmt.Sprintf("status code %d must be >= 200", response.StatusCode))
	case response.StatusCode == 204:
		return response
	}

	response.Headers["Content-Type"] = "application/json"

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			fields := log.Fields{"data": data, "code": response.StatusCode}
			log.WithFields(fields).Errorf("%+v", err)
			response.StatusCode = http.StatusInternalServerError
			response.Body = "Internal Server Error"
			return response
		}
		response.Body = string(jsonData)
	}

	return response
}

func Err(response events.APIGatewayProxyResponse, e *Error, code int) events.APIGatewayProxyResponse {
	if e == nil {
		panic("error must not be empty")
	}
	response.StatusCode = code
	return write(response, e)
}

func (e *Error) AddInvalidError(field string) {
	e.Errors = append(e.Errors, &FieldError{
		ErrorCode: Invalid.String(),
		Field:     field,
	})
}

func New(message, resource string, output *Output, e ...*FieldError) *Error {
	err := &Error{
		Message:     message,
		Resource:    resource,
		Description: output.Description(),
	}
	if len(e) != 0 && e[0] != nil {
		err.Errors = e
	}
	return err
}

func NewInvalidError(message, resource, field string) *Error {
	return New(message, resource, Invalid, &FieldError{
		Field:     field,
		ErrorCode: Invalid.String(),
	})
}

func NewMissingFieldError(message, resource, field string) *Error {
	return New(message, resource, MissingField, &FieldError{
		Field:     field,
		ErrorCode: MissingField.String(),
	})
}

func NewInternalError(resource string) *Error {
	return New(Internal.ToUpper(), resource, Internal, nil)
}

func NewNotFoundError(resource string) *Error {
	return New(NotFound.ToUpper(), resource, NotFound, nil)
}

func OK200(response events.APIGatewayProxyResponse, data interface{}) events.APIGatewayProxyResponse {
	response.StatusCode = 200
	return write(response, data)
}

func Created201(response events.APIGatewayProxyResponse, data interface{}) events.APIGatewayProxyResponse {
	response.StatusCode = 201
	return write(response, data)
}

func BadRequest400(response events.APIGatewayProxyResponse, resource, field string) events.APIGatewayProxyResponse {
	return Err(response, NewInvalidError("BAD_REQUEST", resource, field), 400)
}

func NotFound404(response events.APIGatewayProxyResponse, resource string) events.APIGatewayProxyResponse {
	return Err(response, NewNotFoundError(resource), 404)
}

func InternalError500(response events.APIGatewayProxyResponse, resource string, err error) events.APIGatewayProxyResponse {
	return Err(response, NewInternalError(resource), 500)
}
