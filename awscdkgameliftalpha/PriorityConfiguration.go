package awscdkgameliftalpha


// Custom prioritization settings for use by a game session queue when placing new game sessions with available game servers.
//
// When defined, this configuration replaces the default FleetIQ prioritization process, which is as follows:
//
// - If player latency data is included in a game session request, destinations and locations are prioritized first based on lowest average latency (1), then on lowest hosting cost (2), then on destination list order (3), and finally on location (alphabetical) (4).
// This approach ensures that the queue's top priority is to place game sessions where average player latency is lowest, and--if latency is the same--where the hosting cost is less, etc.
//
// - If player latency data is not included, destinations and locations are prioritized first on destination list order (1), and then on location (alphabetical) (2).
// This approach ensures that the queue's top priority is to place game sessions on the first destination fleet listed. If that fleet has multiple locations, the game session is placed on the first location (when listed alphabetically).
//
// Changing the priority order will affect how game sessions are placed.
//
// Example:
//   var fleet buildFleet
//   var topic topic
//
//
//   gamelift.NewGameSessionQueue(this, jsii.String("MyGameSessionQueue"), &GameSessionQueueProps{
//   	GameSessionQueueName: jsii.String("test-gameSessionQueue"),
//   	CustomEventData: jsii.String("test-event-data"),
//   	AllowedLocations: []*string{
//   		jsii.String("eu-west-1"),
//   		jsii.String("eu-west-2"),
//   	},
//   	Destinations: []iGameSessionQueueDestination{
//   		fleet,
//   	},
//   	NotificationTarget: topic,
//   	PlayerLatencyPolicies: []playerLatencyPolicy{
//   		&playerLatencyPolicy{
//   			MaximumIndividualPlayerLatency: awscdk.Duration_Millis(jsii.Number(100)),
//   			PolicyDuration: awscdk.Duration_Seconds(jsii.Number(300)),
//   		},
//   	},
//   	PriorityConfiguration: &PriorityConfiguration{
//   		LocationOrder: []*string{
//   			jsii.String("eu-west-1"),
//   			jsii.String("eu-west-2"),
//   		},
//   		PriorityOrder: []priorityType{
//   			gamelift.*priorityType_LATENCY,
//   			gamelift.*priorityType_COST,
//   			gamelift.*priorityType_DESTINATION,
//   			gamelift.*priorityType_LOCATION,
//   		},
//   	},
//   	Timeout: awscdk.Duration_*Seconds(jsii.Number(300)),
//   })
//
// Experimental.
type PriorityConfiguration struct {
	// The prioritization order to use for fleet locations, when the PriorityOrder property includes LOCATION.
	//
	// Locations are identified by AWS Region codes such as `us-west-2.
	//
	// Each location can only be listed once.
	// Experimental.
	LocationOrder *[]*string `field:"required" json:"locationOrder" yaml:"locationOrder"`
	// The recommended sequence to use when prioritizing where to place new game sessions.
	//
	// Each type can only be listed once.
	// Experimental.
	PriorityOrder *[]PriorityType `field:"required" json:"priorityOrder" yaml:"priorityOrder"`
}

