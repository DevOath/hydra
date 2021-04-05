// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// CreateGrantTypeJwtBearerReader is a Reader for the CreateGrantTypeJwtBearer structure.
type CreateGrantTypeJwtBearerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGrantTypeJwtBearerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGrantTypeJwtBearerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGrantTypeJwtBearerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateGrantTypeJwtBearerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGrantTypeJwtBearerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGrantTypeJwtBearerCreated creates a CreateGrantTypeJwtBearerCreated with default headers values
func NewCreateGrantTypeJwtBearerCreated() *CreateGrantTypeJwtBearerCreated {
	return &CreateGrantTypeJwtBearerCreated{}
}

/*CreateGrantTypeJwtBearerCreated handles this case with default header values.

grantTypeJwtBearer
*/
type CreateGrantTypeJwtBearerCreated struct {
	Payload *models.GrantTypeJwtBearer
}

func (o *CreateGrantTypeJwtBearerCreated) Error() string {
	return fmt.Sprintf("[POST /grants/jwt-bearer][%d] createGrantTypeJwtBearerCreated  %+v", 201, o.Payload)
}

func (o *CreateGrantTypeJwtBearerCreated) GetPayload() *models.GrantTypeJwtBearer {
	return o.Payload
}

func (o *CreateGrantTypeJwtBearerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrantTypeJwtBearer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGrantTypeJwtBearerBadRequest creates a CreateGrantTypeJwtBearerBadRequest with default headers values
func NewCreateGrantTypeJwtBearerBadRequest() *CreateGrantTypeJwtBearerBadRequest {
	return &CreateGrantTypeJwtBearerBadRequest{}
}

/*CreateGrantTypeJwtBearerBadRequest handles this case with default header values.

genericError
*/
type CreateGrantTypeJwtBearerBadRequest struct {
	Payload *models.GenericError
}

func (o *CreateGrantTypeJwtBearerBadRequest) Error() string {
	return fmt.Sprintf("[POST /grants/jwt-bearer][%d] createGrantTypeJwtBearerBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGrantTypeJwtBearerBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateGrantTypeJwtBearerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGrantTypeJwtBearerConflict creates a CreateGrantTypeJwtBearerConflict with default headers values
func NewCreateGrantTypeJwtBearerConflict() *CreateGrantTypeJwtBearerConflict {
	return &CreateGrantTypeJwtBearerConflict{}
}

/*CreateGrantTypeJwtBearerConflict handles this case with default header values.

genericError
*/
type CreateGrantTypeJwtBearerConflict struct {
	Payload *models.GenericError
}

func (o *CreateGrantTypeJwtBearerConflict) Error() string {
	return fmt.Sprintf("[POST /grants/jwt-bearer][%d] createGrantTypeJwtBearerConflict  %+v", 409, o.Payload)
}

func (o *CreateGrantTypeJwtBearerConflict) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateGrantTypeJwtBearerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGrantTypeJwtBearerInternalServerError creates a CreateGrantTypeJwtBearerInternalServerError with default headers values
func NewCreateGrantTypeJwtBearerInternalServerError() *CreateGrantTypeJwtBearerInternalServerError {
	return &CreateGrantTypeJwtBearerInternalServerError{}
}

/*CreateGrantTypeJwtBearerInternalServerError handles this case with default header values.

genericError
*/
type CreateGrantTypeJwtBearerInternalServerError struct {
	Payload *models.GenericError
}

func (o *CreateGrantTypeJwtBearerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /grants/jwt-bearer][%d] createGrantTypeJwtBearerInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateGrantTypeJwtBearerInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateGrantTypeJwtBearerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}