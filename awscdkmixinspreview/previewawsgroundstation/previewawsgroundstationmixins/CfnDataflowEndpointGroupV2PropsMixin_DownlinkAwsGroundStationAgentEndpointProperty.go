package previewawsgroundstationmixins


// Definition for a downlink agent endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   downlinkAwsGroundStationAgentEndpointProperty := &DownlinkAwsGroundStationAgentEndpointProperty{
//   	DataflowDetails: &DownlinkDataflowDetailsProperty{
//   		AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   			AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   				Mtu: jsii.Number(123),
//   				SocketAddress: &RangedSocketAddressProperty{
//   					Name: jsii.String("name"),
//   					PortRange: &IntegerRangeProperty{
//   						Maximum: jsii.Number(123),
//   						Minimum: jsii.Number(123),
//   					},
//   				},
//   			},
//   			EgressAddressAndPort: &ConnectionDetailsProperty{
//   				Mtu: jsii.Number(123),
//   				SocketAddress: &SocketAddressProperty{
//   					Name: jsii.String("name"),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html
//
type CfnDataflowEndpointGroupV2PropsMixin_DownlinkAwsGroundStationAgentEndpointProperty struct {
	// Dataflow details for the downlink endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint-dataflowdetails
	//
	DataflowDetails interface{} `field:"optional" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Downlink dataflow endpoint name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

