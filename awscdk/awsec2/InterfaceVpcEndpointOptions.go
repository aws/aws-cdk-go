package awsec2


// Options to add an interface endpoint to a VPC.
//
// Example:
//   // Add gateway endpoints when creating the VPC
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &VpcProps{
//   	GatewayEndpoints: map[string]gatewayVpcEndpointOptions{
//   		"S3": &gatewayVpcEndpointOptions{
//   			"service": ec2.GatewayVpcEndpointAwsService_S3(),
//   		},
//   	},
//   })
//
//   // Alternatively gateway endpoints can be added on the VPC
//   dynamoDbEndpoint := vpc.addGatewayEndpoint(jsii.String("DynamoDbEndpoint"), &gatewayVpcEndpointOptions{
//   	Service: ec2.GatewayVpcEndpointAwsService_DYNAMODB(),
//   })
//
//   // This allows to customize the endpoint policy
//   dynamoDbEndpoint.AddToPolicy(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	 // Restrict to listing and describing tables
//   	Principals: []iPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   	Actions: []*string{
//   		jsii.String("dynamodb:DescribeTable"),
//   		jsii.String("dynamodb:ListTables"),
//   	},
//   	Resources: []*string{
//   		jsii.String("*"),
//   	},
//   }))
//
//   // Add an interface endpoint
//   vpc.addInterfaceEndpoint(jsii.String("EcrDockerEndpoint"), &InterfaceVpcEndpointOptions{
//   	Service: ec2.InterfaceVpcEndpointAwsService_ECR_DOCKER(),
//   })
//
type InterfaceVpcEndpointOptions struct {
	// The service to use for this interface VPC endpoint.
	Service IInterfaceVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Limit to only those availability zones where the endpoint service can be created.
	//
	// Setting this to 'true' requires a lookup to be performed at synthesis time. Account
	// and region must be set on the containing stack for this to work.
	// Default: false.
	//
	LookupSupportedAzs *bool `field:"optional" json:"lookupSupportedAzs" yaml:"lookupSupportedAzs"`
	// Whether to automatically allow VPC traffic to the endpoint.
	//
	// If enabled, all traffic to the endpoint from within the VPC will be
	// automatically allowed. This is done based on the VPC's CIDR range.
	// Default: true.
	//
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// Whether to associate a private hosted zone with the specified VPC.
	//
	// This
	// allows you to make requests to the service using its default DNS hostname.
	// Default: set by the instance of IInterfaceVpcEndpointService, or true if
	// not defined by the instance of IInterfaceVpcEndpointService.
	//
	PrivateDnsEnabled *bool `field:"optional" json:"privateDnsEnabled" yaml:"privateDnsEnabled"`
	// The security groups to associate with this interface VPC endpoint.
	// Default: - a new security group is created.
	//
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets in which to create an endpoint network interface.
	//
	// At most one
	// per availability zone.
	// Default: - private subnets.
	//
	Subnets *SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

