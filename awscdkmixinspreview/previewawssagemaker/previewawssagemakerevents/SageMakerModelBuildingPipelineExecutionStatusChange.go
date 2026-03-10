package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerModelBuildingPipelineExecutionStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelBuildingPipelineExecutionStatusChange := awscdkmixinspreview.Events.NewSageMakerModelBuildingPipelineExecutionStatusChange()
//
// Experimental.
type SageMakerModelBuildingPipelineExecutionStatusChange interface {
}

// The jsii proxy struct for SageMakerModelBuildingPipelineExecutionStatusChange
type jsiiProxy_SageMakerModelBuildingPipelineExecutionStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerModelBuildingPipelineExecutionStatusChange() SageMakerModelBuildingPipelineExecutionStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerModelBuildingPipelineExecutionStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelBuildingPipelineExecutionStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerModelBuildingPipelineExecutionStatusChange_Override(s SageMakerModelBuildingPipelineExecutionStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelBuildingPipelineExecutionStatusChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Model Building Pipeline Execution Status Change.
// Experimental.
func SageMakerModelBuildingPipelineExecutionStatusChange_SageMakerModelBuildingPipelineExecutionStatusChangePattern(options *SageMakerModelBuildingPipelineExecutionStatusChange_SageMakerModelBuildingPipelineExecutionStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerModelBuildingPipelineExecutionStatusChange_SageMakerModelBuildingPipelineExecutionStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelBuildingPipelineExecutionStatusChange",
		"sageMakerModelBuildingPipelineExecutionStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

