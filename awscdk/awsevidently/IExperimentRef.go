package awsevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Experiment.
// Experimental.
type IExperimentRef interface {
	constructs.IConstruct
	// A reference to a Experiment resource.
	// Experimental.
	ExperimentRef() *ExperimentReference
}

// The jsii proxy for IExperimentRef
type jsiiProxy_IExperimentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IExperimentRef) ExperimentRef() *ExperimentReference {
	var returns *ExperimentReference
	_jsii_.Get(
		j,
		"experimentRef",
		&returns,
	)
	return returns
}

