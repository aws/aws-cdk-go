package awsgroundstation


// Details for a downlink agent endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   downlinkAwsGroundStationAgentEndpointDetailsProperty := &DownlinkAwsGroundStationAgentEndpointDetailsProperty{
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
//
//   	// the properties below are optional
//   	AgentStatus: jsii.String("agentStatus"),
//   	AuditResults: jsii.String("auditResults"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails.html
//
type CfnDataflowEndpointGroupV2_DownlinkAwsGroundStationAgentEndpointDetailsProperty struct {
	// Dataflow details for the downlink endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails-dataflowdetails
	//
	DataflowDetails interface{} `field:"required" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Downlink dataflow endpoint name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Status of the agent associated with the downlink dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails-agentstatus
	//
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// Health audit results for the downlink dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-downlinkawsgroundstationagentendpointdetails-auditresults
	//
	AuditResults *string `field:"optional" json:"auditResults" yaml:"auditResults"`
}

