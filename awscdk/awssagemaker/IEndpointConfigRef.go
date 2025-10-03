package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointConfig.
// Experimental.
type IEndpointConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEndpointConfigRef
type jsiiProxy_IEndpointConfigRef struct {
	internal.Type__constructsIConstruct
}

