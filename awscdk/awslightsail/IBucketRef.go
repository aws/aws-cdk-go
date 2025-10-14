package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bucket.
// Experimental.
type IBucketRef interface {
	constructs.IConstruct
	// A reference to a Bucket resource.
	// Experimental.
	BucketRef() *BucketReference
}

// The jsii proxy for IBucketRef
type jsiiProxy_IBucketRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBucketRef) BucketRef() *BucketReference {
	var returns *BucketReference
	_jsii_.Get(
		j,
		"bucketRef",
		&returns,
	)
	return returns
}

