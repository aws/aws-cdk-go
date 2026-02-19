package interfacesawsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Outcome.
// Experimental.
type IOutcomeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Outcome resource.
	// Experimental.
	OutcomeRef() *OutcomeReference
}

// The jsii proxy for IOutcomeRef
type jsiiProxy_IOutcomeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOutcomeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOutcomeRef) OutcomeRef() *OutcomeReference {
	var returns *OutcomeReference
	_jsii_.Get(
		j,
		"outcomeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutcomeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutcomeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

