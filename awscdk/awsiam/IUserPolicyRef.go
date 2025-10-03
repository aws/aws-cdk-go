package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPolicy.
// Experimental.
type IUserPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPolicyRef
type jsiiProxy_IUserPolicyRef struct {
	internal.Type__constructsIConstruct
}

