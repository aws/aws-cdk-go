package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Outcome.
// Experimental.
type IOutcomeRef interface {
	constructs.IConstruct
	// A reference to a Outcome resource.
	// Experimental.
	OutcomeRef() *OutcomeReference
}

// The jsii proxy for IOutcomeRef
type jsiiProxy_IOutcomeRef struct {
	internal.Type__constructsIConstruct
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

