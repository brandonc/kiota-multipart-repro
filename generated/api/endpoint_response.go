package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EndpointPostResponseable instead.
type EndpointResponse struct {
    EndpointPostResponse
}
// NewEndpointResponse instantiates a new EndpointResponse and sets the default values.
func NewEndpointResponse()(*EndpointResponse) {
    m := &EndpointResponse{
        EndpointPostResponse: *NewEndpointPostResponse(),
    }
    return m
}
// CreateEndpointResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEndpointResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndpointResponse(), nil
}
// Deprecated: This class is obsolete. Use EndpointPostResponseable instead.
type EndpointResponseable interface {
    EndpointPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
