package awsresiliencehub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresiliencehub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResiliencyPolicy.
// Experimental.
type IResiliencyPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResiliencyPolicyRef
type jsiiProxy_IResiliencyPolicyRef struct {
	internal.Type__constructsIConstruct
}

