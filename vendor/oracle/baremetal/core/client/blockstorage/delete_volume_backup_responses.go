package blockstorage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// DeleteVolumeBackupReader is a Reader for the DeleteVolumeBackup structure.
type DeleteVolumeBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVolumeBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVolumeBackupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteVolumeBackupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteVolumeBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteVolumeBackupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteVolumeBackupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteVolumeBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteVolumeBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVolumeBackupNoContent creates a DeleteVolumeBackupNoContent with default headers values
func NewDeleteVolumeBackupNoContent() *DeleteVolumeBackupNoContent {
	return &DeleteVolumeBackupNoContent{}
}

/*DeleteVolumeBackupNoContent handles this case with default header values.

The volume backup is being deleted.
*/
type DeleteVolumeBackupNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteVolumeBackupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] deleteVolumeBackupNoContent ", 204)
}

func (o *DeleteVolumeBackupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteVolumeBackupBadRequest creates a DeleteVolumeBackupBadRequest with default headers values
func NewDeleteVolumeBackupBadRequest() *DeleteVolumeBackupBadRequest {
	return &DeleteVolumeBackupBadRequest{}
}

/*DeleteVolumeBackupBadRequest handles this case with default header values.

Bad Request
*/
type DeleteVolumeBackupBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeBackupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] deleteVolumeBackupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVolumeBackupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeBackupUnauthorized creates a DeleteVolumeBackupUnauthorized with default headers values
func NewDeleteVolumeBackupUnauthorized() *DeleteVolumeBackupUnauthorized {
	return &DeleteVolumeBackupUnauthorized{}
}

/*DeleteVolumeBackupUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteVolumeBackupUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeBackupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] deleteVolumeBackupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteVolumeBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeBackupNotFound creates a DeleteVolumeBackupNotFound with default headers values
func NewDeleteVolumeBackupNotFound() *DeleteVolumeBackupNotFound {
	return &DeleteVolumeBackupNotFound{}
}

/*DeleteVolumeBackupNotFound handles this case with default header values.

Not Found
*/
type DeleteVolumeBackupNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeBackupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] deleteVolumeBackupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVolumeBackupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeBackupConflict creates a DeleteVolumeBackupConflict with default headers values
func NewDeleteVolumeBackupConflict() *DeleteVolumeBackupConflict {
	return &DeleteVolumeBackupConflict{}
}

/*DeleteVolumeBackupConflict handles this case with default header values.

Conflict
*/
type DeleteVolumeBackupConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeBackupConflict) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] deleteVolumeBackupConflict  %+v", 409, o.Payload)
}

func (o *DeleteVolumeBackupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeBackupInternalServerError creates a DeleteVolumeBackupInternalServerError with default headers values
func NewDeleteVolumeBackupInternalServerError() *DeleteVolumeBackupInternalServerError {
	return &DeleteVolumeBackupInternalServerError{}
}

/*DeleteVolumeBackupInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteVolumeBackupInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeBackupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] deleteVolumeBackupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVolumeBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeBackupDefault creates a DeleteVolumeBackupDefault with default headers values
func NewDeleteVolumeBackupDefault(code int) *DeleteVolumeBackupDefault {
	return &DeleteVolumeBackupDefault{
		_statusCode: code,
	}
}

/*DeleteVolumeBackupDefault handles this case with default header values.

An error has occurred.
*/
type DeleteVolumeBackupDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete volume backup default response
func (o *DeleteVolumeBackupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVolumeBackupDefault) Error() string {
	return fmt.Sprintf("[DELETE /volumeBackups/{volumeBackupId}][%d] DeleteVolumeBackup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVolumeBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}