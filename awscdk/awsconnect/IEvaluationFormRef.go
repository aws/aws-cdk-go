package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EvaluationForm.
// Experimental.
type IEvaluationFormRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEvaluationFormRef
type jsiiProxy_IEvaluationFormRef struct {
	internal.Type__constructsIConstruct
}

