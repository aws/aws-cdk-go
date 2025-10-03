package awsinspector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssessmentTemplate.
// Experimental.
type IAssessmentTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssessmentTemplateRef
type jsiiProxy_IAssessmentTemplateRef struct {
	internal.Type__constructsIConstruct
}

