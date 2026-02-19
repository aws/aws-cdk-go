package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceAction.
// Experimental.
type IServiceActionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceAction resource.
	// Experimental.
	ServiceActionRef() *ServiceActionReference
}

// The jsii proxy for IServiceActionRef
type jsiiProxy_IServiceActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServiceActionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServiceActionRef) ServiceActionRef() *ServiceActionReference {
	var returns *ServiceActionReference
	_jsii_.Get(
		j,
		"serviceActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceActionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

