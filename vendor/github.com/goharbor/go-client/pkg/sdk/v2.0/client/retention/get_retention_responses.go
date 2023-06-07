// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetRetentionReader is a Reader for the GetRetention structure.
type GetRetentionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRetentionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRetentionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRetentionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRetentionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRetentionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRetentionOK creates a GetRetentionOK with default headers values
func NewGetRetentionOK() *GetRetentionOK {
	return &GetRetentionOK{}
}

/*
GetRetentionOK describes a response with status code 200, with default header values.

Get Retention Policy successfully.
*/
type GetRetentionOK struct {
	Payload *models.RetentionPolicy
}

// IsSuccess returns true when this get retention o k response has a 2xx status code
func (o *GetRetentionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get retention o k response has a 3xx status code
func (o *GetRetentionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get retention o k response has a 4xx status code
func (o *GetRetentionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get retention o k response has a 5xx status code
func (o *GetRetentionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get retention o k response a status code equal to that given
func (o *GetRetentionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRetentionOK) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionOK  %+v", 200, o.Payload)
}

func (o *GetRetentionOK) String() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionOK  %+v", 200, o.Payload)
}

func (o *GetRetentionOK) GetPayload() *models.RetentionPolicy {
	return o.Payload
}

func (o *GetRetentionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetentionPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRetentionUnauthorized creates a GetRetentionUnauthorized with default headers values
func NewGetRetentionUnauthorized() *GetRetentionUnauthorized {
	return &GetRetentionUnauthorized{}
}

/*
GetRetentionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRetentionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get retention unauthorized response has a 2xx status code
func (o *GetRetentionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get retention unauthorized response has a 3xx status code
func (o *GetRetentionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get retention unauthorized response has a 4xx status code
func (o *GetRetentionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get retention unauthorized response has a 5xx status code
func (o *GetRetentionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get retention unauthorized response a status code equal to that given
func (o *GetRetentionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRetentionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRetentionUnauthorized) String() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRetentionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRetentionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRetentionForbidden creates a GetRetentionForbidden with default headers values
func NewGetRetentionForbidden() *GetRetentionForbidden {
	return &GetRetentionForbidden{}
}

/*
GetRetentionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRetentionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get retention forbidden response has a 2xx status code
func (o *GetRetentionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get retention forbidden response has a 3xx status code
func (o *GetRetentionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get retention forbidden response has a 4xx status code
func (o *GetRetentionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get retention forbidden response has a 5xx status code
func (o *GetRetentionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get retention forbidden response a status code equal to that given
func (o *GetRetentionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRetentionForbidden) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionForbidden  %+v", 403, o.Payload)
}

func (o *GetRetentionForbidden) String() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionForbidden  %+v", 403, o.Payload)
}

func (o *GetRetentionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRetentionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRetentionInternalServerError creates a GetRetentionInternalServerError with default headers values
func NewGetRetentionInternalServerError() *GetRetentionInternalServerError {
	return &GetRetentionInternalServerError{}
}

/*
GetRetentionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRetentionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get retention internal server error response has a 2xx status code
func (o *GetRetentionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get retention internal server error response has a 3xx status code
func (o *GetRetentionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get retention internal server error response has a 4xx status code
func (o *GetRetentionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get retention internal server error response has a 5xx status code
func (o *GetRetentionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get retention internal server error response a status code equal to that given
func (o *GetRetentionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRetentionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRetentionInternalServerError) String() string {
	return fmt.Sprintf("[GET /retentions/{id}][%d] getRetentionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRetentionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRetentionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
