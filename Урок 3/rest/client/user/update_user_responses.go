// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"rest/models"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/users/{id}] updateUser", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*
UpdateUserOK describes a response with status code 200, with default header values.

User successfully updated
*/
type UpdateUserOK struct {
	Payload *models.UpdateUserResponse
}

// IsSuccess returns true when this update user o k response has a 2xx status code
func (o *UpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user o k response has a 3xx status code
func (o *UpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user o k response has a 4xx status code
func (o *UpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user o k response has a 5xx status code
func (o *UpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user o k response a status code equal to that given
func (o *UpdateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user o k response
func (o *UpdateUserOK) Code() int {
	return 200
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) String() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) GetPayload() *models.UpdateUserResponse {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserNoContent creates a UpdateUserNoContent with default headers values
func NewUpdateUserNoContent() *UpdateUserNoContent {
	return &UpdateUserNoContent{}
}

/*
UpdateUserNoContent describes a response with status code 204, with default header values.

No content
*/
type UpdateUserNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this update user no content response has a 2xx status code
func (o *UpdateUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user no content response has a 3xx status code
func (o *UpdateUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user no content response has a 4xx status code
func (o *UpdateUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user no content response has a 5xx status code
func (o *UpdateUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update user no content response a status code equal to that given
func (o *UpdateUserNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update user no content response
func (o *UpdateUserNoContent) Code() int {
	return 204
}

func (o *UpdateUserNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserNoContent  %+v", 204, o.Payload)
}

func (o *UpdateUserNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserNoContent  %+v", 204, o.Payload)
}

func (o *UpdateUserNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/*
UpdateUserForbidden describes a response with status code 403, with default header values.

Forrbidden
*/
type UpdateUserForbidden struct {
	Payload *models.ErrorMessage
}

// IsSuccess returns true when this update user forbidden response has a 2xx status code
func (o *UpdateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user forbidden response has a 3xx status code
func (o *UpdateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user forbidden response has a 4xx status code
func (o *UpdateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user forbidden response has a 5xx status code
func (o *UpdateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update user forbidden response a status code equal to that given
func (o *UpdateUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update user forbidden response
func (o *UpdateUserForbidden) Code() int {
	return 403
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserInternalServerError creates a UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {
	return &UpdateUserInternalServerError{}
}

/*
UpdateUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateUserInternalServerError struct {
	Payload *models.ErrorMessage
}

// IsSuccess returns true when this update user internal server error response has a 2xx status code
func (o *UpdateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user internal server error response has a 3xx status code
func (o *UpdateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user internal server error response has a 4xx status code
func (o *UpdateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user internal server error response has a 5xx status code
func (o *UpdateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update user internal server error response a status code equal to that given
func (o *UpdateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update user internal server error response
func (o *UpdateUserInternalServerError) Code() int {
	return 500
}

func (o *UpdateUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] updateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
