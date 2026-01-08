package awsgroundstation


// Details for an uplink agent endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uplinkAwsGroundStationAgentEndpointDetailsProperty := &UplinkAwsGroundStationAgentEndpointDetailsProperty{
//   	DataflowDetails: &UplinkDataflowDetailsProperty{
//   		AgentConnectionDetails: &UplinkConnectionDetailsProperty{
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
//   			IngressAddressAndPort: &ConnectionDetailsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html
//
type CfnDataflowEndpointGroupV2_UplinkAwsGroundStationAgentEndpointDetailsProperty struct {
	// Dataflow details for the uplink endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-dataflowdetails
	//
	DataflowDetails interface{} `field:"required" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Uplink dataflow endpoint name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Status of the agent associated with the uplink dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-agentstatus
	//
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// Health audit results for the uplink dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-auditresults
	//
	AuditResults *string `field:"optional" json:"auditResults" yaml:"auditResults"`
}

