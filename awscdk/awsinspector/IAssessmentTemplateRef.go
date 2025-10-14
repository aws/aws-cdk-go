package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssessmentTemplate.
// Experimental.
type IAssessmentTemplateRef interface {
	constructs.IConstruct
	// A reference to a AssessmentTemplate resource.
	// Experimental.
	AssessmentTemplateRef() *AssessmentTemplateReference
}

// The jsii proxy for IAssessmentTemplateRef
type jsiiProxy_IAssessmentTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssessmentTemplateRef) AssessmentTemplateRef() *AssessmentTemplateReference {
	var returns *AssessmentTemplateReference
	_jsii_.Get(
		j,
		"assessmentTemplateRef",
		&returns,
	)
	return returns
}

