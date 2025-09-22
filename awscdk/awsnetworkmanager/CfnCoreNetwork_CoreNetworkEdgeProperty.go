package awsnetworkmanager


// Describes a core network edge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreNetworkEdgeProperty := &CoreNetworkEdgeProperty{
//   	Asn: jsii.Number(123),
//   	EdgeLocation: jsii.String("edgeLocation"),
//   	InsideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html
//
type CfnCoreNetwork_CoreNetworkEdgeProperty struct {
	// The ASN of a core network edge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-asn
	//
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// The Region where a core network edge is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-edgelocation
	//
	EdgeLocation *string `field:"optional" json:"edgeLocation" yaml:"edgeLocation"`
	// The inside IP addresses used for core network edges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworkedge.html#cfn-networkmanager-corenetwork-corenetworkedge-insidecidrblocks
	//
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
}

