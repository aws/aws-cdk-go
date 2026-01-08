package awsgroundstation


// Dataflow details for a downlink endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   downlinkDataflowDetailsProperty := &DownlinkDataflowDetailsProperty{
//   	AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   		AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   			SocketAddress: &RangedSocketAddressProperty{
//   				Name: jsii.String("name"),
//   				PortRange: &IntegerRangeProperty{
//   					Maximum: jsii.Number(123),
//   					Minimum: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Mtu: jsii.Number(123),
//   		},
//   		EgressAddressAndPort: &ConnectionDetailsProperty{
//   			SocketAddress: &SocketAddressProperty{
//   				Name: jsii.String("name"),
//   				Port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			Mtu: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkdataflowdetails.html
//
type CfnDataflowEndpointGroupV2_DownlinkDataflowDetailsProperty struct {
	// Downlink connection details for customer to Agent and Agent to Ground Station.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkdataflowdetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkdataflowdetails-agentconnectiondetails
	//
	AgentConnectionDetails interface{} `field:"required" json:"agentConnectionDetails" yaml:"agentConnectionDetails"`
}

