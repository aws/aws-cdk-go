package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PredefinedAttribute.
// Experimental.
type IPredefinedAttributeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPredefinedAttributeRef
type jsiiProxy_IPredefinedAttributeRef struct {
	internal.Type__constructsIConstruct
}

