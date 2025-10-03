package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPolicy.
// Experimental.
type IAccessPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessPolicyRef
type jsiiProxy_IAccessPolicyRef struct {
	internal.Type__constructsIConstruct
}

