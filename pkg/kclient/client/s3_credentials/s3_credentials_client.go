// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new s3 credentials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s3 credentials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteS3Credential(params *DeleteS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteS3CredentialOK, error)

	GetS3Credential(params *GetS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*GetS3CredentialOK, error)

	ListS3Credentials(params *ListS3CredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListS3CredentialsOK, error)

	NewS3Credential(params *NewS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*NewS3CredentialOK, error)

	UpdateS3Credential(params *UpdateS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateS3CredentialOK, error)

	UpdateS3Credential2(params *UpdateS3Credential2Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateS3Credential2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteS3Credential delete s3 credential API
*/
func (a *Client) DeleteS3Credential(params *DeleteS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteS3CredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteS3CredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteS3Credential",
		Method:             "DELETE",
		PathPattern:        "/v1/s3_credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteS3CredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteS3CredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteS3Credential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetS3Credential get s3 credential API
*/
func (a *Client) GetS3Credential(params *GetS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*GetS3CredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetS3CredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetS3Credential",
		Method:             "GET",
		PathPattern:        "/v1/s3_credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetS3CredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetS3CredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetS3Credential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListS3Credentials list s3 credentials API
*/
func (a *Client) ListS3Credentials(params *ListS3CredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListS3CredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListS3CredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListS3Credentials",
		Method:             "GET",
		PathPattern:        "/v1/s3_credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListS3CredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListS3CredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListS3Credentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NewS3Credential new s3 credential API
*/
func (a *Client) NewS3Credential(params *NewS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*NewS3CredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNewS3CredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NewS3Credential",
		Method:             "POST",
		PathPattern:        "/v1/s3_credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NewS3CredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NewS3CredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NewS3Credential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateS3Credential update s3 credential API
*/
func (a *Client) UpdateS3Credential(params *UpdateS3CredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateS3CredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateS3CredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateS3Credential",
		Method:             "PUT",
		PathPattern:        "/v1/s3_credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateS3CredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateS3CredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateS3Credential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateS3Credential2 update s3 credential2 API
*/
func (a *Client) UpdateS3Credential2(params *UpdateS3Credential2Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateS3Credential2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateS3Credential2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateS3Credential2",
		Method:             "PATCH",
		PathPattern:        "/v1/s3_credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateS3Credential2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateS3Credential2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateS3Credential2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
