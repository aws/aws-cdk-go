package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssessmentTemplate.
// Experimental.
type IAssessmentTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AssessmentTemplate resource.
	// Experimental.
	AssessmentTemplateRef() *AssessmentTemplateReference
}

// The jsii proxy for IAssessmentTemplateRef
type jsiiProxy_IAssessmentTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAssessmentTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssessmentTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

