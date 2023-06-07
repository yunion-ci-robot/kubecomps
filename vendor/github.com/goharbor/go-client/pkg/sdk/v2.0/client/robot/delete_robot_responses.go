// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// DeleteRobotReader is a Reader for the DeleteRobot structure.
type DeleteRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRobotOK creates a DeleteRobotOK with default headers values
func NewDeleteRobotOK() *DeleteRobotOK {
	return &DeleteRobotOK{}
}

/*
DeleteRobotOK describes a response with status code 200, with default header values.

Success
*/
type DeleteRobotOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete robot o k response has a 2xx status code
func (o *DeleteRobotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete robot o k response has a 3xx status code
func (o *DeleteRobotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot o k response has a 4xx status code
func (o *DeleteRobotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete robot o k response has a 5xx status code
func (o *DeleteRobotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot o k response a status code equal to that given
func (o *DeleteRobotOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRobotOK) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotOK ", 200)
}

func (o *DeleteRobotOK) String() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotOK ", 200)
}

func (o *DeleteRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteRobotBadRequest creates a DeleteRobotBadRequest with default headers values
func NewDeleteRobotBadRequest() *DeleteRobotBadRequest {
	return &DeleteRobotBadRequest{}
}

/*
DeleteRobotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteRobotBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot bad request response has a 2xx status code
func (o *DeleteRobotBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot bad request response has a 3xx status code
func (o *DeleteRobotBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot bad request response has a 4xx status code
func (o *DeleteRobotBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot bad request response has a 5xx status code
func (o *DeleteRobotBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot bad request response a status code equal to that given
func (o *DeleteRobotBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRobotBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRobotBadRequest) String() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRobotBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotUnauthorized creates a DeleteRobotUnauthorized with default headers values
func NewDeleteRobotUnauthorized() *DeleteRobotUnauthorized {
	return &DeleteRobotUnauthorized{}
}

/*
DeleteRobotUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRobotUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot unauthorized response has a 2xx status code
func (o *DeleteRobotUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot unauthorized response has a 3xx status code
func (o *DeleteRobotUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot unauthorized response has a 4xx status code
func (o *DeleteRobotUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot unauthorized response has a 5xx status code
func (o *DeleteRobotUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot unauthorized response a status code equal to that given
func (o *DeleteRobotUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRobotUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRobotUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRobotUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotForbidden creates a DeleteRobotForbidden with default headers values
func NewDeleteRobotForbidden() *DeleteRobotForbidden {
	return &DeleteRobotForbidden{}
}

/*
DeleteRobotForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRobotForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot forbidden response has a 2xx status code
func (o *DeleteRobotForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot forbidden response has a 3xx status code
func (o *DeleteRobotForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot forbidden response has a 4xx status code
func (o *DeleteRobotForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot forbidden response has a 5xx status code
func (o *DeleteRobotForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot forbidden response a status code equal to that given
func (o *DeleteRobotForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRobotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRobotForbidden) String() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRobotForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotNotFound creates a DeleteRobotNotFound with default headers values
func NewDeleteRobotNotFound() *DeleteRobotNotFound {
	return &DeleteRobotNotFound{}
}

/*
DeleteRobotNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteRobotNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot not found response has a 2xx status code
func (o *DeleteRobotNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot not found response has a 3xx status code
func (o *DeleteRobotNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot not found response has a 4xx status code
func (o *DeleteRobotNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot not found response has a 5xx status code
func (o *DeleteRobotNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot not found response a status code equal to that given
func (o *DeleteRobotNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRobotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRobotNotFound) String() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRobotNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotInternalServerError creates a DeleteRobotInternalServerError with default headers values
func NewDeleteRobotInternalServerError() *DeleteRobotInternalServerError {
	return &DeleteRobotInternalServerError{}
}

/*
DeleteRobotInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteRobotInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot internal server error response has a 2xx status code
func (o *DeleteRobotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot internal server error response has a 3xx status code
func (o *DeleteRobotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot internal server error response has a 4xx status code
func (o *DeleteRobotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete robot internal server error response has a 5xx status code
func (o *DeleteRobotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete robot internal server error response a status code equal to that given
func (o *DeleteRobotInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRobotInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRobotInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRobotInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
