package awsgamelift


// Custom prioritization settings for use by a game session queue when placing new game sessions with available game servers.
//
// When defined, this configuration replaces the default FleetIQ prioritization process, which is as follows:
//
// - If player latency data is included in a game session request, destinations and locations are prioritized first based on lowest average latency (1), then on lowest hosting cost (2), then on destination list order (3), and finally on location (alphabetical) (4). This approach ensures that the queue's top priority is to place game sessions where average player latency is lowest, and--if latency is the same--where the hosting cost is less, etc.
// - If player latency data is not included, destinations and locations are prioritized first on destination list order (1), and then on location (alphabetical) (2). This approach ensures that the queue's top priority is to place game sessions on the first destination fleet listed. If that fleet has multiple locations, the game session is placed on the first location (when listed alphabetically).
//
// Changing the priority order will affect how game sessions are placed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   priorityConfigurationProperty := &PriorityConfigurationProperty{
//   	LocationOrder: []*string{
//   		jsii.String("locationOrder"),
//   	},
//   	PriorityOrder: []*string{
//   		jsii.String("priorityOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html
//
type CfnGameSessionQueue_PriorityConfigurationProperty struct {
	// The prioritization order to use for fleet locations, when the `PriorityOrder` property includes `LOCATION` .
	//
	// Locations can include AWS Region codes (such as `us-west-2` ), local zones, and custom locations (for Anywhere fleets). Each location must be listed only once. For details, see [Amazon GameLift Servers service locations.](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html#cfn-gamelift-gamesessionqueue-priorityconfiguration-locationorder
	//
	LocationOrder *[]*string `field:"optional" json:"locationOrder" yaml:"locationOrder"`
	// A custom sequence to use when prioritizing where to place new game sessions. Each priority type is listed once.
	//
	// - `LATENCY` -- Amazon GameLift Servers prioritizes locations where the average player latency is lowest. Player latency data is provided in each game session placement request.
	// - `COST` -- Amazon GameLift Servers prioritizes queue destinations with the lowest current hosting costs. Cost is evaluated based on the destination's location, instance type, and fleet type (Spot or On-Demand).
	// - `DESTINATION` -- Amazon GameLift Servers prioritizes based on the list order of destinations in the queue configuration.
	// - `LOCATION` -- Amazon GameLift Servers prioritizes based on the provided order of locations, as defined in `LocationOrder` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html#cfn-gamelift-gamesessionqueue-priorityconfiguration-priorityorder
	//
	PriorityOrder *[]*string `field:"optional" json:"priorityOrder" yaml:"priorityOrder"`
}

