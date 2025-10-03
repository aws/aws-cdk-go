package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Step.
// Experimental.
type IStepRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStepRef
type jsiiProxy_IStepRef struct {
	internal.Type__constructsIConstruct
}

