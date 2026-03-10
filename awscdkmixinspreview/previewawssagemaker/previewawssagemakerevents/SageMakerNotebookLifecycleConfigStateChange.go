package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerNotebookLifecycleConfigStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerNotebookLifecycleConfigStateChange := awscdkmixinspreview.Events.NewSageMakerNotebookLifecycleConfigStateChange()
//
// Experimental.
type SageMakerNotebookLifecycleConfigStateChange interface {
}

// The jsii proxy struct for SageMakerNotebookLifecycleConfigStateChange
type jsiiProxy_SageMakerNotebookLifecycleConfigStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerNotebookLifecycleConfigStateChange() SageMakerNotebookLifecycleConfigStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerNotebookLifecycleConfigStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerNotebookLifecycleConfigStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerNotebookLifecycleConfigStateChange_Override(s SageMakerNotebookLifecycleConfigStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerNotebookLifecycleConfigStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Notebook Lifecycle Config State Change.
// Experimental.
func SageMakerNotebookLifecycleConfigStateChange_SageMakerNotebookLifecycleConfigStateChangePattern(options *SageMakerNotebookLifecycleConfigStateChange_SageMakerNotebookLifecycleConfigStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerNotebookLifecycleConfigStateChange_SageMakerNotebookLifecycleConfigStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerNotebookLifecycleConfigStateChange",
		"sageMakerNotebookLifecycleConfigStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

