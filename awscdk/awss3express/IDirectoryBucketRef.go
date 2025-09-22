package awss3express

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3express/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryBucket.
// Experimental.
type IDirectoryBucketRef interface {
	constructs.IConstruct
	// A reference to a DirectoryBucket resource.
	// Experimental.
	DirectoryBucketRef() *DirectoryBucketReference
}

// The jsii proxy for IDirectoryBucketRef
type jsiiProxy_IDirectoryBucketRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDirectoryBucketRef) DirectoryBucketRef() *DirectoryBucketReference {
	var returns *DirectoryBucketReference
	_jsii_.Get(
		j,
		"directoryBucketRef",
		&returns,
	)
	return returns
}

