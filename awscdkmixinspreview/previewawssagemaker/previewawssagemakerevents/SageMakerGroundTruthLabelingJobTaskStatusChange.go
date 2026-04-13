package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerGroundTruthLabelingJobTaskStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerGroundTruthLabelingJobTaskStatusChange := awscdkmixinspreview.Events.NewSageMakerGroundTruthLabelingJobTaskStatusChange()
//
// Experimental.
type SageMakerGroundTruthLabelingJobTaskStatusChange interface {
}

// The jsii proxy struct for SageMakerGroundTruthLabelingJobTaskStatusChange
type jsiiProxy_SageMakerGroundTruthLabelingJobTaskStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerGroundTruthLabelingJobTaskStatusChange() SageMakerGroundTruthLabelingJobTaskStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerGroundTruthLabelingJobTaskStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerGroundTruthLabelingJobTaskStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerGroundTruthLabelingJobTaskStatusChange_Override(s SageMakerGroundTruthLabelingJobTaskStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerGroundTruthLabelingJobTaskStatusChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Ground Truth Labeling Job Task Status Change.
// Experimental.
func SageMakerGroundTruthLabelingJobTaskStatusChange_EventPattern(options *SageMakerGroundTruthLabelingJobTaskStatusChange_SageMakerGroundTruthLabelingJobTaskStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerGroundTruthLabelingJobTaskStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerGroundTruthLabelingJobTaskStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

