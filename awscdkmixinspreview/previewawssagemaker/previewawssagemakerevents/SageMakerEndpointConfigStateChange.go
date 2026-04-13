package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerEndpointConfigStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerEndpointConfigStateChange := awscdkmixinspreview.Events.NewSageMakerEndpointConfigStateChange()
//
// Experimental.
type SageMakerEndpointConfigStateChange interface {
}

// The jsii proxy struct for SageMakerEndpointConfigStateChange
type jsiiProxy_SageMakerEndpointConfigStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerEndpointConfigStateChange() SageMakerEndpointConfigStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerEndpointConfigStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerEndpointConfigStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerEndpointConfigStateChange_Override(s SageMakerEndpointConfigStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerEndpointConfigStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Endpoint Config State Change.
// Experimental.
func SageMakerEndpointConfigStateChange_EventPattern(options *SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerEndpointConfigStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerEndpointConfigStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

