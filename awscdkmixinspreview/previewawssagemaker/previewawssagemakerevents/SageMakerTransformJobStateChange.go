package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerTransformJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTransformJobStateChange := awscdkmixinspreview.Events.NewSageMakerTransformJobStateChange()
//
// Experimental.
type SageMakerTransformJobStateChange interface {
}

// The jsii proxy struct for SageMakerTransformJobStateChange
type jsiiProxy_SageMakerTransformJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerTransformJobStateChange() SageMakerTransformJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerTransformJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerTransformJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerTransformJobStateChange_Override(s SageMakerTransformJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerTransformJobStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Transform Job State Change.
// Experimental.
func SageMakerTransformJobStateChange_SageMakerTransformJobStateChangePattern(options *SageMakerTransformJobStateChange_SageMakerTransformJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerTransformJobStateChange_SageMakerTransformJobStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerTransformJobStateChange",
		"sageMakerTransformJobStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

