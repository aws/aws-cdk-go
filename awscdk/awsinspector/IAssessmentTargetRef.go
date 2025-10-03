package awsinspector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssessmentTarget.
// Experimental.
type IAssessmentTargetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssessmentTargetRef
type jsiiProxy_IAssessmentTargetRef struct {
	internal.Type__constructsIConstruct
}

