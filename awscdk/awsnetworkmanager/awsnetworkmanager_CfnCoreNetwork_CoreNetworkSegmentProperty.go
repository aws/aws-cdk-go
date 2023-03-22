package awsnetworkmanager


// Describes a core network segment, which are dedicated routes.
//
// Only attachments within this segment can communicate with each other.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreNetworkSegmentProperty := &coreNetworkSegmentProperty{
//   	edgeLocations: []*string{
//   		jsii.String("edgeLocations"),
//   	},
//   	name: jsii.String("name"),
//   	sharedSegments: []*string{
//   		jsii.String("sharedSegments"),
//   	},
//   }
//
type CfnCoreNetwork_CoreNetworkSegmentProperty struct {
	// The Regions where the edges are located.
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
	// The name of a core network segment.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The shared segments of a core network.
	SharedSegments *[]*string `field:"optional" json:"sharedSegments" yaml:"sharedSegments"`
}

