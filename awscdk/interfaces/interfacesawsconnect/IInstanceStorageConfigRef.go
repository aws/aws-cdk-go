package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceStorageConfig.
// Experimental.
type IInstanceStorageConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InstanceStorageConfig resource.
	// Experimental.
	InstanceStorageConfigRef() *InstanceStorageConfigReference
}

// The jsii proxy for IInstanceStorageConfigRef
type jsiiProxy_IInstanceStorageConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInstanceStorageConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInstanceStorageConfigRef) InstanceStorageConfigRef() *InstanceStorageConfigReference {
	var returns *InstanceStorageConfigReference
	_jsii_.Get(
		j,
		"instanceStorageConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceStorageConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceStorageConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

