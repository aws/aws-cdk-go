package interfacesawsfis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SafetyLever.
// Experimental.
type ISafetyLeverRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SafetyLever resource.
	// Experimental.
	SafetyLeverRef() *SafetyLeverReference
}

// The jsii proxy for ISafetyLeverRef
type jsiiProxy_ISafetyLeverRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISafetyLeverRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISafetyLeverRef) SafetyLeverRef() *SafetyLeverReference {
	var returns *SafetyLeverReference
	_jsii_.Get(
		j,
		"safetyLeverRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISafetyLeverRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISafetyLeverRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

