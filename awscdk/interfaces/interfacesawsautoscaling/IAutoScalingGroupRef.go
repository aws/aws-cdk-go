package interfacesawsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoScalingGroup.
// Experimental.
type IAutoScalingGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AutoScalingGroup resource.
	// Experimental.
	AutoScalingGroupRef() *AutoScalingGroupReference
}

// The jsii proxy for IAutoScalingGroupRef
type jsiiProxy_IAutoScalingGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAutoScalingGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAutoScalingGroupRef) AutoScalingGroupRef() *AutoScalingGroupReference {
	var returns *AutoScalingGroupReference
	_jsii_.Get(
		j,
		"autoScalingGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

