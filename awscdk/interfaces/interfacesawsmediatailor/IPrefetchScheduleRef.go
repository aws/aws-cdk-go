package interfacesawsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrefetchSchedule.
// Experimental.
type IPrefetchScheduleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PrefetchSchedule resource.
	// Experimental.
	PrefetchScheduleRef() *PrefetchScheduleReference
}

// The jsii proxy for IPrefetchScheduleRef
type jsiiProxy_IPrefetchScheduleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPrefetchScheduleRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPrefetchScheduleRef) PrefetchScheduleRef() *PrefetchScheduleReference {
	var returns *PrefetchScheduleReference
	_jsii_.Get(
		j,
		"prefetchScheduleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefetchScheduleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefetchScheduleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

