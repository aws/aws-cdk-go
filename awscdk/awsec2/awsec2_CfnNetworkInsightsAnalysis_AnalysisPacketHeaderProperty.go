package awsec2


// Describes a header.
//
// Reflects any changes made by a component as traffic passes through. The fields of an inbound header are null except for the first component of a path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisPacketHeaderProperty := &analysisPacketHeaderProperty{
//   	destinationAddresses: []*string{
//   		jsii.String("destinationAddresses"),
//   	},
//   	destinationPortRanges: []interface{}{
//   		&portRangeProperty{
//   			from: jsii.Number(123),
//   			to: jsii.Number(123),
//   		},
//   	},
//   	protocol: jsii.String("protocol"),
//   	sourceAddresses: []*string{
//   		jsii.String("sourceAddresses"),
//   	},
//   	sourcePortRanges: []interface{}{
//   		&portRangeProperty{
//   			from: jsii.Number(123),
//   			to: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisPacketHeaderProperty struct {
	// The destination addresses.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// The destination port ranges.
	DestinationPortRanges interface{} `field:"optional" json:"destinationPortRanges" yaml:"destinationPortRanges"`
	// The protocol.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The source addresses.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// The source port ranges.
	SourcePortRanges interface{} `field:"optional" json:"sourcePortRanges" yaml:"sourcePortRanges"`
}

