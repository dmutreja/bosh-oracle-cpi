package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/identity/models"
)

// AddUserToGroupReader is a Reader for the AddUserToGroup structure.
type AddUserToGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserToGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddUserToGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddUserToGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddUserToGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddUserToGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddUserToGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddUserToGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddUserToGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddUserToGroupOK creates a AddUserToGroupOK with default headers values
func NewAddUserToGroupOK() *AddUserToGroupOK {
	return &AddUserToGroupOK{}
}

/*AddUserToGroupOK handles this case with default header values.

The user is being added to the group.
*/
type AddUserToGroupOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.UserGroupMembership
}

func (o *AddUserToGroupOK) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] addUserToGroupOK  %+v", 200, o.Payload)
}

func (o *AddUserToGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.UserGroupMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupBadRequest creates a AddUserToGroupBadRequest with default headers values
func NewAddUserToGroupBadRequest() *AddUserToGroupBadRequest {
	return &AddUserToGroupBadRequest{}
}

/*AddUserToGroupBadRequest handles this case with default header values.

Bad Request
*/
type AddUserToGroupBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AddUserToGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] addUserToGroupBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserToGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupForbidden creates a AddUserToGroupForbidden with default headers values
func NewAddUserToGroupForbidden() *AddUserToGroupForbidden {
	return &AddUserToGroupForbidden{}
}

/*AddUserToGroupForbidden handles this case with default header values.

Forbidden
*/
type AddUserToGroupForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AddUserToGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] addUserToGroupForbidden  %+v", 403, o.Payload)
}

func (o *AddUserToGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupNotFound creates a AddUserToGroupNotFound with default headers values
func NewAddUserToGroupNotFound() *AddUserToGroupNotFound {
	return &AddUserToGroupNotFound{}
}

/*AddUserToGroupNotFound handles this case with default header values.

Not Found
*/
type AddUserToGroupNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AddUserToGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] addUserToGroupNotFound  %+v", 404, o.Payload)
}

func (o *AddUserToGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupConflict creates a AddUserToGroupConflict with default headers values
func NewAddUserToGroupConflict() *AddUserToGroupConflict {
	return &AddUserToGroupConflict{}
}

/*AddUserToGroupConflict handles this case with default header values.

Conflict
*/
type AddUserToGroupConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AddUserToGroupConflict) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] addUserToGroupConflict  %+v", 409, o.Payload)
}

func (o *AddUserToGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupInternalServerError creates a AddUserToGroupInternalServerError with default headers values
func NewAddUserToGroupInternalServerError() *AddUserToGroupInternalServerError {
	return &AddUserToGroupInternalServerError{}
}

/*AddUserToGroupInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddUserToGroupInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AddUserToGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] addUserToGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserToGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToGroupDefault creates a AddUserToGroupDefault with default headers values
func NewAddUserToGroupDefault(code int) *AddUserToGroupDefault {
	return &AddUserToGroupDefault{
		_statusCode: code,
	}
}

/*AddUserToGroupDefault handles this case with default header values.

An error has occurred.

*/
type AddUserToGroupDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the add user to group default response
func (o *AddUserToGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddUserToGroupDefault) Error() string {
	return fmt.Sprintf("[POST /userGroupMemberships/][%d] AddUserToGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddUserToGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}