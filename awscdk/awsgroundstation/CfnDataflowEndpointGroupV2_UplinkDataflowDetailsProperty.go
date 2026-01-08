package awsgroundstation


// Dataflow details for an uplink endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkDataflowDetailsProperty := &UplinkDataflowDetailsProperty{
//   	AgentConnectionDetails: &UplinkConnectionDetailsProperty{
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
//   		IngressAddressAndPort: &ConnectionDetailsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkdataflowdetails.html
//
type CfnDataflowEndpointGroupV2_UplinkDataflowDetailsProperty struct {
	// Uplink connection details for customer to Agent and Agent to Ground Station.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkdataflowdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkdataflowdetails-agentconnectiondetails
	//
	AgentConnectionDetails interface{} `field:"required" json:"agentConnectionDetails" yaml:"agentConnectionDetails"`
}

