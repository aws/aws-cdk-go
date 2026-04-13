package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerEndpointStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerEndpointStateChange := awscdkmixinspreview.Events.NewSageMakerEndpointStateChange()
//
// Experimental.
type SageMakerEndpointStateChange interface {
}

// The jsii proxy struct for SageMakerEndpointStateChange
type jsiiProxy_SageMakerEndpointStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerEndpointStateChange() SageMakerEndpointStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerEndpointStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerEndpointStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerEndpointStateChange_Override(s SageMakerEndpointStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerEndpointStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Endpoint State Change.
// Experimental.
func SageMakerEndpointStateChange_EventPattern(options *SageMakerEndpointStateChange_SageMakerEndpointStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerEndpointStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerEndpointStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

