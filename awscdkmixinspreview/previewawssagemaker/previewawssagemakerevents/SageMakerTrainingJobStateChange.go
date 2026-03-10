package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerTrainingJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTrainingJobStateChange := awscdkmixinspreview.Events.NewSageMakerTrainingJobStateChange()
//
// Experimental.
type SageMakerTrainingJobStateChange interface {
}

// The jsii proxy struct for SageMakerTrainingJobStateChange
type jsiiProxy_SageMakerTrainingJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerTrainingJobStateChange() SageMakerTrainingJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerTrainingJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerTrainingJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerTrainingJobStateChange_Override(s SageMakerTrainingJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerTrainingJobStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Training Job State Change.
// Experimental.
func SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangePattern(options *SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerTrainingJobStateChange",
		"sageMakerTrainingJobStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

