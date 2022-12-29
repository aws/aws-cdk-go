package awsec2


// Options to add an interface endpoint to a VPC.
//
// Example:
//   // Add gateway endpoints when creating the VPC
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
//   	gatewayEndpoints: map[string]gatewayVpcEndpointOptions{
//   		"S3": &gatewayVpcEndpointOptions{
//   			"service": ec2.GatewayVpcEndpointAwsService_S3(),
//   		},
//   	},
//   })
//
//   // Alternatively gateway endpoints can be added on the VPC
//   dynamoDbEndpoint := vpc.addGatewayEndpoint(jsii.String("DynamoDbEndpoint"), &gatewayVpcEndpointOptions{
//   	service: ec2.gatewayVpcEndpointAwsService_DYNAMODB(),
//   })
//
//   // This allows to customize the endpoint policy
//   dynamoDbEndpoint.addToPolicy(
//   iam.NewPolicyStatement(&policyStatementProps{
//   	 // Restrict to listing and describing tables
//   	principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	actions: []*string{
//   		jsii.String("dynamodb:DescribeTable"),
//   		jsii.String("dynamodb:ListTables"),
//   	},
//   	resources: []*string{
//   		jsii.String("*"),
//   	},
//   }))
//
//   // Add an interface endpoint
//   vpc.addInterfaceEndpoint(jsii.String("EcrDockerEndpoint"), &interfaceVpcEndpointOptions{
//   	service: ec2.interfaceVpcEndpointAwsService_ECR_DOCKER(),
//   })
//
type InterfaceVpcEndpointOptions struct {
	// The service to use for this interface VPC endpoint.
	Service IInterfaceVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Limit to only those availability zones where the endpoint service can be created.
	//
	// Setting this to 'true' requires a lookup to be performed at synthesis time. Account
	// and region must be set on the containing stack for this to work.
	LookupSupportedAzs *bool `field:"optional" json:"lookupSupportedAzs" yaml:"lookupSupportedAzs"`
	// Whether to automatically allow VPC traffic to the endpoint.
	//
	// If enabled, all traffic to the endpoint from within the VPC will be
	// automatically allowed. This is done based on the VPC's CIDR range.
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// Whether to associate a private hosted zone with the specified VPC.
	//
	// This
	// allows you to make requests to the service using its default DNS hostname.
	PrivateDnsEnabled *bool `field:"optional" json:"privateDnsEnabled" yaml:"privateDnsEnabled"`
	// The security groups to associate with this interface VPC endpoint.
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets in which to create an endpoint network interface.
	//
	// At most one
	// per availability zone.
	Subnets *SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

