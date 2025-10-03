package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CachePolicy.
// Experimental.
type ICachePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICachePolicyRef
type jsiiProxy_ICachePolicyRef struct {
	internal.Type__constructsIConstruct
}

