package awss3express

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3express/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryBucket.
// Experimental.
type IDirectoryBucketRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDirectoryBucketRef
type jsiiProxy_IDirectoryBucketRef struct {
	internal.Type__constructsIConstruct
}

