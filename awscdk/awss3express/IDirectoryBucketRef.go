package awss3express

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3express/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryBucket.
// Experimental.
type IDirectoryBucketRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DirectoryBucket resource.
	// Experimental.
	DirectoryBucketRef() *DirectoryBucketReference
}

// The jsii proxy for IDirectoryBucketRef
type jsiiProxy_IDirectoryBucketRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IDirectoryBucketRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectoryBucketRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

