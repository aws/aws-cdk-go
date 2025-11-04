package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageLens.
// Experimental.
type IStorageLensRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StorageLens resource.
	// Experimental.
	StorageLensRef() *StorageLensReference
}

// The jsii proxy for IStorageLensRef
type jsiiProxy_IStorageLensRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStorageLensRef) StorageLensRef() *StorageLensReference {
	var returns *StorageLensReference
	_jsii_.Get(
		j,
		"storageLensRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageLensRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageLensRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

