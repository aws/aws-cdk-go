package interfacesawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageTierPolicy.
// Experimental.
type IStorageTierPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StorageTierPolicy resource.
	// Experimental.
	StorageTierPolicyRef() *StorageTierPolicyReference
}

// The jsii proxy for IStorageTierPolicyRef
type jsiiProxy_IStorageTierPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IStorageTierPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IStorageTierPolicyRef) StorageTierPolicyRef() *StorageTierPolicyReference {
	var returns *StorageTierPolicyReference
	_jsii_.Get(
		j,
		"storageTierPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageTierPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageTierPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

