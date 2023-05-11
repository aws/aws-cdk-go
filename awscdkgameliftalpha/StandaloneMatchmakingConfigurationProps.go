// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a new standalone matchmaking configuration.
//
// Example:
//   var ruleSet matchmakingRuleSet
//
//
//   gamelift.NewStandaloneMatchmakingConfiguration(this, jsii.String("StandaloneMatchmaking"), &StandaloneMatchmakingConfigurationProps{
//   	MatchmakingConfigurationName: jsii.String("test-standalone-config-name"),
//   	RuleSet: ruleSet,
//   })
//
// Experimental.
type StandaloneMatchmakingConfigurationProps struct {
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
}

