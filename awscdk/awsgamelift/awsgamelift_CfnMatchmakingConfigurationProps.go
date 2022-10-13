package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMatchmakingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMatchmakingConfigurationProps := &cfnMatchmakingConfigurationProps{
//   	acceptanceRequired: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	requestTimeoutSeconds: jsii.Number(123),
//   	ruleSetName: jsii.String("ruleSetName"),
//
//   	// the properties below are optional
//   	acceptanceTimeoutSeconds: jsii.Number(123),
//   	additionalPlayerCount: jsii.Number(123),
//   	backfillMode: jsii.String("backfillMode"),
//   	customEventData: jsii.String("customEventData"),
//   	description: jsii.String("description"),
//   	flexMatchMode: jsii.String("flexMatchMode"),
//   	gameProperties: []interface{}{
//   		&gamePropertyProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	gameSessionData: jsii.String("gameSessionData"),
//   	gameSessionQueueArns: []*string{
//   		jsii.String("gameSessionQueueArns"),
//   	},
//   	notificationTarget: jsii.String("notificationTarget"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMatchmakingConfigurationProps struct {
	// A flag that determines whether a match that was created with this configuration must be accepted by the matched players.
	//
	// To require acceptance, set to `TRUE` . With this option enabled, matchmaking tickets use the status `REQUIRES_ACCEPTANCE` to indicate when a completed potential match is waiting for player acceptance.
	AcceptanceRequired interface{} `field:"required" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// A unique identifier for the matchmaking configuration.
	//
	// This name is used to identify the configuration associated with a matchmaking request or ticket.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
	//
	// Requests that fail due to timing out can be resubmitted as needed.
	RequestTimeoutSeconds *float64 `field:"required" json:"requestTimeoutSeconds" yaml:"requestTimeoutSeconds"`
	// A unique identifier for the matchmaking rule set to use with this configuration.
	//
	// You can use either the rule set name or ARN value. A matchmaking configuration can only use rule sets that are defined in the same Region.
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
	// The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
	AcceptanceTimeoutSeconds *float64 `field:"optional" json:"acceptanceTimeoutSeconds" yaml:"acceptanceTimeoutSeconds"`
	// The number of player slots in a match to keep open for future players.
	//
	// For example, if the configuration's rule set specifies a match for a single 12-person team, and the additional player count is set to 2, only 10 players are selected for the match. This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	AdditionalPlayerCount *float64 `field:"optional" json:"additionalPlayerCount" yaml:"additionalPlayerCount"`
	// The method used to backfill game sessions that are created with this matchmaking configuration.
	//
	// Specify `MANUAL` when your game manages backfill requests manually or does not use the match backfill feature. Specify `AUTOMATIC` to have GameLift create a `StartMatchBackfill` request whenever a game session has one or more open slots. Learn more about manual and automatic backfill in [Backfill Existing Games with FlexMatch](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html) . Automatic backfill is not available when `FlexMatchMode` is set to `STANDALONE` .
	BackfillMode *string `field:"optional" json:"backfillMode" yaml:"backfillMode"`
	// Information to add to all events related to the matchmaking configuration.
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A descriptive label that is associated with matchmaking configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether this matchmaking configuration is being used with GameLift hosting or as a standalone matchmaking solution.
	//
	// - *STANDALONE* - FlexMatch forms matches and returns match information, including players and team assignments, in a [MatchmakingSucceeded](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html#match-events-matchmakingsucceeded) event.
	// - *WITH_QUEUE* - FlexMatch forms matches and uses the specified GameLift queue to start a game session for the match.
	FlexMatchMode *string `field:"optional" json:"flexMatchMode" yaml:"flexMatchMode"`
	// A set of custom properties for a game session, formatted as key-value pairs.
	//
	// These properties are passed to a game server process with a request to start a new game session. See [Start a Game Session](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession) . This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	GameProperties interface{} `field:"optional" json:"gameProperties" yaml:"gameProperties"`
	// A set of custom game session properties, formatted as a single string value.
	//
	// This data is passed to a game server process with a request to start a new game session. See [Start a Game Session](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession) . This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	GameSessionData *string `field:"optional" json:"gameSessionData" yaml:"gameSessionData"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to a GameLift game session queue resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::gamesessionqueue/<queue name>` . Queues can be located in any Region. Queues are used to start new GameLift-hosted game sessions for matches that are created with this matchmaking configuration. If `FlexMatchMode` is set to `STANDALONE` , do not set this parameter.
	GameSessionQueueArns *[]*string `field:"optional" json:"gameSessionQueueArns" yaml:"gameSessionQueueArns"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	//
	// See [Setting up notifications for matchmaking](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html) for more information.
	NotificationTarget *string `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// A list of labels to assign to the new matchmaking configuration resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

