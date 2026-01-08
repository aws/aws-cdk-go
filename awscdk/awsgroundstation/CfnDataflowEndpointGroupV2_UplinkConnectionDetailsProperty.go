package awsgroundstation


// Connection details for customer to Agent and Agent to Ground Station.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkConnectionDetailsProperty := &UplinkConnectionDetailsProperty{
//   	AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   		SocketAddress: &RangedSocketAddressProperty{
//   			Name: jsii.String("name"),
//   			PortRange: &IntegerRangeProperty{
//   				Maximum: jsii.Number(123),
//   				Minimum: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Mtu: jsii.Number(123),
//   	},
//   	IngressAddressAndPort: &ConnectionDetailsProperty{
//   		SocketAddress: &SocketAddressProperty{
//   			Name: jsii.String("name"),
//   			Port: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		Mtu: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkconnectiondetails.html
//
type CfnDataflowEndpointGroupV2_UplinkConnectionDetailsProperty struct {
	// Agent IP and port address for the uplink connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkconnectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkconnectiondetails-agentipandportaddress
	//
	AgentIpAndPortAddress interface{} `field:"required" json:"agentIpAndPortAddress" yaml:"agentIpAndPortAddress"`
	// Ingress address and port for the uplink connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkconnectiondetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkconnectiondetails-ingressaddressandport
	//
	IngressAddressAndPort interface{} `field:"required" json:"ingressAddressAndPort" yaml:"ingressAddressAndPort"`
}

