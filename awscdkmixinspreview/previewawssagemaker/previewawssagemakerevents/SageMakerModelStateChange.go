package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerModelStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelStateChange := awscdkmixinspreview.Events.NewSageMakerModelStateChange()
//
// Experimental.
type SageMakerModelStateChange interface {
}

// The jsii proxy struct for SageMakerModelStateChange
type jsiiProxy_SageMakerModelStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerModelStateChange() SageMakerModelStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerModelStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerModelStateChange_Override(s SageMakerModelStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Model State Change.
// Experimental.
func SageMakerModelStateChange_EventPattern(options *SageMakerModelStateChange_SageMakerModelStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerModelStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

