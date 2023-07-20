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
//   analysisPacketHeaderProperty := &AnalysisPacketHeaderProperty{
//   	DestinationAddresses: []*string{
//   		jsii.String("destinationAddresses"),
//   	},
//   	DestinationPortRanges: []interface{}{
//   		&PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   	},
//   	Protocol: jsii.String("protocol"),
//   	SourceAddresses: []*string{
//   		jsii.String("sourceAddresses"),
//   	},
//   	SourcePortRanges: []interface{}{
//   		&PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysispacketheader.html
//
type CfnNetworkInsightsAnalysis_AnalysisPacketHeaderProperty struct {
	// The destination addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysispacketheader.html#cfn-ec2-networkinsightsanalysis-analysispacketheader-destinationaddresses
	//
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// The destination port ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysispacketheader.html#cfn-ec2-networkinsightsanalysis-analysispacketheader-destinationportranges
	//
	DestinationPortRanges interface{} `field:"optional" json:"destinationPortRanges" yaml:"destinationPortRanges"`
	// The protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysispacketheader.html#cfn-ec2-networkinsightsanalysis-analysispacketheader-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The source addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysispacketheader.html#cfn-ec2-networkinsightsanalysis-analysispacketheader-sourceaddresses
	//
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// The source port ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysispacketheader.html#cfn-ec2-networkinsightsanalysis-analysispacketheader-sourceportranges
	//
	SourcePortRanges interface{} `field:"optional" json:"sourcePortRanges" yaml:"sourcePortRanges"`
}

