package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker"
)

// EventBridge event patterns for EndpointConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var endpointConfigRef IEndpointConfigRef
//
//   endpointConfigEvents := awscdkmixinspreview.Events.EndpointConfigEvents_FromEndpointConfig(endpointConfigRef)
//
// Experimental.
type EndpointConfigEvents interface {
	// EventBridge event pattern for EndpointConfig SageMaker Endpoint Config State Change.
	// Experimental.
	SageMakerEndpointConfigStateChangePattern(options *EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for EndpointConfigEvents
type jsiiProxy_EndpointConfigEvents struct {
	_ byte // padding
}

// Create EndpointConfigEvents from a EndpointConfig reference.
// Experimental.
func EndpointConfigEvents_FromEndpointConfig(endpointConfigRef interfacesawssagemaker.IEndpointConfigRef) EndpointConfigEvents {
	_init_.Initialize()

	if err := validateEndpointConfigEvents_FromEndpointConfigParameters(endpointConfigRef); err != nil {
		panic(err)
	}
	var returns EndpointConfigEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents",
		"fromEndpointConfig",
		[]interface{}{endpointConfigRef},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointConfigEvents) SageMakerEndpointConfigStateChangePattern(options *EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps) *awsevents.EventPattern {
	if err := e.validateSageMakerEndpointConfigStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		e,
		"sageMakerEndpointConfigStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

