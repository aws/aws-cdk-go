package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupPolicy.
// Experimental.
type IGroupPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGroupPolicyRef
type jsiiProxy_IGroupPolicyRef struct {
	internal.Type__constructsIConstruct
}

