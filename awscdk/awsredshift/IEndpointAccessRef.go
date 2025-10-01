package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointAccess.
// Experimental.
type IEndpointAccessRef interface {
	constructs.IConstruct
	// A reference to a EndpointAccess resource.
	// Experimental.
	EndpointAccessRef() *EndpointAccessReference
}

// The jsii proxy for IEndpointAccessRef
type jsiiProxy_IEndpointAccessRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEndpointAccessRef) EndpointAccessRef() *EndpointAccessReference {
	var returns *EndpointAccessReference
	_jsii_.Get(
		j,
		"endpointAccessRef",
		&returns,
	)
	return returns
}

