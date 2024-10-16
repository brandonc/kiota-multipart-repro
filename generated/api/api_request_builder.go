package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApiRequestBuilder builds and executes requests for operations under \api
type ApiRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewApiRequestBuilderInternal instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    m := &ApiRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api", pathParameters),
    }
    return m
}
// NewApiRequestBuilder instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint the endpoint property
// returns a *EndpointRequestBuilder when successful
func (m *ApiRequestBuilder) Endpoint()(*EndpointRequestBuilder) {
    return NewEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
