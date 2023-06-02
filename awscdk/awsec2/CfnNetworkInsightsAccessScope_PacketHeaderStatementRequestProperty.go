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
type CfnNetworkInsightsAccessScope_PacketHeaderStatementRequestProperty struct {
	// The destination addresses.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// The destination ports.
	DestinationPorts *[]*string `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// The destination prefix lists.
	DestinationPrefixLists *[]*string `field:"optional" json:"destinationPrefixLists" yaml:"destinationPrefixLists"`
	// The protocols.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The source addresses.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// The source ports.
	SourcePorts *[]*string `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
	// The source prefix lists.
	SourcePrefixLists *[]*string `field:"optional" json:"sourcePrefixLists" yaml:"sourcePrefixLists"`
}

