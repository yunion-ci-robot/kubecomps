// Code generated by go-swagger; DO NOT EDIT.

package statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetStatisticReader is a Reader for the GetStatistic structure.
type GetStatisticReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatisticReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatisticOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStatisticUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStatisticInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatisticOK creates a GetStatisticOK with default headers values
func NewGetStatisticOK() *GetStatisticOK {
	return &GetStatisticOK{}
}

/*
GetStatisticOK describes a response with status code 200, with default header values.

The statistic information
*/
type GetStatisticOK struct {
	Payload *models.Statistic
}

// IsSuccess returns true when this get statistic o k response has a 2xx status code
func (o *GetStatisticOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get statistic o k response has a 3xx status code
func (o *GetStatisticOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get statistic o k response has a 4xx status code
func (o *GetStatisticOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get statistic o k response has a 5xx status code
func (o *GetStatisticOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get statistic o k response a status code equal to that given
func (o *GetStatisticOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetStatisticOK) Error() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticOK  %+v", 200, o.Payload)
}

func (o *GetStatisticOK) String() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticOK  %+v", 200, o.Payload)
}

func (o *GetStatisticOK) GetPayload() *models.Statistic {
	return o.Payload
}

func (o *GetStatisticOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Statistic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatisticUnauthorized creates a GetStatisticUnauthorized with default headers values
func NewGetStatisticUnauthorized() *GetStatisticUnauthorized {
	return &GetStatisticUnauthorized{}
}

/*
GetStatisticUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetStatisticUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get statistic unauthorized response has a 2xx status code
func (o *GetStatisticUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get statistic unauthorized response has a 3xx status code
func (o *GetStatisticUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get statistic unauthorized response has a 4xx status code
func (o *GetStatisticUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get statistic unauthorized response has a 5xx status code
func (o *GetStatisticUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get statistic unauthorized response a status code equal to that given
func (o *GetStatisticUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetStatisticUnauthorized) Error() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStatisticUnauthorized) String() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStatisticUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetStatisticUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatisticInternalServerError creates a GetStatisticInternalServerError with default headers values
func NewGetStatisticInternalServerError() *GetStatisticInternalServerError {
	return &GetStatisticInternalServerError{}
}

/*
GetStatisticInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetStatisticInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get statistic internal server error response has a 2xx status code
func (o *GetStatisticInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get statistic internal server error response has a 3xx status code
func (o *GetStatisticInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get statistic internal server error response has a 4xx status code
func (o *GetStatisticInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get statistic internal server error response has a 5xx status code
func (o *GetStatisticInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get statistic internal server error response a status code equal to that given
func (o *GetStatisticInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetStatisticInternalServerError) Error() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStatisticInternalServerError) String() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStatisticInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetStatisticInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
