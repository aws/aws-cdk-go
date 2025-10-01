package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointGroup.
// Experimental.
type IEndpointGroupRef interface {
	constructs.IConstruct
	// A reference to a EndpointGroup resource.
	// Experimental.
	EndpointGroupRef() *EndpointGroupReference
}

// The jsii proxy for IEndpointGroupRef
type jsiiProxy_IEndpointGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEndpointGroupRef) EndpointGroupRef() *EndpointGroupReference {
	var returns *EndpointGroupReference
	_jsii_.Get(
		j,
		"endpointGroupRef",
		&returns,
	)
	return returns
}

