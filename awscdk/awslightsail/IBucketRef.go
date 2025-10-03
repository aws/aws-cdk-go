package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bucket.
// Experimental.
type IBucketRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBucketRef
type jsiiProxy_IBucketRef struct {
	internal.Type__constructsIConstruct
}

