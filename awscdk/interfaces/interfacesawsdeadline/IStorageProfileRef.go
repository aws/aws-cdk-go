package interfacesawsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageProfile.
// Experimental.
type IStorageProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StorageProfile resource.
	// Experimental.
	StorageProfileRef() *StorageProfileReference
}

// The jsii proxy for IStorageProfileRef
type jsiiProxy_IStorageProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStorageProfileRef) StorageProfileRef() *StorageProfileReference {
	var returns *StorageProfileReference
	_jsii_.Get(
		j,
		"storageProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

