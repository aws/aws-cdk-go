package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsinspector"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for an Inspector Assessment Template.
type IAssessmentTemplate interface {
	interfacesawsinspector.IAssessmentTemplateRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the assessment template.
	AssessmentTemplateArn() *string
}

// The jsii proxy for IAssessmentTemplate
type jsiiProxy_IAssessmentTemplate struct {
	internal.Type__interfacesawsinspectorIAssessmentTemplateRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAssessmentTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IAssessmentTemplate) AssessmentTemplateRef() *interfacesawsinspector.AssessmentTemplateReference {
	var returns *interfacesawsinspector.AssessmentTemplateReference
	_jsii_.Get(
		j,
		"assessmentTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssessmentTemplate) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssessmentTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssessmentTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

