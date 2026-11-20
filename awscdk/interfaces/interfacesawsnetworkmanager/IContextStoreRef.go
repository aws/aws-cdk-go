package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContextStore.
// Experimental.
type IContextStoreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContextStore resource.
	// Experimental.
	ContextStoreRef() *ContextStoreReference
}

// The jsii proxy for IContextStoreRef
type jsiiProxy_IContextStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IContextStoreRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IContextStoreRef) ContextStoreRef() *ContextStoreReference {
	var returns *ContextStoreReference
	_jsii_.Get(
		j,
		"contextStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextStoreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

