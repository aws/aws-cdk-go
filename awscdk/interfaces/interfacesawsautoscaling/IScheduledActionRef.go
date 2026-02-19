package interfacesawsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledAction.
// Experimental.
type IScheduledActionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ScheduledAction resource.
	// Experimental.
	ScheduledActionRef() *ScheduledActionReference
}

// The jsii proxy for IScheduledActionRef
type jsiiProxy_IScheduledActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IScheduledActionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IScheduledActionRef) ScheduledActionRef() *ScheduledActionReference {
	var returns *ScheduledActionReference
	_jsii_.Get(
		j,
		"scheduledActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduledActionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduledActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

