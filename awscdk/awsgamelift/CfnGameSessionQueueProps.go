package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGameSessionQueue`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGameSessionQueueProps := &CfnGameSessionQueueProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CustomEventData: jsii.String("customEventData"),
//   	Destinations: []interface{}{
//   		&DestinationProperty{
//   			DestinationArn: jsii.String("destinationArn"),
//   		},
//   	},
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		AllowedLocations: []*string{
//   			jsii.String("allowedLocations"),
//   		},
//   	},
//   	NotificationTarget: jsii.String("notificationTarget"),
//   	PlayerLatencyPolicies: []interface{}{
//   		&PlayerLatencyPolicyProperty{
//   			MaximumIndividualPlayerLatencyMilliseconds: jsii.Number(123),
//   			PolicyDurationSeconds: jsii.Number(123),
//   		},
//   	},
//   	PriorityConfiguration: &PriorityConfigurationProperty{
//   		LocationOrder: []*string{
//   			jsii.String("locationOrder"),
//   		},
//   		PriorityOrder: []*string{
//   			jsii.String("priorityOrder"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html
//
type CfnGameSessionQueueProps struct {
	// A descriptive label that is associated with game session queue.
	//
	// Queue names must be unique within each Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Information to be added to all events that are related to this game session queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-customeventdata
	//
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.
	//
	// Destinations are identified by either a fleet ARN or a fleet alias ARN, and are listed in order of placement preference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// A list of locations where a queue is allowed to place new game sessions.
	//
	// Locations are specified in the form of AWS Region codes, such as `us-west-2` . If this parameter is not set, game sessions can be placed in any queue location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// An SNS topic ARN that is set up to receive game session placement notifications.
	//
	// See [Setting up notifications for game session placement](https://docs.aws.amazon.com/gamelift/latest/developerguide/queue-notification.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-notificationtarget
	//
	NotificationTarget *string `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// A set of policies that enforce a sliding cap on player latency when processing game sessions placement requests.
	//
	// Use multiple policies to gradually relax the cap over time if Amazon GameLift can't make a placement. Policies are evaluated in order starting with the lowest maximum latency value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-playerlatencypolicies
	//
	PlayerLatencyPolicies interface{} `field:"optional" json:"playerLatencyPolicies" yaml:"playerLatencyPolicies"`
	// Custom settings to use when prioritizing destinations and locations for game session placements.
	//
	// This configuration replaces the FleetIQ default prioritization process. Priority types that are not explicitly named will be automatically applied at the end of the prioritization process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-priorityconfiguration
	//
	PriorityConfiguration interface{} `field:"optional" json:"priorityConfiguration" yaml:"priorityConfiguration"`
	// A list of labels to assign to the new game session queue resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The maximum time, in seconds, that a new game session placement request remains in the queue.
	//
	// When a request exceeds this time, the game session placement changes to a `TIMED_OUT` status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-timeoutinseconds
	//
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

