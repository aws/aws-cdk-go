package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.sagemaker@SageMakerTransformJobStateChange event types for Model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTransformJobStateChange := #error#.NewSageMakerTransformJobStateChange()
//
// Experimental.
type ModelEvents_SageMakerTransformJobStateChange interface {
}

// The jsii proxy struct for ModelEvents_SageMakerTransformJobStateChange
type jsiiProxy_ModelEvents_SageMakerTransformJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewModelEvents_SageMakerTransformJobStateChange() ModelEvents_SageMakerTransformJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_ModelEvents_SageMakerTransformJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewModelEvents_SageMakerTransformJobStateChange_Override(m ModelEvents_SageMakerTransformJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange",
		nil, // no parameters
		m,
	)
}

