package awsec2


// Describes a path statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathStatementRequestProperty := &PathStatementRequestProperty{
//   	PacketHeaderStatement: &PacketHeaderStatementRequestProperty{
//   		DestinationAddresses: []*string{
//   			jsii.String("destinationAddresses"),
//   		},
//   		DestinationPorts: []*string{
//   			jsii.String("destinationPorts"),
//   		},
//   		DestinationPrefixLists: []*string{
//   			jsii.String("destinationPrefixLists"),
//   		},
//   		Protocols: []*string{
//   			jsii.String("protocols"),
//   		},
//   		SourceAddresses: []*string{
//   			jsii.String("sourceAddresses"),
//   		},
//   		SourcePorts: []*string{
//   			jsii.String("sourcePorts"),
//   		},
//   		SourcePrefixLists: []*string{
//   			jsii.String("sourcePrefixLists"),
//   		},
//   	},
//   	ResourceStatement: &ResourceStatementRequestProperty{
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.html
//
type CfnNetworkInsightsAccessScope_PathStatementRequestProperty struct {
	// The packet header statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.html#cfn-ec2-networkinsightsaccessscope-pathstatementrequest-packetheaderstatement
	//
	PacketHeaderStatement interface{} `field:"optional" json:"packetHeaderStatement" yaml:"packetHeaderStatement"`
	// The resource statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.html#cfn-ec2-networkinsightsaccessscope-pathstatementrequest-resourcestatement
	//
	ResourceStatement interface{} `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

