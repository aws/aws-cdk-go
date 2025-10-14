package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointAuthorization.
// Experimental.
type IEndpointAuthorizationRef interface {
	constructs.IConstruct
	// A reference to a EndpointAuthorization resource.
	// Experimental.
	EndpointAuthorizationRef() *EndpointAuthorizationReference
}

// The jsii proxy for IEndpointAuthorizationRef
type jsiiProxy_IEndpointAuthorizationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEndpointAuthorizationRef) EndpointAuthorizationRef() *EndpointAuthorizationReference {
	var returns *EndpointAuthorizationReference
	_jsii_.Get(
		j,
		"endpointAuthorizationRef",
		&returns,
	)
	return returns
}

