package awsfis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExperimentTemplate.
// Experimental.
type IExperimentTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IExperimentTemplateRef
type jsiiProxy_IExperimentTemplateRef struct {
	internal.Type__constructsIConstruct
}

