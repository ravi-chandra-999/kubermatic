// Code generated by go-swagger; DO NOT EDIT.

package alibaba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListAlibabaInstanceTypesNoCredentialsReader is a Reader for the ListAlibabaInstanceTypesNoCredentials structure.
type ListAlibabaInstanceTypesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAlibabaInstanceTypesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAlibabaInstanceTypesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAlibabaInstanceTypesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAlibabaInstanceTypesNoCredentialsOK creates a ListAlibabaInstanceTypesNoCredentialsOK with default headers values
func NewListAlibabaInstanceTypesNoCredentialsOK() *ListAlibabaInstanceTypesNoCredentialsOK {
	return &ListAlibabaInstanceTypesNoCredentialsOK{}
}

/*ListAlibabaInstanceTypesNoCredentialsOK handles this case with default header values.

AlibabaInstanceTypeList
*/
type ListAlibabaInstanceTypesNoCredentialsOK struct {
	Payload models.AlibabaInstanceTypeList
}

func (o *ListAlibabaInstanceTypesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/alibaba/instancetypes][%d] listAlibabaInstanceTypesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListAlibabaInstanceTypesNoCredentialsOK) GetPayload() models.AlibabaInstanceTypeList {
	return o.Payload
}

func (o *ListAlibabaInstanceTypesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlibabaInstanceTypesNoCredentialsDefault creates a ListAlibabaInstanceTypesNoCredentialsDefault with default headers values
func NewListAlibabaInstanceTypesNoCredentialsDefault(code int) *ListAlibabaInstanceTypesNoCredentialsDefault {
	return &ListAlibabaInstanceTypesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListAlibabaInstanceTypesNoCredentialsDefault handles this case with default header values.

errorResponse
*/
type ListAlibabaInstanceTypesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list alibaba instance types no credentials default response
func (o *ListAlibabaInstanceTypesNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListAlibabaInstanceTypesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/alibaba/instancetypes][%d] listAlibabaInstanceTypesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListAlibabaInstanceTypesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAlibabaInstanceTypesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
