package previewawsgroundstationmixins


// Details for an uplink agent endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uplinkAwsGroundStationAgentEndpointDetailsProperty := &UplinkAwsGroundStationAgentEndpointDetailsProperty{
//   	AgentStatus: jsii.String("agentStatus"),
//   	AuditResults: jsii.String("auditResults"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html
//
type CfnDataflowEndpointGroupV2PropsMixin_UplinkAwsGroundStationAgentEndpointDetailsProperty struct {
	// Status of the agent associated with the uplink dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-agentstatus
	//
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// Health audit results for the uplink dataflow endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-auditresults
	//
	AuditResults *string `field:"optional" json:"auditResults" yaml:"auditResults"`
	// Dataflow details for the uplink endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-dataflowdetails
	//
	DataflowDetails interface{} `field:"optional" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Uplink dataflow endpoint name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails.html#cfn-groundstation-dataflowendpointgroupv2-uplinkawsgroundstationagentendpointdetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

