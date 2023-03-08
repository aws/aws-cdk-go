package awsgamelift


// A list of fleet locations where a game session queue can place new game sessions.
//
// You can use a filter to temporarily turn off placements for specific locations. For queues that have multi-location fleets, you can use a filter configuration allow placement with some, but not all of these locations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterConfigurationProperty := &filterConfigurationProperty{
//   	allowedLocations: []*string{
//   		jsii.String("allowedLocations"),
//   	},
//   }
//
type CfnGameSessionQueue_FilterConfigurationProperty struct {
	// A list of locations to allow game session placement in, in the form of AWS Region codes such as `us-west-2` .
	AllowedLocations *[]*string `field:"optional" json:"allowedLocations" yaml:"allowedLocations"`
}

