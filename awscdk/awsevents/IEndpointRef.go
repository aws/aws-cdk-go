package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Endpoint.
// Experimental.
type IEndpointRef interface {
	constructs.IConstruct
	// A reference to a Endpoint resource.
	// Experimental.
	EndpointRef() *EndpointReference
}

// The jsii proxy for IEndpointRef
type jsiiProxy_IEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEndpointRef) EndpointRef() *EndpointReference {
	var returns *EndpointReference
	_jsii_.Get(
		j,
		"endpointRef",
		&returns,
	)
	return returns
}

