package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Outcome.
// Experimental.
type IOutcomeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOutcomeRef
type jsiiProxy_IOutcomeRef struct {
	internal.Type__constructsIConstruct
}

