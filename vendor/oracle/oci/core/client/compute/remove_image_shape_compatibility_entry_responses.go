package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// RemoveImageShapeCompatibilityEntryReader is a Reader for the RemoveImageShapeCompatibilityEntry structure.
type RemoveImageShapeCompatibilityEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveImageShapeCompatibilityEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveImageShapeCompatibilityEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveImageShapeCompatibilityEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveImageShapeCompatibilityEntryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveImageShapeCompatibilityEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveImageShapeCompatibilityEntryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRemoveImageShapeCompatibilityEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveImageShapeCompatibilityEntryNoContent creates a RemoveImageShapeCompatibilityEntryNoContent with default headers values
func NewRemoveImageShapeCompatibilityEntryNoContent() *RemoveImageShapeCompatibilityEntryNoContent {
	return &RemoveImageShapeCompatibilityEntryNoContent{}
}

/*RemoveImageShapeCompatibilityEntryNoContent handles this case with default header values.

The shape has been removed from the compatible shapes list.
*/
type RemoveImageShapeCompatibilityEntryNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *RemoveImageShapeCompatibilityEntryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageId}/shapes/{shapeName}][%d] removeImageShapeCompatibilityEntryNoContent ", 204)
}

func (o *RemoveImageShapeCompatibilityEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewRemoveImageShapeCompatibilityEntryBadRequest creates a RemoveImageShapeCompatibilityEntryBadRequest with default headers values
func NewRemoveImageShapeCompatibilityEntryBadRequest() *RemoveImageShapeCompatibilityEntryBadRequest {
	return &RemoveImageShapeCompatibilityEntryBadRequest{}
}

/*RemoveImageShapeCompatibilityEntryBadRequest handles this case with default header values.

Bad Request
*/
type RemoveImageShapeCompatibilityEntryBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveImageShapeCompatibilityEntryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageId}/shapes/{shapeName}][%d] removeImageShapeCompatibilityEntryBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveImageShapeCompatibilityEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveImageShapeCompatibilityEntryUnauthorized creates a RemoveImageShapeCompatibilityEntryUnauthorized with default headers values
func NewRemoveImageShapeCompatibilityEntryUnauthorized() *RemoveImageShapeCompatibilityEntryUnauthorized {
	return &RemoveImageShapeCompatibilityEntryUnauthorized{}
}

/*RemoveImageShapeCompatibilityEntryUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveImageShapeCompatibilityEntryUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveImageShapeCompatibilityEntryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageId}/shapes/{shapeName}][%d] removeImageShapeCompatibilityEntryUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveImageShapeCompatibilityEntryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveImageShapeCompatibilityEntryNotFound creates a RemoveImageShapeCompatibilityEntryNotFound with default headers values
func NewRemoveImageShapeCompatibilityEntryNotFound() *RemoveImageShapeCompatibilityEntryNotFound {
	return &RemoveImageShapeCompatibilityEntryNotFound{}
}

/*RemoveImageShapeCompatibilityEntryNotFound handles this case with default header values.

Not Found
*/
type RemoveImageShapeCompatibilityEntryNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveImageShapeCompatibilityEntryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageId}/shapes/{shapeName}][%d] removeImageShapeCompatibilityEntryNotFound  %+v", 404, o.Payload)
}

func (o *RemoveImageShapeCompatibilityEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveImageShapeCompatibilityEntryInternalServerError creates a RemoveImageShapeCompatibilityEntryInternalServerError with default headers values
func NewRemoveImageShapeCompatibilityEntryInternalServerError() *RemoveImageShapeCompatibilityEntryInternalServerError {
	return &RemoveImageShapeCompatibilityEntryInternalServerError{}
}

/*RemoveImageShapeCompatibilityEntryInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoveImageShapeCompatibilityEntryInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *RemoveImageShapeCompatibilityEntryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageId}/shapes/{shapeName}][%d] removeImageShapeCompatibilityEntryInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveImageShapeCompatibilityEntryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveImageShapeCompatibilityEntryDefault creates a RemoveImageShapeCompatibilityEntryDefault with default headers values
func NewRemoveImageShapeCompatibilityEntryDefault(code int) *RemoveImageShapeCompatibilityEntryDefault {
	return &RemoveImageShapeCompatibilityEntryDefault{
		_statusCode: code,
	}
}

/*RemoveImageShapeCompatibilityEntryDefault handles this case with default header values.

An error has occurred.
*/
type RemoveImageShapeCompatibilityEntryDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the remove image shape compatibility entry default response
func (o *RemoveImageShapeCompatibilityEntryDefault) Code() int {
	return o._statusCode
}

func (o *RemoveImageShapeCompatibilityEntryDefault) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageId}/shapes/{shapeName}][%d] RemoveImageShapeCompatibilityEntry default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveImageShapeCompatibilityEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
