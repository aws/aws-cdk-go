package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an Event Bridge event bus as a Lambda destination.
//
// If no event bus is specified, the default event bus is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventBus eventBus
//
//   eventBridgeDestination := awscdk.Aws_lambda_destinations.NewEventBridgeDestination(eventBus)
//
type EventBridgeDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for EventBridgeDestination
type jsiiProxy_EventBridgeDestination struct {
	internal.Type__awslambdaIDestination
}

func NewEventBridgeDestination(eventBus awsevents.IEventBus) EventBridgeDestination {
	_init_.Initialize()

	j := jsiiProxy_EventBridgeDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.EventBridgeDestination",
		[]interface{}{eventBus},
		&j,
	)

	return &j
}

func NewEventBridgeDestination_Override(e EventBridgeDestination, eventBus awsevents.IEventBus) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.EventBridgeDestination",
		[]interface{}{eventBus},
		e,
	)
}

func (e *jsiiProxy_EventBridgeDestination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

