package interfacesawsredshiftserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecoveryPoint.
// Experimental.
type IRecoveryPointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RecoveryPoint resource.
	// Experimental.
	RecoveryPointRef() *RecoveryPointReference
}

// The jsii proxy for IRecoveryPointRef
type jsiiProxy_IRecoveryPointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRecoveryPointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRecoveryPointRef) RecoveryPointRef() *RecoveryPointReference {
	var returns *RecoveryPointReference
	_jsii_.Get(
		j,
		"recoveryPointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecoveryPointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecoveryPointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

