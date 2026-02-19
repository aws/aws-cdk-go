package interfacesawsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComponentType.
// Experimental.
type IComponentTypeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ComponentType resource.
	// Experimental.
	ComponentTypeRef() *ComponentTypeReference
}

// The jsii proxy for IComponentTypeRef
type jsiiProxy_IComponentTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IComponentTypeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IComponentTypeRef) ComponentTypeRef() *ComponentTypeReference {
	var returns *ComponentTypeReference
	_jsii_.Get(
		j,
		"componentTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentTypeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

