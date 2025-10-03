package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourcePolicy.
// Experimental.
type IResourcePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourcePolicyRef
type jsiiProxy_IResourcePolicyRef struct {
	internal.Type__constructsIConstruct
}

