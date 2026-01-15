package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Use an Amazon EventBridge event bus as an event destination.
//
// Example:
//   bus := events.EventBus_FromEventBusName(this, jsii.String("MyEventBus"), jsii.String("default"))
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []Action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []ActionPoint{
//   				appconfig.ActionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewEventBridgeDestination(bus),
//   		}),
//   	},
//   })
//
type EventBridgeDestination interface {
	IEventDestination
	// The URI of the extension event destination.
	ExtensionUri() *string
	// The type of the extension event destination.
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


func NewEventBridgeDestination(bus interfacesawsevents.IEventBusRef) EventBridgeDestination {
	_init_.Initialize()

	if err := validateNewEventBridgeDestinationParameters(bus); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgeDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.EventBridgeDestination",
		[]interface{}{bus},
		&j,
	)

	return &j
}

func NewEventBridgeDestination_Override(e EventBridgeDestination, bus interfacesawsevents.IEventBusRef) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.EventBridgeDestination",
		[]interface{}{bus},
		e,
	)
}

