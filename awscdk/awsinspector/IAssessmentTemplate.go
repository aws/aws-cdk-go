package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
)

// Interface for an Inspector Assessment Template.
type IAssessmentTemplate interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the assessment template.
	AssessmentTemplateArn() *string
}

// The jsii proxy for IAssessmentTemplate
type jsiiProxy_IAssessmentTemplate struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAssessmentTemplate) AssessmentTemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assessmentTemplateArn",
		&returns,
	)
	return returns
}

