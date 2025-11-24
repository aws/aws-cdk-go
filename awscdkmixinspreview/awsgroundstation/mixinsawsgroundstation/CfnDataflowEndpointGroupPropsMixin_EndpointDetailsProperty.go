package mixinsawsgroundstation


// The security details and endpoint information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointDetailsProperty := &EndpointDetailsProperty{
//   	AwsGroundStationAgentEndpoint: &AwsGroundStationAgentEndpointProperty{
//   		AgentStatus: jsii.String("agentStatus"),
//   		AuditResults: jsii.String("auditResults"),
//   		EgressAddress: &ConnectionDetailsProperty{
//   			Mtu: jsii.Number(123),
//   			SocketAddress: &SocketAddressProperty{
//   				Name: jsii.String("name"),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   		IngressAddress: &RangedConnectionDetailsProperty{
//   			Mtu: jsii.Number(123),
//   			SocketAddress: &RangedSocketAddressProperty{
//   				Name: jsii.String("name"),
//   				PortRange: &IntegerRangeProperty{
//   					Maximum: jsii.Number(123),
//   					Minimum: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//   	},
//   	Endpoint: &DataflowEndpointProperty{
//   		Address: &SocketAddressProperty{
//   			Name: jsii.String("name"),
//   			Port: jsii.Number(123),
//   		},
//   		Mtu: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	SecurityDetails: &SecurityDetailsProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-endpointdetails.html
//
type CfnDataflowEndpointGroupPropsMixin_EndpointDetailsProperty struct {
	// An agent endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-endpointdetails.html#cfn-groundstation-dataflowendpointgroup-endpointdetails-awsgroundstationagentendpoint
	//
	AwsGroundStationAgentEndpoint interface{} `field:"optional" json:"awsGroundStationAgentEndpoint" yaml:"awsGroundStationAgentEndpoint"`
	// Information about the endpoint such as name and the endpoint address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-endpointdetails.html#cfn-groundstation-dataflowendpointgroup-endpointdetails-endpoint
	//
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The role ARN, and IDs for security groups and subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-endpointdetails.html#cfn-groundstation-dataflowendpointgroup-endpointdetails-securitydetails
	//
	SecurityDetails interface{} `field:"optional" json:"securityDetails" yaml:"securityDetails"`
}

