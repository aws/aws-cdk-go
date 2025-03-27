package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// A full specification of a matchmaking configuration that can be used to import it fluently into the CDK application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   matchmakingConfigurationAttributes := &MatchmakingConfigurationAttributes{
//   	MatchmakingConfigurationArn: jsii.String("matchmakingConfigurationArn"),
//   	MatchmakingConfigurationName: jsii.String("matchmakingConfigurationName"),
//   	NotificationTarget: topic,
//   }
//
// Experimental.
type MatchmakingConfigurationAttributes struct {
	// The ARN of the Matchmaking configuration.
	//
	// At least one of `matchmakingConfigurationArn` and `matchmakingConfigurationName` must be provided.
	// Default: derived from `matchmakingConfigurationName`.
	//
	// Experimental.
	MatchmakingConfigurationArn *string `field:"optional" json:"matchmakingConfigurationArn" yaml:"matchmakingConfigurationArn"`
	// The identifier of the Matchmaking configuration.
	//
	// At least one of `matchmakingConfigurationName` and `matchmakingConfigurationArn`  must be provided.
	// Default: derived from `matchmakingConfigurationArn`.
	//
	// Experimental.
	MatchmakingConfigurationName *string `field:"optional" json:"matchmakingConfigurationName" yaml:"matchmakingConfigurationName"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	// See: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html
	//
	// Default: no notification target binded to imported ressource.
	//
	// Experimental.
	NotificationTarget awssns.ITopic `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
}

