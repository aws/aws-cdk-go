package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Capability.
// Experimental.
type ICapabilityRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICapabilityRef
type jsiiProxy_ICapabilityRef struct {
	internal.Type__constructsIConstruct
}

