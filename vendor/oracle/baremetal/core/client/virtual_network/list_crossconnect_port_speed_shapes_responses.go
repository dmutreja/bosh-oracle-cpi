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

// ListCrossconnectPortSpeedShapesReader is a Reader for the ListCrossconnectPortSpeedShapes structure.
type ListCrossconnectPortSpeedShapesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCrossconnectPortSpeedShapesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCrossconnectPortSpeedShapesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListCrossconnectPortSpeedShapesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListCrossconnectPortSpeedShapesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListCrossconnectPortSpeedShapesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListCrossconnectPortSpeedShapesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCrossconnectPortSpeedShapesOK creates a ListCrossconnectPortSpeedShapesOK with default headers values
func NewListCrossconnectPortSpeedShapesOK() *ListCrossconnectPortSpeedShapesOK {
	return &ListCrossconnectPortSpeedShapesOK{}
}

/*ListCrossconnectPortSpeedShapesOK handles this case with default header values.

The list is being retrieved.
*/
type ListCrossconnectPortSpeedShapesOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.CrossConnectPortSpeedShape
}

func (o *ListCrossconnectPortSpeedShapesOK) Error() string {
	return fmt.Sprintf("[GET /crossConnectPortSpeedShapes][%d] listCrossconnectPortSpeedShapesOK  %+v", 200, o.Payload)
}

func (o *ListCrossconnectPortSpeedShapesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-next-page
	o.OpcNextPage = response.GetHeader("opc-next-page")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCrossconnectPortSpeedShapesUnauthorized creates a ListCrossconnectPortSpeedShapesUnauthorized with default headers values
func NewListCrossconnectPortSpeedShapesUnauthorized() *ListCrossconnectPortSpeedShapesUnauthorized {
	return &ListCrossconnectPortSpeedShapesUnauthorized{}
}

/*ListCrossconnectPortSpeedShapesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCrossconnectPortSpeedShapesUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListCrossconnectPortSpeedShapesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /crossConnectPortSpeedShapes][%d] listCrossconnectPortSpeedShapesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCrossconnectPortSpeedShapesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCrossconnectPortSpeedShapesNotFound creates a ListCrossconnectPortSpeedShapesNotFound with default headers values
func NewListCrossconnectPortSpeedShapesNotFound() *ListCrossconnectPortSpeedShapesNotFound {
	return &ListCrossconnectPortSpeedShapesNotFound{}
}

/*ListCrossconnectPortSpeedShapesNotFound handles this case with default header values.

Not Found
*/
type ListCrossconnectPortSpeedShapesNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListCrossconnectPortSpeedShapesNotFound) Error() string {
	return fmt.Sprintf("[GET /crossConnectPortSpeedShapes][%d] listCrossconnectPortSpeedShapesNotFound  %+v", 404, o.Payload)
}

func (o *ListCrossconnectPortSpeedShapesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCrossconnectPortSpeedShapesInternalServerError creates a ListCrossconnectPortSpeedShapesInternalServerError with default headers values
func NewListCrossconnectPortSpeedShapesInternalServerError() *ListCrossconnectPortSpeedShapesInternalServerError {
	return &ListCrossconnectPortSpeedShapesInternalServerError{}
}

/*ListCrossconnectPortSpeedShapesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCrossconnectPortSpeedShapesInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListCrossconnectPortSpeedShapesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /crossConnectPortSpeedShapes][%d] listCrossconnectPortSpeedShapesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCrossconnectPortSpeedShapesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCrossconnectPortSpeedShapesDefault creates a ListCrossconnectPortSpeedShapesDefault with default headers values
func NewListCrossconnectPortSpeedShapesDefault(code int) *ListCrossconnectPortSpeedShapesDefault {
	return &ListCrossconnectPortSpeedShapesDefault{
		_statusCode: code,
	}
}

/*ListCrossconnectPortSpeedShapesDefault handles this case with default header values.

An error has occurred.
*/
type ListCrossconnectPortSpeedShapesDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list crossconnect port speed shapes default response
func (o *ListCrossconnectPortSpeedShapesDefault) Code() int {
	return o._statusCode
}

func (o *ListCrossconnectPortSpeedShapesDefault) Error() string {
	return fmt.Sprintf("[GET /crossConnectPortSpeedShapes][%d] ListCrossconnectPortSpeedShapes default  %+v", o._statusCode, o.Payload)
}

func (o *ListCrossconnectPortSpeedShapesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}