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
//   coreNetworkSegmentProperty := &CoreNetworkSegmentProperty{
//   	EdgeLocations: []*string{
//   		jsii.String("edgeLocations"),
//   	},
//   	Name: jsii.String("name"),
//   	SharedSegments: []*string{
//   		jsii.String("sharedSegments"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html
//
type CfnCoreNetwork_CoreNetworkSegmentProperty struct {
	// The Regions where the edges are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html#cfn-networkmanager-corenetwork-corenetworksegment-edgelocations
	//
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
	// The name of a core network segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html#cfn-networkmanager-corenetwork-corenetworksegment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The shared segments of a core network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html#cfn-networkmanager-corenetwork-corenetworksegment-sharedsegments
	//
	SharedSegments *[]*string `field:"optional" json:"sharedSegments" yaml:"sharedSegments"`
}

