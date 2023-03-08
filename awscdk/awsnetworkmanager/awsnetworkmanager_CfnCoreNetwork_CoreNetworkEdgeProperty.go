package awsnetworkmanager


// Describes a core network edge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreNetworkEdgeProperty := &coreNetworkEdgeProperty{
//   	asn: jsii.Number(123),
//   	edgeLocation: jsii.String("edgeLocation"),
//   	insideCidrBlocks: []*string{
//   		jsii.String("insideCidrBlocks"),
//   	},
//   }
//
type CfnCoreNetwork_CoreNetworkEdgeProperty struct {
	// The ASN of a core network edge.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// The Region where a core network edge is located.
	EdgeLocation *string `field:"optional" json:"edgeLocation" yaml:"edgeLocation"`
	// The inside IP addresses used for core network edges.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
}

