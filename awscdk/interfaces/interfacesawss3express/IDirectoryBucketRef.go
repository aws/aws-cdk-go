package interfacesawss3express

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3express/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryBucket.
// Experimental.
type IDirectoryBucketRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DirectoryBucket resource.
	// Experimental.
	DirectoryBucketRef() *DirectoryBucketReference
}

// The jsii proxy for IDirectoryBucketRef
type jsiiProxy_IDirectoryBucketRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDirectoryBucketRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IDirectoryBucketRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

