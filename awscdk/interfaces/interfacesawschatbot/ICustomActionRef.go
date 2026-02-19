package interfacesawschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomAction.
// Experimental.
type ICustomActionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CustomAction resource.
	// Experimental.
	CustomActionRef() *CustomActionReference
}

// The jsii proxy for ICustomActionRef
type jsiiProxy_ICustomActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICustomActionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICustomActionRef) CustomActionRef() *CustomActionReference {
	var returns *CustomActionReference
	_jsii_.Get(
		j,
		"customActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomActionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

