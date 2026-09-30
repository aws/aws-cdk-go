package interfacesawsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DelegationSet.
// Experimental.
type IDelegationSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DelegationSet resource.
	// Experimental.
	DelegationSetRef() *DelegationSetReference
}

// The jsii proxy for IDelegationSetRef
type jsiiProxy_IDelegationSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDelegationSetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDelegationSetRef) DelegationSetRef() *DelegationSetReference {
	var returns *DelegationSetReference
	_jsii_.Get(
		j,
		"delegationSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelegationSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelegationSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

