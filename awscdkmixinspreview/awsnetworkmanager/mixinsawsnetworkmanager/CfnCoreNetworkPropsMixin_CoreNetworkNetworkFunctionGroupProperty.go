package mixinsawsnetworkmanager


// Describes a network function group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   coreNetworkNetworkFunctionGroupProperty := &CoreNetworkNetworkFunctionGroupProperty{
//   	EdgeLocations: []*string{
//   		jsii.String("edgeLocations"),
//   	},
//   	Name: jsii.String("name"),
//   	Segments: &SegmentsProperty{
//   		SendTo: []*string{
//   			jsii.String("sendTo"),
//   		},
//   		SendVia: []*string{
//   			jsii.String("sendVia"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworknetworkfunctiongroup.html
//
type CfnCoreNetworkPropsMixin_CoreNetworkNetworkFunctionGroupProperty struct {
	// The core network edge locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworknetworkfunctiongroup.html#cfn-networkmanager-corenetwork-corenetworknetworkfunctiongroup-edgelocations
	//
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
	// The name of the network function group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworknetworkfunctiongroup.html#cfn-networkmanager-corenetwork-corenetworknetworkfunctiongroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The segments associated with the network function group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworknetworkfunctiongroup.html#cfn-networkmanager-corenetwork-corenetworknetworkfunctiongroup-segments
	//
	Segments interface{} `field:"optional" json:"segments" yaml:"segments"`
}

