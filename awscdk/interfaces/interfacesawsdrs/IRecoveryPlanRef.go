package interfacesawsdrs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdrs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecoveryPlan.
// Experimental.
type IRecoveryPlanRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RecoveryPlan resource.
	// Experimental.
	RecoveryPlanRef() *RecoveryPlanReference
}

// The jsii proxy for IRecoveryPlanRef
type jsiiProxy_IRecoveryPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRecoveryPlanRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRecoveryPlanRef) RecoveryPlanRef() *RecoveryPlanReference {
	var returns *RecoveryPlanReference
	_jsii_.Get(
		j,
		"recoveryPlanRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecoveryPlanRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecoveryPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

