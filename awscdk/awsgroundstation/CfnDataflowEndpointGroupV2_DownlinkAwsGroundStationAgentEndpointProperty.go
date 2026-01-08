package awsgroundstation


// Definition for a downlink agent endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   downlinkAwsGroundStationAgentEndpointProperty := &DownlinkAwsGroundStationAgentEndpointProperty{
//   	DataflowDetails: &DownlinkDataflowDetailsProperty{
//   		AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   			AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   				SocketAddress: &RangedSocketAddressProperty{
//   					Name: jsii.String("name"),
//   					PortRange: &IntegerRangeProperty{
//   						Maximum: jsii.Number(123),
//   						Minimum: jsii.Number(123),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Mtu: jsii.Number(123),
//   			},
//   			EgressAddressAndPort: &ConnectionDetailsProperty{
//   				SocketAddress: &SocketAddressProperty{
//   					Name: jsii.String("name"),
//   					Port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Mtu: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html
//
type CfnDataflowEndpointGroupV2_DownlinkAwsGroundStationAgentEndpointProperty struct {
	// Dataflow details for the downlink endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint-dataflowdetails
	//
	DataflowDetails interface{} `field:"required" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Downlink dataflow endpoint name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

