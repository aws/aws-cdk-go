package previewawsgameliftevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.gamelift@GameLiftMatchmakingEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameLiftMatchmakingEvent := awscdkmixinspreview.Events.NewGameLiftMatchmakingEvent()
//
// Experimental.
type GameLiftMatchmakingEvent interface {
}

// The jsii proxy struct for GameLiftMatchmakingEvent
type jsiiProxy_GameLiftMatchmakingEvent struct {
	_ byte // padding
}

// Experimental.
func NewGameLiftMatchmakingEvent() GameLiftMatchmakingEvent {
	_init_.Initialize()

	j := jsiiProxy_GameLiftMatchmakingEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGameLiftMatchmakingEvent_Override(g GameLiftMatchmakingEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for GameLift Matchmaking Event.
// Experimental.
func GameLiftMatchmakingEvent_EventPattern(options *GameLiftMatchmakingEvent_GameLiftMatchmakingEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGameLiftMatchmakingEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

