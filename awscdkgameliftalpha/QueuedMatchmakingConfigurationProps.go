package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a new queued matchmaking configuration.
//
// Example:
//   var queue gameSessionQueue
//   var ruleSet matchmakingRuleSet
//
//
//   gamelift.NewQueuedMatchmakingConfiguration(this, jsii.String("QueuedMatchmakingConfiguration"), &QueuedMatchmakingConfigurationProps{
//   	MatchmakingConfigurationName: jsii.String("test-queued-config-name"),
//   	GameSessionQueues: []iGameSessionQueue{
//   		queue,
//   	},
//   	RuleSet: ruleSet,
//   })
//
// Experimental.
type QueuedMatchmakingConfigurationProps struct {
	// A unique identifier for the matchmaking configuration.
	//
	// This name is used to identify the configuration associated with a matchmaking request or ticket.
	// Experimental.
	MatchmakingConfigurationName *string `field:"required" json:"matchmakingConfigurationName" yaml:"matchmakingConfigurationName"`
	// A matchmaking rule set to use with this configuration.
	//
	// A matchmaking configuration can only use rule sets that are defined in the same Region.
	// Experimental.
	RuleSet IMatchmakingRuleSet `field:"required" json:"ruleSet" yaml:"ruleSet"`
	// The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
	// Experimental.
	AcceptanceTimeout awscdk.Duration `field:"optional" json:"acceptanceTimeout" yaml:"acceptanceTimeout"`
	// Information to add to all events related to the matchmaking configuration.
	// Experimental.
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A human-readable description of the matchmaking configuration.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	// See: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html
	//
	// Experimental.
	NotificationTarget awssns.ITopic `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The maximum duration, that a matchmaking ticket can remain in process before timing out.
	//
	// Requests that fail due to timing out can be resubmitted as needed.
	// Experimental.
	RequestTimeout awscdk.Duration `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// A flag that determines whether a match that was created with this configuration must be accepted by the matched players.
	//
	// With this option enabled, matchmaking tickets use the status `REQUIRES_ACCEPTANCE` to indicate when a completed potential match is waiting for player acceptance.
	// Experimental.
	RequireAcceptance *bool `field:"optional" json:"requireAcceptance" yaml:"requireAcceptance"`
	// Queues are used to start new GameLift-hosted game sessions for matches that are created with this matchmaking configuration.
	//
	// Queues can be located in any Region.
	// Experimental.
	GameSessionQueues *[]IGameSessionQueue `field:"required" json:"gameSessionQueues" yaml:"gameSessionQueues"`
	// The number of player slots in a match to keep open for future players.
	//
	// For example, if the configuration's rule set specifies a match for a single 12-person team, and the additional player count is set to 2, only 10 players are selected for the match.
	// Experimental.
	AdditionalPlayerCount *float64 `field:"optional" json:"additionalPlayerCount" yaml:"additionalPlayerCount"`
	// A set of custom properties for a game session, formatted as key-value pairs.
	//
	// These properties are passed to a game server process with a request to start a new game session.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession
	//
	// Experimental.
	GameProperties *[]*GameProperty `field:"optional" json:"gameProperties" yaml:"gameProperties"`
	// A set of custom game session properties, formatted as a single string value.
	//
	// This data is passed to a game server process with a request to start a new game session.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession
	//
	// Experimental.
	GameSessionData *string `field:"optional" json:"gameSessionData" yaml:"gameSessionData"`
	// The method used to backfill game sessions that are created with this matchmaking configuration.
	//
	// - Choose manual when your game manages backfill requests manually or does not use the match backfill feature.
	// - Otherwise backfill is settled to automatic to have GameLift create a `StartMatchBackfill` request whenever a game session has one or more open slots.
	// See: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html
	//
	// Experimental.
	ManualBackfillMode *bool `field:"optional" json:"manualBackfillMode" yaml:"manualBackfillMode"`
}

