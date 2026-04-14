package interfacesawsbraket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbraket/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SpendingLimit.
// Experimental.
type ISpendingLimitRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SpendingLimit resource.
	// Experimental.
	SpendingLimitRef() *SpendingLimitReference
}

// The jsii proxy for ISpendingLimitRef
type jsiiProxy_ISpendingLimitRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISpendingLimitRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISpendingLimitRef) SpendingLimitRef() *SpendingLimitReference {
	var returns *SpendingLimitReference
	_jsii_.Get(
		j,
		"spendingLimitRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpendingLimitRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpendingLimitRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

