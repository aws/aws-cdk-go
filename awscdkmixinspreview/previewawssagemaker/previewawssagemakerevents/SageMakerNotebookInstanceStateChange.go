package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerNotebookInstanceStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerNotebookInstanceStateChange := awscdkmixinspreview.Events.NewSageMakerNotebookInstanceStateChange()
//
// Experimental.
type SageMakerNotebookInstanceStateChange interface {
}

// The jsii proxy struct for SageMakerNotebookInstanceStateChange
type jsiiProxy_SageMakerNotebookInstanceStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerNotebookInstanceStateChange() SageMakerNotebookInstanceStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerNotebookInstanceStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerNotebookInstanceStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerNotebookInstanceStateChange_Override(s SageMakerNotebookInstanceStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerNotebookInstanceStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Notebook Instance State Change.
// Experimental.
func SageMakerNotebookInstanceStateChange_SageMakerNotebookInstanceStateChangePattern(options *SageMakerNotebookInstanceStateChange_SageMakerNotebookInstanceStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerNotebookInstanceStateChange_SageMakerNotebookInstanceStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerNotebookInstanceStateChange",
		"sageMakerNotebookInstanceStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

