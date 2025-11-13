package interfacesawss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageLensGroup.
// Experimental.
type IStorageLensGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StorageLensGroup resource.
	// Experimental.
	StorageLensGroupRef() *StorageLensGroupReference
}

// The jsii proxy for IStorageLensGroupRef
type jsiiProxy_IStorageLensGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStorageLensGroupRef) StorageLensGroupRef() *StorageLensGroupReference {
	var returns *StorageLensGroupReference
	_jsii_.Get(
		j,
		"storageLensGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageLensGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageLensGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

