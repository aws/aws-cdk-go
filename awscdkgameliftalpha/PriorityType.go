package awscdkgameliftalpha


// Priority to condider when placing new game sessions.
//
// Example:
//   var fleet BuildFleet
//   var topic Topic
//
//
//   gamelift.NewGameSessionQueue(this, jsii.String("MyGameSessionQueue"), &GameSessionQueueProps{
//   	GameSessionQueueName: jsii.String("test-gameSessionQueue"),
//   	CustomEventData: jsii.String("test-event-data"),
//   	AllowedLocations: []*string{
//   		jsii.String("eu-west-1"),
//   		jsii.String("eu-west-2"),
//   	},
//   	Destinations: []IGameSessionQueueDestination{
//   		fleet,
//   	},
//   	NotificationTarget: topic,
//   	PlayerLatencyPolicies: []PlayerLatencyPolicy{
//   		&PlayerLatencyPolicy{
//   			MaximumIndividualPlayerLatency: awscdk.Duration_Millis(jsii.Number(100)),
//   			PolicyDuration: awscdk.Duration_Seconds(jsii.Number(300)),
//   		},
//   	},
//   	PriorityConfiguration: &PriorityConfiguration{
//   		LocationOrder: []*string{
//   			jsii.String("eu-west-1"),
//   			jsii.String("eu-west-2"),
//   		},
//   		PriorityOrder: []PriorityType{
//   			gamelift.PriorityType_LATENCY,
//   			gamelift.PriorityType_COST,
//   			gamelift.PriorityType_DESTINATION,
//   			gamelift.PriorityType_LOCATION,
//   		},
//   	},
//   	Timeout: awscdk.Duration_*Seconds(jsii.Number(300)),
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

