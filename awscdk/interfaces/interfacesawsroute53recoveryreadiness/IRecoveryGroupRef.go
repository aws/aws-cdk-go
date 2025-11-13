package interfacesawsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecoveryGroup.
// Experimental.
type IRecoveryGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RecoveryGroup resource.
	// Experimental.
	RecoveryGroupRef() *RecoveryGroupReference
}

// The jsii proxy for IRecoveryGroupRef
type jsiiProxy_IRecoveryGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRecoveryGroupRef) RecoveryGroupRef() *RecoveryGroupReference {
	var returns *RecoveryGroupReference
	_jsii_.Get(
		j,
		"recoveryGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecoveryGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecoveryGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

