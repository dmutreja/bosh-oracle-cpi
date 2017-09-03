package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// CreateCrossConnectReader is a Reader for the CreateCrossConnect structure.
type CreateCrossConnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCrossConnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCrossConnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateCrossConnectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateCrossConnectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateCrossConnectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateCrossConnectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateCrossConnectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCrossConnectOK creates a CreateCrossConnectOK with default headers values
func NewCreateCrossConnectOK() *CreateCrossConnectOK {
	return &CreateCrossConnectOK{}
}

/*CreateCrossConnectOK handles this case with default header values.

The cross-connect was created.
*/
type CreateCrossConnectOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.CrossConnect
}

func (o *CreateCrossConnectOK) Error() string {
	return fmt.Sprintf("[POST /crossConnects][%d] createCrossConnectOK  %+v", 200, o.Payload)
}

func (o *CreateCrossConnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.CrossConnect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectBadRequest creates a CreateCrossConnectBadRequest with default headers values
func NewCreateCrossConnectBadRequest() *CreateCrossConnectBadRequest {
	return &CreateCrossConnectBadRequest{}
}

/*CreateCrossConnectBadRequest handles this case with default header values.

Bad Request
*/
type CreateCrossConnectBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectBadRequest) Error() string {
	return fmt.Sprintf("[POST /crossConnects][%d] createCrossConnectBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCrossConnectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectUnauthorized creates a CreateCrossConnectUnauthorized with default headers values
func NewCreateCrossConnectUnauthorized() *CreateCrossConnectUnauthorized {
	return &CreateCrossConnectUnauthorized{}
}

/*CreateCrossConnectUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCrossConnectUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /crossConnects][%d] createCrossConnectUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCrossConnectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectConflict creates a CreateCrossConnectConflict with default headers values
func NewCreateCrossConnectConflict() *CreateCrossConnectConflict {
	return &CreateCrossConnectConflict{}
}

/*CreateCrossConnectConflict handles this case with default header values.

Conflict
*/
type CreateCrossConnectConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectConflict) Error() string {
	return fmt.Sprintf("[POST /crossConnects][%d] createCrossConnectConflict  %+v", 409, o.Payload)
}

func (o *CreateCrossConnectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectInternalServerError creates a CreateCrossConnectInternalServerError with default headers values
func NewCreateCrossConnectInternalServerError() *CreateCrossConnectInternalServerError {
	return &CreateCrossConnectInternalServerError{}
}

/*CreateCrossConnectInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateCrossConnectInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateCrossConnectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /crossConnects][%d] createCrossConnectInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCrossConnectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCrossConnectDefault creates a CreateCrossConnectDefault with default headers values
func NewCreateCrossConnectDefault(code int) *CreateCrossConnectDefault {
	return &CreateCrossConnectDefault{
		_statusCode: code,
	}
}

/*CreateCrossConnectDefault handles this case with default header values.

An error has occurred.
*/
type CreateCrossConnectDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create cross connect default response
func (o *CreateCrossConnectDefault) Code() int {
	return o._statusCode
}

func (o *CreateCrossConnectDefault) Error() string {
	return fmt.Sprintf("[POST /crossConnects][%d] CreateCrossConnect default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCrossConnectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}