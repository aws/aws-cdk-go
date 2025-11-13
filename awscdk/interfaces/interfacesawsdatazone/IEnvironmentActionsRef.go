package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentActions.
// Experimental.
type IEnvironmentActionsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnvironmentActions resource.
	// Experimental.
	EnvironmentActionsRef() *EnvironmentActionsReference
}

// The jsii proxy for IEnvironmentActionsRef
type jsiiProxy_IEnvironmentActionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEnvironmentActionsRef) EnvironmentActionsRef() *EnvironmentActionsReference {
	var returns *EnvironmentActionsReference
	_jsii_.Get(
		j,
		"environmentActionsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentActionsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentActionsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

