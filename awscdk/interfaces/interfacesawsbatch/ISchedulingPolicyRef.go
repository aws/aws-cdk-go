package interfacesawsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchedulingPolicy.
// Experimental.
type ISchedulingPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SchedulingPolicy resource.
	// Experimental.
	SchedulingPolicyRef() *SchedulingPolicyReference
}

// The jsii proxy for ISchedulingPolicyRef
type jsiiProxy_ISchedulingPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISchedulingPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISchedulingPolicyRef) SchedulingPolicyRef() *SchedulingPolicyReference {
	var returns *SchedulingPolicyReference
	_jsii_.Get(
		j,
		"schedulingPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedulingPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedulingPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

