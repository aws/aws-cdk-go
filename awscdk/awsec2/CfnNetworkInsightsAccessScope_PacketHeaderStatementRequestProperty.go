package awsec2


// Describes a packet header statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packetHeaderStatementRequestProperty := &PacketHeaderStatementRequestProperty{
//   	DestinationAddresses: []*string{
//   		jsii.String("destinationAddresses"),
//   	},
//   	DestinationPorts: []*string{
//   		jsii.String("destinationPorts"),
//   	},
//   	DestinationPrefixLists: []*string{
//   		jsii.String("destinationPrefixLists"),
//   	},
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	SourceAddresses: []*string{
//   		jsii.String("sourceAddresses"),
//   	},
//   	SourcePorts: []*string{
//   		jsii.String("sourcePorts"),
//   	},
//   	SourcePrefixLists: []*string{
//   		jsii.String("sourcePrefixLists"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html
//
type CfnNetworkInsightsAccessScope_PacketHeaderStatementRequestProperty struct {
	// The destination addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-destinationaddresses
	//
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// The destination ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-destinationports
	//
	DestinationPorts *[]*string `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// The destination prefix lists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-destinationprefixlists
	//
	DestinationPrefixLists *[]*string `field:"optional" json:"destinationPrefixLists" yaml:"destinationPrefixLists"`
	// The protocols.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-protocols
	//
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The source addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-sourceaddresses
	//
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// The source ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-sourceports
	//
	SourcePorts *[]*string `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
	// The source prefix lists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.html#cfn-ec2-networkinsightsaccessscope-packetheaderstatementrequest-sourceprefixlists
	//
	SourcePrefixLists *[]*string `field:"optional" json:"sourcePrefixLists" yaml:"sourcePrefixLists"`
}

