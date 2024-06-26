// Code generated by go-swagger; DO NOT EDIT.

package chart_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostChartrepoChartsReader is a Reader for the PostChartrepoCharts structure.
type PostChartrepoChartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChartrepoChartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostChartrepoChartsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostChartrepoChartsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostChartrepoChartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostChartrepoChartsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 507:
		result := NewPostChartrepoChartsInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostChartrepoChartsCreated creates a PostChartrepoChartsCreated with default headers values
func NewPostChartrepoChartsCreated() *PostChartrepoChartsCreated {
	return &PostChartrepoChartsCreated{}
}

/*
PostChartrepoChartsCreated describes a response with status code 201, with default header values.

The specified chart is successfully uploaded.
*/
type PostChartrepoChartsCreated struct {
}

// IsSuccess returns true when this post chartrepo charts created response has a 2xx status code
func (o *PostChartrepoChartsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post chartrepo charts created response has a 3xx status code
func (o *PostChartrepoChartsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo charts created response has a 4xx status code
func (o *PostChartrepoChartsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post chartrepo charts created response has a 5xx status code
func (o *PostChartrepoChartsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post chartrepo charts created response a status code equal to that given
func (o *PostChartrepoChartsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostChartrepoChartsCreated) Error() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsCreated ", 201)
}

func (o *PostChartrepoChartsCreated) String() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsCreated ", 201)
}

func (o *PostChartrepoChartsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoChartsUnauthorized creates a PostChartrepoChartsUnauthorized with default headers values
func NewPostChartrepoChartsUnauthorized() *PostChartrepoChartsUnauthorized {
	return &PostChartrepoChartsUnauthorized{}
}

/*
PostChartrepoChartsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostChartrepoChartsUnauthorized struct {
}

// IsSuccess returns true when this post chartrepo charts unauthorized response has a 2xx status code
func (o *PostChartrepoChartsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo charts unauthorized response has a 3xx status code
func (o *PostChartrepoChartsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo charts unauthorized response has a 4xx status code
func (o *PostChartrepoChartsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post chartrepo charts unauthorized response has a 5xx status code
func (o *PostChartrepoChartsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post chartrepo charts unauthorized response a status code equal to that given
func (o *PostChartrepoChartsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostChartrepoChartsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsUnauthorized ", 401)
}

func (o *PostChartrepoChartsUnauthorized) String() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsUnauthorized ", 401)
}

func (o *PostChartrepoChartsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoChartsForbidden creates a PostChartrepoChartsForbidden with default headers values
func NewPostChartrepoChartsForbidden() *PostChartrepoChartsForbidden {
	return &PostChartrepoChartsForbidden{}
}

/*
PostChartrepoChartsForbidden describes a response with status code 403, with default header values.

Operation is forbidden or quota exceeded
*/
type PostChartrepoChartsForbidden struct {
}

// IsSuccess returns true when this post chartrepo charts forbidden response has a 2xx status code
func (o *PostChartrepoChartsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo charts forbidden response has a 3xx status code
func (o *PostChartrepoChartsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo charts forbidden response has a 4xx status code
func (o *PostChartrepoChartsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post chartrepo charts forbidden response has a 5xx status code
func (o *PostChartrepoChartsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post chartrepo charts forbidden response a status code equal to that given
func (o *PostChartrepoChartsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostChartrepoChartsForbidden) Error() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsForbidden ", 403)
}

func (o *PostChartrepoChartsForbidden) String() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsForbidden ", 403)
}

func (o *PostChartrepoChartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoChartsInternalServerError creates a PostChartrepoChartsInternalServerError with default headers values
func NewPostChartrepoChartsInternalServerError() *PostChartrepoChartsInternalServerError {
	return &PostChartrepoChartsInternalServerError{}
}

/*
PostChartrepoChartsInternalServerError describes a response with status code 500, with default header values.

Internal server error occurred
*/
type PostChartrepoChartsInternalServerError struct {
}

// IsSuccess returns true when this post chartrepo charts internal server error response has a 2xx status code
func (o *PostChartrepoChartsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo charts internal server error response has a 3xx status code
func (o *PostChartrepoChartsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo charts internal server error response has a 4xx status code
func (o *PostChartrepoChartsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post chartrepo charts internal server error response has a 5xx status code
func (o *PostChartrepoChartsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post chartrepo charts internal server error response a status code equal to that given
func (o *PostChartrepoChartsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostChartrepoChartsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsInternalServerError ", 500)
}

func (o *PostChartrepoChartsInternalServerError) String() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsInternalServerError ", 500)
}

func (o *PostChartrepoChartsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoChartsInsufficientStorage creates a PostChartrepoChartsInsufficientStorage with default headers values
func NewPostChartrepoChartsInsufficientStorage() *PostChartrepoChartsInsufficientStorage {
	return &PostChartrepoChartsInsufficientStorage{}
}

/*
PostChartrepoChartsInsufficientStorage describes a response with status code 507, with default header values.

Insufficient storage
*/
type PostChartrepoChartsInsufficientStorage struct {
}

// IsSuccess returns true when this post chartrepo charts insufficient storage response has a 2xx status code
func (o *PostChartrepoChartsInsufficientStorage) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo charts insufficient storage response has a 3xx status code
func (o *PostChartrepoChartsInsufficientStorage) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo charts insufficient storage response has a 4xx status code
func (o *PostChartrepoChartsInsufficientStorage) IsClientError() bool {
	return false
}

// IsServerError returns true when this post chartrepo charts insufficient storage response has a 5xx status code
func (o *PostChartrepoChartsInsufficientStorage) IsServerError() bool {
	return true
}

// IsCode returns true when this post chartrepo charts insufficient storage response a status code equal to that given
func (o *PostChartrepoChartsInsufficientStorage) IsCode(code int) bool {
	return code == 507
}

func (o *PostChartrepoChartsInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsInsufficientStorage ", 507)
}

func (o *PostChartrepoChartsInsufficientStorage) String() string {
	return fmt.Sprintf("[POST /chartrepo/charts][%d] postChartrepoChartsInsufficientStorage ", 507)
}

func (o *PostChartrepoChartsInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
