package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerHyperParameterTuningJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerHyperParameterTuningJobStateChange := awscdkmixinspreview.Events.NewSageMakerHyperParameterTuningJobStateChange()
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange interface {
}

// The jsii proxy struct for SageMakerHyperParameterTuningJobStateChange
type jsiiProxy_SageMakerHyperParameterTuningJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerHyperParameterTuningJobStateChange() SageMakerHyperParameterTuningJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerHyperParameterTuningJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerHyperParameterTuningJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerHyperParameterTuningJobStateChange_Override(s SageMakerHyperParameterTuningJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerHyperParameterTuningJobStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker HyperParameter Tuning Job State Change.
// Experimental.
func SageMakerHyperParameterTuningJobStateChange_EventPattern(options *SageMakerHyperParameterTuningJobStateChange_SageMakerHyperParameterTuningJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerHyperParameterTuningJobStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerHyperParameterTuningJobStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

