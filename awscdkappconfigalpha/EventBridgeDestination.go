package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Use an Amazon EventBridge event bus as an event destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventBus eventBus
//
//   eventBridgeDestination := appconfig_alpha.NewEventBridgeDestination(eventBus)
//
// Deprecated.
type EventBridgeDestination interface {
	IEventDestination
	// The URI of the extension event destination.
	// Deprecated.
	ExtensionUri() *string
	// The type of the extension event destination.
	// Deprecated.
	Type() SourceType
}

// The jsii proxy struct for EventBridgeDestination
type jsiiProxy_EventBridgeDestination struct {
	jsiiProxy_IEventDestination
}

func (j *jsiiProxy_EventBridgeDestination) ExtensionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeDestination) Type() SourceType {
	var returns SourceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Deprecated.
func NewEventBridgeDestination(bus awsevents.IEventBus) EventBridgeDestination {
	_init_.Initialize()

	if err := validateNewEventBridgeDestinationParameters(bus); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgeDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.EventBridgeDestination",
		[]interface{}{bus},
		&j,
	)

	return &j
}

// Deprecated.
func NewEventBridgeDestination_Override(e EventBridgeDestination, bus awsevents.IEventBus) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.EventBridgeDestination",
		[]interface{}{bus},
		e,
	)
}

