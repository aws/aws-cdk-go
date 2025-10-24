package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMatchmakingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMatchmakingConfigurationProps := &CfnMatchmakingConfigurationProps{
//   	AcceptanceRequired: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	RequestTimeoutSeconds: jsii.Number(123),
//   	RuleSetName: jsii.String("ruleSetName"),
//
//   	// the properties below are optional
//   	AcceptanceTimeoutSeconds: jsii.Number(123),
//   	AdditionalPlayerCount: jsii.Number(123),
//   	BackfillMode: jsii.String("backfillMode"),
//   	CreationTime: jsii.String("creationTime"),
//   	CustomEventData: jsii.String("customEventData"),
//   	Description: jsii.String("description"),
//   	FlexMatchMode: jsii.String("flexMatchMode"),
//   	GameProperties: []interface{}{
//   		&GamePropertyProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	GameSessionData: jsii.String("gameSessionData"),
//   	GameSessionQueueArns: []*string{
//   		jsii.String("gameSessionQueueArns"),
//   	},
//   	NotificationTarget: jsii.String("notificationTarget"),
//   	RuleSetArn: jsii.String("ruleSetArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html
//
type CfnMatchmakingConfigurationProps struct {
	// A flag that determines whether a match that was created with this configuration must be accepted by the matched players.
	//
	// To require acceptance, set to `TRUE` . With this option enabled, matchmaking tickets use the status `REQUIRES_ACCEPTANCE` to indicate when a completed potential match is waiting for player acceptance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-acceptancerequired
	//
	AcceptanceRequired interface{} `field:"required" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// A unique identifier for the matchmaking configuration.
	//
	// This name is used to identify the configuration associated with a matchmaking request or ticket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
	//
	// Requests that fail due to timing out can be resubmitted as needed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-requesttimeoutseconds
	//
	RequestTimeoutSeconds *float64 `field:"required" json:"requestTimeoutSeconds" yaml:"requestTimeoutSeconds"`
	// A unique identifier for the matchmaking rule set to use with this configuration.
	//
	// You can use either the rule set name or ARN value. A matchmaking configuration can only use rule sets that are defined in the same Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-rulesetname
	//
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
	// The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-acceptancetimeoutseconds
	//
	AcceptanceTimeoutSeconds *float64 `field:"optional" json:"acceptanceTimeoutSeconds" yaml:"acceptanceTimeoutSeconds"`
	// The number of player slots in a match to keep open for future players.
	//
	// For example, if the configuration's rule set specifies a match for a single 12-person team, and the additional player count is set to 2, only 10 players are selected for the match. This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-additionalplayercount
	//
	AdditionalPlayerCount *float64 `field:"optional" json:"additionalPlayerCount" yaml:"additionalPlayerCount"`
	// The method used to backfill game sessions that are created with this matchmaking configuration.
	//
	// Specify `MANUAL` when your game manages backfill requests manually or does not use the match backfill feature. Specify `AUTOMATIC` to have GameLift create a `StartMatchBackfill` request whenever a game session has one or more open slots. Learn more about manual and automatic backfill in [Backfill Existing Games with FlexMatch](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html) . Automatic backfill is not available when `FlexMatchMode` is set to `STANDALONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-backfillmode
	//
	BackfillMode *string `field:"optional" json:"backfillMode" yaml:"backfillMode"`
	// A time stamp indicating when this data object was created.
	//
	// Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-creationtime
	//
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Information to add to all events related to the matchmaking configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-customeventdata
	//
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A description for the matchmaking configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether this matchmaking configuration is being used with Amazon GameLift Servers hosting or as a standalone matchmaking solution.
	//
	// - *STANDALONE* - FlexMatch forms matches and returns match information, including players and team assignments, in a [MatchmakingSucceeded](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html#match-events-matchmakingsucceeded) event.
	// - *WITH_QUEUE* - FlexMatch forms matches and uses the specified Amazon GameLift Servers queue to start a game session for the match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-flexmatchmode
	//
	FlexMatchMode *string `field:"optional" json:"flexMatchMode" yaml:"flexMatchMode"`
	// A set of custom properties for a game session, formatted as key-value pairs.
	//
	// These properties are passed to a game server process with a request to start a new game session. See [Start a Game Session](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession) . This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-gameproperties
	//
	GameProperties interface{} `field:"optional" json:"gameProperties" yaml:"gameProperties"`
	// A set of custom game session properties, formatted as a single string value.
	//
	// This data is passed to a game server process with a request to start a new game session. See [Start a Game Session](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession) . This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-gamesessiondata
	//
	GameSessionData *string `field:"optional" json:"gameSessionData" yaml:"gameSessionData"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to a Amazon GameLift Servers game session queue resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::gamesessionqueue/<queue name>` . Queues can be located in any Region. Queues are used to start new Amazon GameLift Servers-hosted game sessions for matches that are created with this matchmaking configuration. If `FlexMatchMode` is set to `STANDALONE` , do not set this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-gamesessionqueuearns
	//
	GameSessionQueueArns *[]*string `field:"optional" json:"gameSessionQueueArns" yaml:"gameSessionQueueArns"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	//
	// See [Setting up notifications for matchmaking](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html) for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-notificationtarget
	//
	NotificationTarget *string `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) associated with the GameLift matchmaking rule set resource that this configuration uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-rulesetarn
	//
	RuleSetArn *string `field:"optional" json:"ruleSetArn" yaml:"ruleSetArn"`
	// A list of labels to assign to the new matchmaking configuration resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

