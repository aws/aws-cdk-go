package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ActionTarget.
// Experimental.
type IActionTargetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ActionTarget resource.
	// Experimental.
	ActionTargetRef() *ActionTargetReference
}

// The jsii proxy for IActionTargetRef
type jsiiProxy_IActionTargetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IActionTargetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IActionTargetRef) ActionTargetRef() *ActionTargetReference {
	var returns *ActionTargetReference
	_jsii_.Get(
		j,
		"actionTargetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActionTargetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActionTargetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

