package previewawsgroundstationmixins


// Endpoint definition used for creating a dataflow endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createEndpointDetailsProperty := &CreateEndpointDetailsProperty{
//   	DownlinkAwsGroundStationAgentEndpoint: &DownlinkAwsGroundStationAgentEndpointProperty{
//   		DataflowDetails: &DownlinkDataflowDetailsProperty{
//   			AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   				AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   					Mtu: jsii.Number(123),
//   					SocketAddress: &RangedSocketAddressProperty{
//   						Name: jsii.String("name"),
//   						PortRange: &IntegerRangeProperty{
//   							Maximum: jsii.Number(123),
//   							Minimum: jsii.Number(123),
//   						},
//   					},
//   				},
//   				EgressAddressAndPort: &ConnectionDetailsProperty{
//   					Mtu: jsii.Number(123),
//   					SocketAddress: &SocketAddressProperty{
//   						Name: jsii.String("name"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//   	},
//   	UplinkAwsGroundStationAgentEndpoint: &UplinkAwsGroundStationAgentEndpointProperty{
//   		DataflowDetails: &UplinkDataflowDetailsProperty{
//   			AgentConnectionDetails: &UplinkConnectionDetailsProperty{
//   				AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   					Mtu: jsii.Number(123),
//   					SocketAddress: &RangedSocketAddressProperty{
//   						Name: jsii.String("name"),
//   						PortRange: &IntegerRangeProperty{
//   							Maximum: jsii.Number(123),
//   							Minimum: jsii.Number(123),
//   						},
//   					},
//   				},
//   				IngressAddressAndPort: &ConnectionDetailsProperty{
//   					Mtu: jsii.Number(123),
//   					SocketAddress: &SocketAddressProperty{
//   						Name: jsii.String("name"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-createendpointdetails.html
//
type CfnDataflowEndpointGroupV2PropsMixin_CreateEndpointDetailsProperty struct {
	// Definition for a downlink agent endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-createendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-createendpointdetails-downlinkawsgroundstationagentendpoint
	//
	DownlinkAwsGroundStationAgentEndpoint interface{} `field:"optional" json:"downlinkAwsGroundStationAgentEndpoint" yaml:"downlinkAwsGroundStationAgentEndpoint"`
	// Definition for an uplink agent endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-createendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-createendpointdetails-uplinkawsgroundstationagentendpoint
	//
	UplinkAwsGroundStationAgentEndpoint interface{} `field:"optional" json:"uplinkAwsGroundStationAgentEndpoint" yaml:"uplinkAwsGroundStationAgentEndpoint"`
}

