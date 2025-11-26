package previewawsgameliftmixins


// A list of fleet locations where a game session queue can place new game sessions.
//
// You can use a filter to temporarily turn off placements for specific locations. For queues that have multi-location fleets, you can use a filter configuration allow placement with some, but not all of these locations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterConfigurationProperty := &FilterConfigurationProperty{
//   	AllowedLocations: []*string{
//   		jsii.String("allowedLocations"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-filterconfiguration.html
//
type CfnGameSessionQueuePropsMixin_FilterConfigurationProperty struct {
	// A list of locations to allow game session placement in, in the form of AWS Region codes such as `us-west-2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-filterconfiguration.html#cfn-gamelift-gamesessionqueue-filterconfiguration-allowedlocations
	//
	AllowedLocations *[]*string `field:"optional" json:"allowedLocations" yaml:"allowedLocations"`
}

