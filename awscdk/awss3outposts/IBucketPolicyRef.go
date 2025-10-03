package awss3outposts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3outposts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BucketPolicy.
// Experimental.
type IBucketPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBucketPolicyRef
type jsiiProxy_IBucketPolicyRef struct {
	internal.Type__constructsIConstruct
}

