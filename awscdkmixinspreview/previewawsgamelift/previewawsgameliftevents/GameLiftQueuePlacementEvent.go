package previewawsgameliftevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.gamelift@GameLiftQueuePlacementEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameLiftQueuePlacementEvent := awscdkmixinspreview.Events.NewGameLiftQueuePlacementEvent()
//
// Experimental.
type GameLiftQueuePlacementEvent interface {
}

// The jsii proxy struct for GameLiftQueuePlacementEvent
type jsiiProxy_GameLiftQueuePlacementEvent struct {
	_ byte // padding
}

// Experimental.
func NewGameLiftQueuePlacementEvent() GameLiftQueuePlacementEvent {
	_init_.Initialize()

	j := jsiiProxy_GameLiftQueuePlacementEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftQueuePlacementEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGameLiftQueuePlacementEvent_Override(g GameLiftQueuePlacementEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftQueuePlacementEvent",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for GameLift Queue Placement Event.
// Experimental.
func GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventPattern(options *GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGameLiftQueuePlacementEvent_GameLiftQueuePlacementEventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftQueuePlacementEvent",
		"gameLiftQueuePlacementEventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

