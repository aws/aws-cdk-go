package awsec2


// Describes a path statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathStatementRequestProperty := &pathStatementRequestProperty{
//   	packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   		destinationAddresses: []*string{
//   			jsii.String("destinationAddresses"),
//   		},
//   		destinationPorts: []*string{
//   			jsii.String("destinationPorts"),
//   		},
//   		destinationPrefixLists: []*string{
//   			jsii.String("destinationPrefixLists"),
//   		},
//   		protocols: []*string{
//   			jsii.String("protocols"),
//   		},
//   		sourceAddresses: []*string{
//   			jsii.String("sourceAddresses"),
//   		},
//   		sourcePorts: []*string{
//   			jsii.String("sourcePorts"),
//   		},
//   		sourcePrefixLists: []*string{
//   			jsii.String("sourcePrefixLists"),
//   		},
//   	},
//   	resourceStatement: &resourceStatementRequestProperty{
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		resourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAccessScope_PathStatementRequestProperty struct {
	// The packet header statement.
	PacketHeaderStatement interface{} `field:"optional" json:"packetHeaderStatement" yaml:"packetHeaderStatement"`
	// The resource statement.
	ResourceStatement interface{} `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

