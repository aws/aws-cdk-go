package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a new Gamelift matchmaking configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var matchmakingRuleSet MatchmakingRuleSet
//   var topic Topic
//
//   matchmakingConfigurationProps := &MatchmakingConfigurationProps{
//   	MatchmakingConfigurationName: jsii.String("matchmakingConfigurationName"),
//   	RuleSet: matchmakingRuleSet,
//
//   	// the properties below are optional
//   	AcceptanceTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	CustomEventData: jsii.String("customEventData"),
//   	Description: jsii.String("description"),
//   	NotificationTarget: topic,
//   	RequestTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	RequireAcceptance: jsii.Boolean(false),
//   }
//
// Experimental.
type MatchmakingConfigurationProps struct {
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
	// Default: 300 seconds.
	//
	// Experimental.
	AcceptanceTimeout awscdk.Duration `field:"optional" json:"acceptanceTimeout" yaml:"acceptanceTimeout"`
	// Information to add to all events related to the matchmaking configuration.
	// Default: no custom data added to events.
	//
	// Experimental.
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A human-readable description of the matchmaking configuration.
	// Default: no description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	// See: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html
	//
	// Default: no notification target.
	//
	// Experimental.
	NotificationTarget awssns.ITopic `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The maximum duration, that a matchmaking ticket can remain in process before timing out.
	//
	// Requests that fail due to timing out can be resubmitted as needed.
	// Default: 300 seconds.
	//
	// Experimental.
	RequestTimeout awscdk.Duration `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// A flag that determines whether a match that was created with this configuration must be accepted by the matched players.
	//
	// With this option enabled, matchmaking tickets use the status `REQUIRES_ACCEPTANCE` to indicate when a completed potential match is waiting for player acceptance.
	// Default: Acceptance is not required.
	//
	// Experimental.
	RequireAcceptance *bool `field:"optional" json:"requireAcceptance" yaml:"requireAcceptance"`
}

