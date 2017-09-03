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

// ListRouteTablesReader is a Reader for the ListRouteTables structure.
type ListRouteTablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRouteTablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRouteTablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListRouteTablesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListRouteTablesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListRouteTablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListRouteTablesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRouteTablesOK creates a ListRouteTablesOK with default headers values
func NewListRouteTablesOK() *ListRouteTablesOK {
	return &ListRouteTablesOK{}
}

/*ListRouteTablesOK handles this case with default header values.

The list is being retrieved.
*/
type ListRouteTablesOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.RouteTable
}

func (o *ListRouteTablesOK) Error() string {
	return fmt.Sprintf("[GET /routeTables][%d] listRouteTablesOK  %+v", 200, o.Payload)
}

func (o *ListRouteTablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRouteTablesUnauthorized creates a ListRouteTablesUnauthorized with default headers values
func NewListRouteTablesUnauthorized() *ListRouteTablesUnauthorized {
	return &ListRouteTablesUnauthorized{}
}

/*ListRouteTablesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRouteTablesUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListRouteTablesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /routeTables][%d] listRouteTablesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRouteTablesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRouteTablesNotFound creates a ListRouteTablesNotFound with default headers values
func NewListRouteTablesNotFound() *ListRouteTablesNotFound {
	return &ListRouteTablesNotFound{}
}

/*ListRouteTablesNotFound handles this case with default header values.

Not Found
*/
type ListRouteTablesNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListRouteTablesNotFound) Error() string {
	return fmt.Sprintf("[GET /routeTables][%d] listRouteTablesNotFound  %+v", 404, o.Payload)
}

func (o *ListRouteTablesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRouteTablesInternalServerError creates a ListRouteTablesInternalServerError with default headers values
func NewListRouteTablesInternalServerError() *ListRouteTablesInternalServerError {
	return &ListRouteTablesInternalServerError{}
}

/*ListRouteTablesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListRouteTablesInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListRouteTablesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /routeTables][%d] listRouteTablesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRouteTablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRouteTablesDefault creates a ListRouteTablesDefault with default headers values
func NewListRouteTablesDefault(code int) *ListRouteTablesDefault {
	return &ListRouteTablesDefault{
		_statusCode: code,
	}
}

/*ListRouteTablesDefault handles this case with default header values.

An error has occurred.
*/
type ListRouteTablesDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list route tables default response
func (o *ListRouteTablesDefault) Code() int {
	return o._statusCode
}

func (o *ListRouteTablesDefault) Error() string {
	return fmt.Sprintf("[GET /routeTables][%d] ListRouteTables default  %+v", o._statusCode, o.Payload)
}

func (o *ListRouteTablesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}