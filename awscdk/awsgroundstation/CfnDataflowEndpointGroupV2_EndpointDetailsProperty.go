package awsgroundstation


// Information about the endpoint details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &EndpointDetailsProperty{
//   	DownlinkAwsGroundStationAgentEndpoint: &DownlinkAwsGroundStationAgentEndpointDetailsProperty{
//   		DataflowDetails: &DownlinkDataflowDetailsProperty{
//   			AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   				AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   					SocketAddress: &RangedSocketAddressProperty{
//   						Name: jsii.String("name"),
//   						PortRange: &IntegerRangeProperty{
//   							Maximum: jsii.Number(123),
//   							Minimum: jsii.Number(123),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Mtu: jsii.Number(123),
//   				},
//   				EgressAddressAndPort: &ConnectionDetailsProperty{
//   					SocketAddress: &SocketAddressProperty{
//   						Name: jsii.String("name"),
//   						Port: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					Mtu: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		AgentStatus: jsii.String("agentStatus"),
//   		AuditResults: jsii.String("auditResults"),
//   	},
//   	UplinkAwsGroundStationAgentEndpoint: &UplinkAwsGroundStationAgentEndpointDetailsProperty{
//   		DataflowDetails: &UplinkDataflowDetailsProperty{
//   			AgentConnectionDetails: &UplinkConnectionDetailsProperty{
//   				AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   					SocketAddress: &RangedSocketAddressProperty{
//   						Name: jsii.String("name"),
//   						PortRange: &IntegerRangeProperty{
//   							Maximum: jsii.Number(123),
//   							Minimum: jsii.Number(123),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Mtu: jsii.Number(123),
//   				},
//   				IngressAddressAndPort: &ConnectionDetailsProperty{
//   					SocketAddress: &SocketAddressProperty{
//   						Name: jsii.String("name"),
//   						Port: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					Mtu: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		AgentStatus: jsii.String("agentStatus"),
//   		AuditResults: jsii.String("auditResults"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-endpointdetails.html
//
type CfnDataflowEndpointGroupV2_EndpointDetailsProperty struct {
	// Definition for a downlink agent endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-endpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-endpointdetails-downlinkawsgroundstationagentendpoint
	//
	DownlinkAwsGroundStationAgentEndpoint interface{} `field:"optional" json:"downlinkAwsGroundStationAgentEndpoint" yaml:"downlinkAwsGroundStationAgentEndpoint"`
	// Definition for an uplink agent endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-endpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-endpointdetails-uplinkawsgroundstationagentendpoint
	//
	UplinkAwsGroundStationAgentEndpoint interface{} `field:"optional" json:"uplinkAwsGroundStationAgentEndpoint" yaml:"uplinkAwsGroundStationAgentEndpoint"`
}

