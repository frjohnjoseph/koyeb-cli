// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCatalogIntegration(params *GetCatalogIntegrationParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogIntegrationOK, error)

	GetCatalogServiceMixin0(params *GetCatalogServiceMixin0Params, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogServiceMixin0OK, error)

	ListCatalogIntegrationsMixin0(params *ListCatalogIntegrationsMixin0Params, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogIntegrationsMixin0OK, error)

	ListCatalogServiceIntegrations(params *ListCatalogServiceIntegrationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogServiceIntegrationsOK, error)

	ListCatalogServices(params *ListCatalogServicesParams, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogServicesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCatalogIntegration get catalog integration API
*/
func (a *Client) GetCatalogIntegration(params *GetCatalogIntegrationParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogIntegrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogIntegration",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/integrations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCatalogIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCatalogServiceMixin0 get catalog service mixin0 API
*/
func (a *Client) GetCatalogServiceMixin0(params *GetCatalogServiceMixin0Params, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogServiceMixin0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogServiceMixin0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogServiceMixin0",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/services/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogServiceMixin0Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogServiceMixin0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCatalogServiceMixin0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogIntegrationsMixin0 list catalog integrations mixin0 API
*/
func (a *Client) ListCatalogIntegrationsMixin0(params *ListCatalogIntegrationsMixin0Params, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogIntegrationsMixin0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogIntegrationsMixin0Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogIntegrationsMixin0",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/integrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogIntegrationsMixin0Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogIntegrationsMixin0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogIntegrationsMixin0: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogServiceIntegrations list catalog service integrations API
*/
func (a *Client) ListCatalogServiceIntegrations(params *ListCatalogServiceIntegrationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogServiceIntegrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogServiceIntegrationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogServiceIntegrations",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/services/{id}/integrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogServiceIntegrationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogServiceIntegrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogServiceIntegrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogServices list catalog services API
*/
func (a *Client) ListCatalogServices(params *ListCatalogServicesParams, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogServices",
		Method:             "GET",
		PathPattern:        "/v1/public/catalog/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
