package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a new Fleet gameSessionQueue.
//
// Example:
//   var fleet buildFleet
//   var alias alias
//
//
//   queue := gamelift.NewGameSessionQueue(this, jsii.String("GameSessionQueue"), &GameSessionQueueProps{
//   	GameSessionQueueName: jsii.String("my-queue-name"),
//   	Destinations: []iGameSessionQueueDestination{
//   		fleet,
//   	},
//   })
//   queue.AddDestination(alias)
//
// Experimental.
type GameSessionQueueProps struct {
	// A list of fleets and/or fleet alias that can be used to fulfill game session placement requests in the queue.
	//
	// Destinations are listed in order of placement preference.
	// Experimental.
	Destinations *[]IGameSessionQueueDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Name of this gameSessionQueue.
	// Experimental.
	GameSessionQueueName *string `field:"required" json:"gameSessionQueueName" yaml:"gameSessionQueueName"`
	// A list of locations where a queue is allowed to place new game sessions.
	//
	// Locations are specified in the form of AWS Region codes, such as `us-west-2`.
	//
	// For queues that have multi-location fleets, you can use a filter configuration allow placement with some, but not all of these locations.
	// Default: game sessions can be placed in any queue location.
	//
	// Experimental.
	AllowedLocations *[]*string `field:"optional" json:"allowedLocations" yaml:"allowedLocations"`
	// Information to be added to all events that are related to this game session queue.
	// Default: no customer event data.
	//
	// Experimental.
	CustomEventData *string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// An SNS topic is set up to receive game session placement notifications.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/queue-notification.html
	//
	// Default: no notification.
	//
	// Experimental.
	NotificationTarget awssns.ITopic `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// A set of policies that act as a sliding cap on player latency.
	//
	// FleetIQ works to deliver low latency for most players in a game session.
	// These policies ensure that no individual player can be placed into a game with unreasonably high latency.
	// Use multiple policies to gradually relax latency requirements a step at a time.
	// Multiple policies are applied based on their maximum allowed latency, starting with the lowest value.
	// Default: no player latency policy.
	//
	// Experimental.
	PlayerLatencyPolicies *[]*PlayerLatencyPolicy `field:"optional" json:"playerLatencyPolicies" yaml:"playerLatencyPolicies"`
	// Custom settings to use when prioritizing destinations and locations for game session placements.
	//
	// This configuration replaces the FleetIQ default prioritization process.
	//
	// Priority types that are not explicitly named will be automatically applied at the end of the prioritization process.
	// Default: no priority configuration.
	//
	// Experimental.
	PriorityConfiguration *PriorityConfiguration `field:"optional" json:"priorityConfiguration" yaml:"priorityConfiguration"`
	// The maximum time, that a new game session placement request remains in the queue.
	//
	// When a request exceeds this time, the game session placement changes to a `TIMED_OUT` status.
	// Default: 50 seconds.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

