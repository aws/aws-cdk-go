package interfacesawscontroltower

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledControl.
// Experimental.
type IEnabledControlRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnabledControl resource.
	// Experimental.
	EnabledControlRef() *EnabledControlReference
}

// The jsii proxy for IEnabledControlRef
type jsiiProxy_IEnabledControlRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEnabledControlRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEnabledControlRef) EnabledControlRef() *EnabledControlReference {
	var returns *EnabledControlReference
	_jsii_.Get(
		j,
		"enabledControlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnabledControlRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnabledControlRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

