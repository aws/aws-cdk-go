package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointGroup.
// Experimental.
type IEndpointGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEndpointGroupRef
type jsiiProxy_IEndpointGroupRef struct {
	internal.Type__constructsIConstruct
}

