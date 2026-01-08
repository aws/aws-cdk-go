package previewawsgroundstationmixins


// Connection details for Ground Station to Agent and Agent to customer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   downlinkConnectionDetailsProperty := &DownlinkConnectionDetailsProperty{
//   	AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   		Mtu: jsii.Number(123),
//   		SocketAddress: &RangedSocketAddressProperty{
//   			Name: jsii.String("name"),
//   			PortRange: &IntegerRangeProperty{
//   				Maximum: jsii.Number(123),
//   				Minimum: jsii.Number(123),
//   			},
//   		},
//   	},
//   	EgressAddressAndPort: &ConnectionDetailsProperty{
//   		Mtu: jsii.Number(123),
//   		SocketAddress: &SocketAddressProperty{
//   			Name: jsii.String("name"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkconnectiondetails.html
//
type CfnDataflowEndpointGroupV2PropsMixin_DownlinkConnectionDetailsProperty struct {
	// Agent IP and port address for the downlink connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkconnectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkconnectiondetails-agentipandportaddress
	//
	AgentIpAndPortAddress interface{} `field:"optional" json:"agentIpAndPortAddress" yaml:"agentIpAndPortAddress"`
	// Egress address and port for the downlink connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkconnectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkconnectiondetails-egressaddressandport
	//
	EgressAddressAndPort interface{} `field:"optional" json:"egressAddressAndPort" yaml:"egressAddressAndPort"`
}

