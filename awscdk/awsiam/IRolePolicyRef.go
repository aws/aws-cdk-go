package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RolePolicy.
// Experimental.
type IRolePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRolePolicyRef
type jsiiProxy_IRolePolicyRef struct {
	internal.Type__constructsIConstruct
}

