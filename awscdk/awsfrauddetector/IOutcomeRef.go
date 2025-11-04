package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Outcome.
// Experimental.
type IOutcomeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Outcome resource.
	// Experimental.
	OutcomeRef() *OutcomeReference
}

// The jsii proxy for IOutcomeRef
type jsiiProxy_IOutcomeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IOutcomeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

