// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// ListTasksOfWebhookExecutionReader is a Reader for the ListTasksOfWebhookExecution structure.
type ListTasksOfWebhookExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksOfWebhookExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTasksOfWebhookExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTasksOfWebhookExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListTasksOfWebhookExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTasksOfWebhookExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTasksOfWebhookExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTasksOfWebhookExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTasksOfWebhookExecutionOK creates a ListTasksOfWebhookExecutionOK with default headers values
func NewListTasksOfWebhookExecutionOK() *ListTasksOfWebhookExecutionOK {
	return &ListTasksOfWebhookExecutionOK{}
}

/*
ListTasksOfWebhookExecutionOK describes a response with status code 200, with default header values.

List tasks of webhook executions success
*/
type ListTasksOfWebhookExecutionOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of tasks
	 */
	XTotalCount int64

	Payload []*models.Task
}

// IsSuccess returns true when this list tasks of webhook execution o k response has a 2xx status code
func (o *ListTasksOfWebhookExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tasks of webhook execution o k response has a 3xx status code
func (o *ListTasksOfWebhookExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks of webhook execution o k response has a 4xx status code
func (o *ListTasksOfWebhookExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tasks of webhook execution o k response has a 5xx status code
func (o *ListTasksOfWebhookExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks of webhook execution o k response a status code equal to that given
func (o *ListTasksOfWebhookExecutionOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListTasksOfWebhookExecutionOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionOK  %+v", 200, o.Payload)
}

func (o *ListTasksOfWebhookExecutionOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionOK  %+v", 200, o.Payload)
}

func (o *ListTasksOfWebhookExecutionOK) GetPayload() []*models.Task {
	return o.Payload
}

func (o *ListTasksOfWebhookExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksOfWebhookExecutionBadRequest creates a ListTasksOfWebhookExecutionBadRequest with default headers values
func NewListTasksOfWebhookExecutionBadRequest() *ListTasksOfWebhookExecutionBadRequest {
	return &ListTasksOfWebhookExecutionBadRequest{}
}

/*
ListTasksOfWebhookExecutionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListTasksOfWebhookExecutionBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list tasks of webhook execution bad request response has a 2xx status code
func (o *ListTasksOfWebhookExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks of webhook execution bad request response has a 3xx status code
func (o *ListTasksOfWebhookExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks of webhook execution bad request response has a 4xx status code
func (o *ListTasksOfWebhookExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tasks of webhook execution bad request response has a 5xx status code
func (o *ListTasksOfWebhookExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks of webhook execution bad request response a status code equal to that given
func (o *ListTasksOfWebhookExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListTasksOfWebhookExecutionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *ListTasksOfWebhookExecutionBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *ListTasksOfWebhookExecutionBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListTasksOfWebhookExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTasksOfWebhookExecutionUnauthorized creates a ListTasksOfWebhookExecutionUnauthorized with default headers values
func NewListTasksOfWebhookExecutionUnauthorized() *ListTasksOfWebhookExecutionUnauthorized {
	return &ListTasksOfWebhookExecutionUnauthorized{}
}

/*
ListTasksOfWebhookExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListTasksOfWebhookExecutionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list tasks of webhook execution unauthorized response has a 2xx status code
func (o *ListTasksOfWebhookExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks of webhook execution unauthorized response has a 3xx status code
func (o *ListTasksOfWebhookExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks of webhook execution unauthorized response has a 4xx status code
func (o *ListTasksOfWebhookExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tasks of webhook execution unauthorized response has a 5xx status code
func (o *ListTasksOfWebhookExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks of webhook execution unauthorized response a status code equal to that given
func (o *ListTasksOfWebhookExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListTasksOfWebhookExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListTasksOfWebhookExecutionUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListTasksOfWebhookExecutionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListTasksOfWebhookExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTasksOfWebhookExecutionForbidden creates a ListTasksOfWebhookExecutionForbidden with default headers values
func NewListTasksOfWebhookExecutionForbidden() *ListTasksOfWebhookExecutionForbidden {
	return &ListTasksOfWebhookExecutionForbidden{}
}

/*
ListTasksOfWebhookExecutionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListTasksOfWebhookExecutionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list tasks of webhook execution forbidden response has a 2xx status code
func (o *ListTasksOfWebhookExecutionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks of webhook execution forbidden response has a 3xx status code
func (o *ListTasksOfWebhookExecutionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks of webhook execution forbidden response has a 4xx status code
func (o *ListTasksOfWebhookExecutionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tasks of webhook execution forbidden response has a 5xx status code
func (o *ListTasksOfWebhookExecutionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks of webhook execution forbidden response a status code equal to that given
func (o *ListTasksOfWebhookExecutionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListTasksOfWebhookExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionForbidden  %+v", 403, o.Payload)
}

func (o *ListTasksOfWebhookExecutionForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionForbidden  %+v", 403, o.Payload)
}

func (o *ListTasksOfWebhookExecutionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListTasksOfWebhookExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTasksOfWebhookExecutionNotFound creates a ListTasksOfWebhookExecutionNotFound with default headers values
func NewListTasksOfWebhookExecutionNotFound() *ListTasksOfWebhookExecutionNotFound {
	return &ListTasksOfWebhookExecutionNotFound{}
}

/*
ListTasksOfWebhookExecutionNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListTasksOfWebhookExecutionNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list tasks of webhook execution not found response has a 2xx status code
func (o *ListTasksOfWebhookExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks of webhook execution not found response has a 3xx status code
func (o *ListTasksOfWebhookExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks of webhook execution not found response has a 4xx status code
func (o *ListTasksOfWebhookExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tasks of webhook execution not found response has a 5xx status code
func (o *ListTasksOfWebhookExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks of webhook execution not found response a status code equal to that given
func (o *ListTasksOfWebhookExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListTasksOfWebhookExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionNotFound  %+v", 404, o.Payload)
}

func (o *ListTasksOfWebhookExecutionNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionNotFound  %+v", 404, o.Payload)
}

func (o *ListTasksOfWebhookExecutionNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListTasksOfWebhookExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTasksOfWebhookExecutionInternalServerError creates a ListTasksOfWebhookExecutionInternalServerError with default headers values
func NewListTasksOfWebhookExecutionInternalServerError() *ListTasksOfWebhookExecutionInternalServerError {
	return &ListTasksOfWebhookExecutionInternalServerError{}
}

/*
ListTasksOfWebhookExecutionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListTasksOfWebhookExecutionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list tasks of webhook execution internal server error response has a 2xx status code
func (o *ListTasksOfWebhookExecutionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks of webhook execution internal server error response has a 3xx status code
func (o *ListTasksOfWebhookExecutionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks of webhook execution internal server error response has a 4xx status code
func (o *ListTasksOfWebhookExecutionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tasks of webhook execution internal server error response has a 5xx status code
func (o *ListTasksOfWebhookExecutionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list tasks of webhook execution internal server error response a status code equal to that given
func (o *ListTasksOfWebhookExecutionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListTasksOfWebhookExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTasksOfWebhookExecutionInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks][%d] listTasksOfWebhookExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTasksOfWebhookExecutionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListTasksOfWebhookExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
