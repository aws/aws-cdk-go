package interfacesawscontroltower

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledBaseline.
// Experimental.
type IEnabledBaselineRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnabledBaseline resource.
	// Experimental.
	EnabledBaselineRef() *EnabledBaselineReference
}

// The jsii proxy for IEnabledBaselineRef
type jsiiProxy_IEnabledBaselineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEnabledBaselineRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEnabledBaselineRef) EnabledBaselineRef() *EnabledBaselineReference {
	var returns *EnabledBaselineReference
	_jsii_.Get(
		j,
		"enabledBaselineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnabledBaselineRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnabledBaselineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

