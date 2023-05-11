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
type CfnDataflowEndpointGroup_AwsGroundStationAgentEndpointProperty struct {
	// `CfnDataflowEndpointGroup.AwsGroundStationAgentEndpointProperty.AgentStatus`.
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// `CfnDataflowEndpointGroup.AwsGroundStationAgentEndpointProperty.AuditResults`.
	AuditResults *string `field:"optional" json:"auditResults" yaml:"auditResults"`
	// `CfnDataflowEndpointGroup.AwsGroundStationAgentEndpointProperty.EgressAddress`.
	EgressAddress interface{} `field:"optional" json:"egressAddress" yaml:"egressAddress"`
	// `CfnDataflowEndpointGroup.AwsGroundStationAgentEndpointProperty.IngressAddress`.
	IngressAddress interface{} `field:"optional" json:"ingressAddress" yaml:"ingressAddress"`
	// `CfnDataflowEndpointGroup.AwsGroundStationAgentEndpointProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

