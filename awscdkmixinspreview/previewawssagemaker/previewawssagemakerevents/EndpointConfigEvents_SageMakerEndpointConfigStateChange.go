package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.sagemaker@SageMakerEndpointConfigStateChange event types for EndpointConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerEndpointConfigStateChange := #error#.NewSageMakerEndpointConfigStateChange()
//
// Experimental.
type EndpointConfigEvents_SageMakerEndpointConfigStateChange interface {
}

// The jsii proxy struct for EndpointConfigEvents_SageMakerEndpointConfigStateChange
type jsiiProxy_EndpointConfigEvents_SageMakerEndpointConfigStateChange struct {
	_ byte // padding
}

// Experimental.
func NewEndpointConfigEvents_SageMakerEndpointConfigStateChange() EndpointConfigEvents_SageMakerEndpointConfigStateChange {
	_init_.Initialize()

	j := jsiiProxy_EndpointConfigEvents_SageMakerEndpointConfigStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents.SageMakerEndpointConfigStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEndpointConfigEvents_SageMakerEndpointConfigStateChange_Override(e EndpointConfigEvents_SageMakerEndpointConfigStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents.SageMakerEndpointConfigStateChange",
		nil, // no parameters
		e,
	)
}

