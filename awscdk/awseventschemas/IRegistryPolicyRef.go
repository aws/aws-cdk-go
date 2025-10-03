package awseventschemas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryPolicy.
// Experimental.
type IRegistryPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRegistryPolicyRef
type jsiiProxy_IRegistryPolicyRef struct {
	internal.Type__constructsIConstruct
}

