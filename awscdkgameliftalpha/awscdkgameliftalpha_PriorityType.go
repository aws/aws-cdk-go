// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// Priority to condider when placing new game sessions.
//
// Example:
//   var fleet buildFleet
//   var topic topic
//
//
//   gamelift.NewGameSessionQueue(this, jsii.String("MyGameSessionQueue"), &gameSessionQueueProps{
//   	gameSessionQueueName: jsii.String("test-gameSessionQueue"),
//   	customEventData: jsii.String("test-event-data"),
//   	allowedLocations: []*string{
//   		jsii.String("eu-west-1"),
//   		jsii.String("eu-west-2"),
//   	},
//   	destinations: []iGameSessionQueueDestination{
//   		fleet,
//   	},
//   	notificationTarget: topic,
//   	playerLatencyPolicies: []playerLatencyPolicy{
//   		&playerLatencyPolicy{
//   			maximumIndividualPlayerLatency: awscdk.Duration.millis(jsii.Number(100)),
//   			policyDuration: awscdk.Duration.seconds(jsii.Number(300)),
//   		},
//   	},
//   	priorityConfiguration: &priorityConfiguration{
//   		locationOrder: []*string{
//   			jsii.String("eu-west-1"),
//   			jsii.String("eu-west-2"),
//   		},
//   		priorityOrder: []priorityType{
//   			gamelift.*priorityType_LATENCY,
//   			gamelift.*priorityType_COST,
//   			gamelift.*priorityType_DESTINATION,
//   			gamelift.*priorityType_LOCATION,
//   		},
//   	},
//   	timeout: awscdk.Duration.seconds(jsii.Number(300)),
//   })
//
// Experimental.
type PriorityType string

const (
	// FleetIQ prioritizes locations where the average player latency (provided in each game session request) is lowest.
	// Experimental.
	PriorityType_LATENCY PriorityType = "LATENCY"
	// FleetIQ prioritizes destinations with the lowest current hosting costs.
	//
	// Cost is evaluated based on the location, instance type, and fleet type (Spot or On-Demand) for each destination in the queue.
	// Experimental.
	PriorityType_COST PriorityType = "COST"
	// FleetIQ prioritizes based on the order that destinations are listed in the queue configuration.
	// Experimental.
	PriorityType_DESTINATION PriorityType = "DESTINATION"
	// FleetIQ prioritizes based on the provided order of locations, as defined in `LocationOrder`.
	// Experimental.
	PriorityType_LOCATION PriorityType = "LOCATION"
)

