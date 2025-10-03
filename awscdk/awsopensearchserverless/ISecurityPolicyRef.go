package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityPolicy.
// Experimental.
type ISecurityPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecurityPolicyRef
type jsiiProxy_ISecurityPolicyRef struct {
	internal.Type__constructsIConstruct
}

