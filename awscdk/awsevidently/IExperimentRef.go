package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Experiment.
// Experimental.
type IExperimentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IExperimentRef
type jsiiProxy_IExperimentRef struct {
	internal.Type__constructsIConstruct
}

