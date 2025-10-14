package awss3outposts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3outposts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BucketPolicy.
// Experimental.
type IBucketPolicyRef interface {
	constructs.IConstruct
	// A reference to a BucketPolicy resource.
	// Experimental.
	BucketPolicyRef() *BucketPolicyReference
}

// The jsii proxy for IBucketPolicyRef
type jsiiProxy_IBucketPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBucketPolicyRef) BucketPolicyRef() *BucketPolicyReference {
	var returns *BucketPolicyReference
	_jsii_.Get(
		j,
		"bucketPolicyRef",
		&returns,
	)
	return returns
}

