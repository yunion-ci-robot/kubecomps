// Code generated by go-swagger; DO NOT EDIT.

package chart_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostChartrepoRepoChartsReader is a Reader for the PostChartrepoRepoCharts structure.
type PostChartrepoRepoChartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChartrepoRepoChartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostChartrepoRepoChartsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostChartrepoRepoChartsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostChartrepoRepoChartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostChartrepoRepoChartsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 507:
		result := NewPostChartrepoRepoChartsInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostChartrepoRepoChartsCreated creates a PostChartrepoRepoChartsCreated with default headers values
func NewPostChartrepoRepoChartsCreated() *PostChartrepoRepoChartsCreated {
	return &PostChartrepoRepoChartsCreated{}
}

/*
PostChartrepoRepoChartsCreated describes a response with status code 201, with default header values.

The specified chart is successfully uploaded.
*/
type PostChartrepoRepoChartsCreated struct {
}

// IsSuccess returns true when this post chartrepo repo charts created response has a 2xx status code
func (o *PostChartrepoRepoChartsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post chartrepo repo charts created response has a 3xx status code
func (o *PostChartrepoRepoChartsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo repo charts created response has a 4xx status code
func (o *PostChartrepoRepoChartsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post chartrepo repo charts created response has a 5xx status code
func (o *PostChartrepoRepoChartsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post chartrepo repo charts created response a status code equal to that given
func (o *PostChartrepoRepoChartsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostChartrepoRepoChartsCreated) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsCreated ", 201)
}

func (o *PostChartrepoRepoChartsCreated) String() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsCreated ", 201)
}

func (o *PostChartrepoRepoChartsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsUnauthorized creates a PostChartrepoRepoChartsUnauthorized with default headers values
func NewPostChartrepoRepoChartsUnauthorized() *PostChartrepoRepoChartsUnauthorized {
	return &PostChartrepoRepoChartsUnauthorized{}
}

/*
PostChartrepoRepoChartsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostChartrepoRepoChartsUnauthorized struct {
}

// IsSuccess returns true when this post chartrepo repo charts unauthorized response has a 2xx status code
func (o *PostChartrepoRepoChartsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo repo charts unauthorized response has a 3xx status code
func (o *PostChartrepoRepoChartsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo repo charts unauthorized response has a 4xx status code
func (o *PostChartrepoRepoChartsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post chartrepo repo charts unauthorized response has a 5xx status code
func (o *PostChartrepoRepoChartsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post chartrepo repo charts unauthorized response a status code equal to that given
func (o *PostChartrepoRepoChartsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostChartrepoRepoChartsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsUnauthorized ", 401)
}

func (o *PostChartrepoRepoChartsUnauthorized) String() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsUnauthorized ", 401)
}

func (o *PostChartrepoRepoChartsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsForbidden creates a PostChartrepoRepoChartsForbidden with default headers values
func NewPostChartrepoRepoChartsForbidden() *PostChartrepoRepoChartsForbidden {
	return &PostChartrepoRepoChartsForbidden{}
}

/*
PostChartrepoRepoChartsForbidden describes a response with status code 403, with default header values.

Operation is forbidden or quota exceeded
*/
type PostChartrepoRepoChartsForbidden struct {
}

// IsSuccess returns true when this post chartrepo repo charts forbidden response has a 2xx status code
func (o *PostChartrepoRepoChartsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo repo charts forbidden response has a 3xx status code
func (o *PostChartrepoRepoChartsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo repo charts forbidden response has a 4xx status code
func (o *PostChartrepoRepoChartsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post chartrepo repo charts forbidden response has a 5xx status code
func (o *PostChartrepoRepoChartsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post chartrepo repo charts forbidden response a status code equal to that given
func (o *PostChartrepoRepoChartsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostChartrepoRepoChartsForbidden) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsForbidden ", 403)
}

func (o *PostChartrepoRepoChartsForbidden) String() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsForbidden ", 403)
}

func (o *PostChartrepoRepoChartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsInternalServerError creates a PostChartrepoRepoChartsInternalServerError with default headers values
func NewPostChartrepoRepoChartsInternalServerError() *PostChartrepoRepoChartsInternalServerError {
	return &PostChartrepoRepoChartsInternalServerError{}
}

/*
PostChartrepoRepoChartsInternalServerError describes a response with status code 500, with default header values.

Internal server error occurred
*/
type PostChartrepoRepoChartsInternalServerError struct {
}

// IsSuccess returns true when this post chartrepo repo charts internal server error response has a 2xx status code
func (o *PostChartrepoRepoChartsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo repo charts internal server error response has a 3xx status code
func (o *PostChartrepoRepoChartsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo repo charts internal server error response has a 4xx status code
func (o *PostChartrepoRepoChartsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post chartrepo repo charts internal server error response has a 5xx status code
func (o *PostChartrepoRepoChartsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post chartrepo repo charts internal server error response a status code equal to that given
func (o *PostChartrepoRepoChartsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostChartrepoRepoChartsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsInternalServerError ", 500)
}

func (o *PostChartrepoRepoChartsInternalServerError) String() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsInternalServerError ", 500)
}

func (o *PostChartrepoRepoChartsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsInsufficientStorage creates a PostChartrepoRepoChartsInsufficientStorage with default headers values
func NewPostChartrepoRepoChartsInsufficientStorage() *PostChartrepoRepoChartsInsufficientStorage {
	return &PostChartrepoRepoChartsInsufficientStorage{}
}

/*
PostChartrepoRepoChartsInsufficientStorage describes a response with status code 507, with default header values.

Insufficient storage
*/
type PostChartrepoRepoChartsInsufficientStorage struct {
}

// IsSuccess returns true when this post chartrepo repo charts insufficient storage response has a 2xx status code
func (o *PostChartrepoRepoChartsInsufficientStorage) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post chartrepo repo charts insufficient storage response has a 3xx status code
func (o *PostChartrepoRepoChartsInsufficientStorage) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post chartrepo repo charts insufficient storage response has a 4xx status code
func (o *PostChartrepoRepoChartsInsufficientStorage) IsClientError() bool {
	return false
}

// IsServerError returns true when this post chartrepo repo charts insufficient storage response has a 5xx status code
func (o *PostChartrepoRepoChartsInsufficientStorage) IsServerError() bool {
	return true
}

// IsCode returns true when this post chartrepo repo charts insufficient storage response a status code equal to that given
func (o *PostChartrepoRepoChartsInsufficientStorage) IsCode(code int) bool {
	return code == 507
}

func (o *PostChartrepoRepoChartsInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsInsufficientStorage ", 507)
}

func (o *PostChartrepoRepoChartsInsufficientStorage) String() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsInsufficientStorage ", 507)
}

func (o *PostChartrepoRepoChartsInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
