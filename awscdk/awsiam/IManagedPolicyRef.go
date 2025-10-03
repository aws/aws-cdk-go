package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedPolicy.
// Experimental.
type IManagedPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IManagedPolicyRef
type jsiiProxy_IManagedPolicyRef struct {
	internal.Type__constructsIConstruct
}

