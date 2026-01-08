package previewawsgroundstationmixins


// Definition for an uplink agent endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uplinkAwsGroundStationAgentEndpointProperty := &UplinkAwsGroundStationAgentEndpointProperty{
//   	DataflowDetails: &UplinkDataflowDetailsProperty{
//   		AgentConnectionDetails: &UplinkConnectionDetailsProperty{
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
//   			IngressAddressAndPort: &ConnectionDetailsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint.html
//
type CfnDataflowEndpointGroupV2PropsMixin_UplinkAwsGroundStationAgentEndpointProperty struct {
	// Dataflow details for the uplink endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint-dataflowdetails
	//
	DataflowDetails interface{} `field:"optional" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Uplink dataflow endpoint name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

