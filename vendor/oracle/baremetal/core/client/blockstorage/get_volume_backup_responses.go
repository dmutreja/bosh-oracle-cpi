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

// GetVolumeBackupReader is a Reader for the GetVolumeBackup structure.
type GetVolumeBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumeBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVolumeBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetVolumeBackupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetVolumeBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetVolumeBackupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetVolumeBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVolumeBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVolumeBackupOK creates a GetVolumeBackupOK with default headers values
func NewGetVolumeBackupOK() *GetVolumeBackupOK {
	return &GetVolumeBackupOK{}
}

/*GetVolumeBackupOK handles this case with default header values.

The volume backup was retrieved.
*/
type GetVolumeBackupOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.VolumeBackup
}

func (o *GetVolumeBackupOK) Error() string {
	return fmt.Sprintf("[GET /volumeBackups/{volumeBackupId}][%d] getVolumeBackupOK  %+v", 200, o.Payload)
}

func (o *GetVolumeBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.VolumeBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupBadRequest creates a GetVolumeBackupBadRequest with default headers values
func NewGetVolumeBackupBadRequest() *GetVolumeBackupBadRequest {
	return &GetVolumeBackupBadRequest{}
}

/*GetVolumeBackupBadRequest handles this case with default header values.

Bad Request
*/
type GetVolumeBackupBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupBadRequest) Error() string {
	return fmt.Sprintf("[GET /volumeBackups/{volumeBackupId}][%d] getVolumeBackupBadRequest  %+v", 400, o.Payload)
}

func (o *GetVolumeBackupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupUnauthorized creates a GetVolumeBackupUnauthorized with default headers values
func NewGetVolumeBackupUnauthorized() *GetVolumeBackupUnauthorized {
	return &GetVolumeBackupUnauthorized{}
}

/*GetVolumeBackupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetVolumeBackupUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /volumeBackups/{volumeBackupId}][%d] getVolumeBackupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVolumeBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupNotFound creates a GetVolumeBackupNotFound with default headers values
func NewGetVolumeBackupNotFound() *GetVolumeBackupNotFound {
	return &GetVolumeBackupNotFound{}
}

/*GetVolumeBackupNotFound handles this case with default header values.

Not Found
*/
type GetVolumeBackupNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupNotFound) Error() string {
	return fmt.Sprintf("[GET /volumeBackups/{volumeBackupId}][%d] getVolumeBackupNotFound  %+v", 404, o.Payload)
}

func (o *GetVolumeBackupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupInternalServerError creates a GetVolumeBackupInternalServerError with default headers values
func NewGetVolumeBackupInternalServerError() *GetVolumeBackupInternalServerError {
	return &GetVolumeBackupInternalServerError{}
}

/*GetVolumeBackupInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetVolumeBackupInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumeBackups/{volumeBackupId}][%d] getVolumeBackupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVolumeBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupDefault creates a GetVolumeBackupDefault with default headers values
func NewGetVolumeBackupDefault(code int) *GetVolumeBackupDefault {
	return &GetVolumeBackupDefault{
		_statusCode: code,
	}
}

/*GetVolumeBackupDefault handles this case with default header values.

An error has occurred.
*/
type GetVolumeBackupDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get volume backup default response
func (o *GetVolumeBackupDefault) Code() int {
	return o._statusCode
}

func (o *GetVolumeBackupDefault) Error() string {
	return fmt.Sprintf("[GET /volumeBackups/{volumeBackupId}][%d] GetVolumeBackup default  %+v", o._statusCode, o.Payload)
}

func (o *GetVolumeBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}