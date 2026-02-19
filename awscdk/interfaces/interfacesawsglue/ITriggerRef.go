package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Trigger.
// Experimental.
type ITriggerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Trigger resource.
	// Experimental.
	TriggerRef() *TriggerReference
}

// The jsii proxy for ITriggerRef
type jsiiProxy_ITriggerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITriggerRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITriggerRef) TriggerRef() *TriggerReference {
	var returns *TriggerReference
	_jsii_.Get(
		j,
		"triggerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITriggerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITriggerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

