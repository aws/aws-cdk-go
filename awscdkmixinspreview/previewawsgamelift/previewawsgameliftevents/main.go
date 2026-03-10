package previewawsgameliftevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent",
		reflect.TypeOf((*GameLiftMatchmakingEvent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GameLiftMatchmakingEvent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent.GameLiftMatchmakingEventProps",
		reflect.TypeOf((*GameLiftMatchmakingEvent_GameLiftMatchmakingEventProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent.GameSessionInfo",
		reflect.TypeOf((*GameLiftMatchmakingEvent_GameSessionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent.GameSessionInfoItem",
		reflect.TypeOf((*GameLiftMatchmakingEvent_GameSessionInfoItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent.RuleEvaluationMetrics",
		reflect.TypeOf((*GameLiftMatchmakingEvent_RuleEvaluationMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftMatchmakingEvent.Ticket",
		reflect.TypeOf((*GameLiftMatchmakingEvent_Ticket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftQueuePlacementEvent",
		reflect.TypeOf((*GameLiftQueuePlacementEvent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GameLiftQueuePlacementEvent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftQueuePlacementEvent.GameLiftQueuePlacementEventItem",
		reflect.TypeOf((*GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_gamelift.events.GameLiftQueuePlacementEvent.GameLiftQueuePlacementEventProps",
		reflect.TypeOf((*GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventProps)(nil)).Elem(),
	)
}
