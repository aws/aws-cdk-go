package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsGroundStationAgentEndpointProperty := &AwsGroundStationAgentEndpointProperty{
//   	AgentStatus: jsii.String("agentStatus"),
//   	AuditResults: jsii.String("auditResults"),
//   	EgressAddress: &ConnectionDetailsProperty{
//   		Mtu: jsii.Number(123),
//   		SocketAddress: &SocketAddressProperty{
//   			Name: jsii.String("name"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	IngressAddress: &RangedConnectionDetailsProperty{
//   		Mtu: jsii.Number(123),
//   		SocketAddress: &RangedSocketAddressProperty{
//   			Name: jsii.String("name"),
//   			PortRange: &IntegerRangeProperty{
//   				Maximum: jsii.Number(123),
//   				Minimum: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.html
//
type CfnDataflowEndpointGroup_AwsGroundStationAgentEndpointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint-agentstatus
	//
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint-auditresults
	//
	AuditResults *string `field:"optional" json:"auditResults" yaml:"auditResults"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint-egressaddress
	//
	EgressAddress interface{} `field:"optional" json:"egressAddress" yaml:"egressAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint-ingressaddress
	//
	IngressAddress interface{} `field:"optional" json:"ingressAddress" yaml:"ingressAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.html#cfn-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

