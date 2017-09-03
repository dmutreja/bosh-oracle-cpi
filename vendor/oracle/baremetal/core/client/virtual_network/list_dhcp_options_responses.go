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

// ListDhcpOptionsReader is a Reader for the ListDhcpOptions structure.
type ListDhcpOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDhcpOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDhcpOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListDhcpOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListDhcpOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListDhcpOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListDhcpOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDhcpOptionsOK creates a ListDhcpOptionsOK with default headers values
func NewListDhcpOptionsOK() *ListDhcpOptionsOK {
	return &ListDhcpOptionsOK{}
}

/*ListDhcpOptionsOK handles this case with default header values.

The list is being retrieved.
*/
type ListDhcpOptionsOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.DhcpOptions
}

func (o *ListDhcpOptionsOK) Error() string {
	return fmt.Sprintf("[GET /dhcps][%d] listDhcpOptionsOK  %+v", 200, o.Payload)
}

func (o *ListDhcpOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListDhcpOptionsUnauthorized creates a ListDhcpOptionsUnauthorized with default headers values
func NewListDhcpOptionsUnauthorized() *ListDhcpOptionsUnauthorized {
	return &ListDhcpOptionsUnauthorized{}
}

/*ListDhcpOptionsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListDhcpOptionsUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListDhcpOptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dhcps][%d] listDhcpOptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListDhcpOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDhcpOptionsNotFound creates a ListDhcpOptionsNotFound with default headers values
func NewListDhcpOptionsNotFound() *ListDhcpOptionsNotFound {
	return &ListDhcpOptionsNotFound{}
}

/*ListDhcpOptionsNotFound handles this case with default header values.

Not Found
*/
type ListDhcpOptionsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListDhcpOptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /dhcps][%d] listDhcpOptionsNotFound  %+v", 404, o.Payload)
}

func (o *ListDhcpOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDhcpOptionsInternalServerError creates a ListDhcpOptionsInternalServerError with default headers values
func NewListDhcpOptionsInternalServerError() *ListDhcpOptionsInternalServerError {
	return &ListDhcpOptionsInternalServerError{}
}

/*ListDhcpOptionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListDhcpOptionsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListDhcpOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dhcps][%d] listDhcpOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDhcpOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDhcpOptionsDefault creates a ListDhcpOptionsDefault with default headers values
func NewListDhcpOptionsDefault(code int) *ListDhcpOptionsDefault {
	return &ListDhcpOptionsDefault{
		_statusCode: code,
	}
}

/*ListDhcpOptionsDefault handles this case with default header values.

An error has occurred.
*/
type ListDhcpOptionsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list dhcp options default response
func (o *ListDhcpOptionsDefault) Code() int {
	return o._statusCode
}

func (o *ListDhcpOptionsDefault) Error() string {
	return fmt.Sprintf("[GET /dhcps][%d] ListDhcpOptions default  %+v", o._statusCode, o.Payload)
}

func (o *ListDhcpOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}