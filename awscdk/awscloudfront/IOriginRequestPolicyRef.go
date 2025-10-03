package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginRequestPolicy.
// Experimental.
type IOriginRequestPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOriginRequestPolicyRef
type jsiiProxy_IOriginRequestPolicyRef struct {
	internal.Type__constructsIConstruct
}

