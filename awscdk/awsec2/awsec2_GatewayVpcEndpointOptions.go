package awsec2


// Options to add a gateway endpoint to a VPC.
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
type GatewayVpcEndpointOptions struct {
	// The service to use for this gateway VPC endpoint.
	Service IGatewayVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Where to add endpoint routing.
	//
	// By default, this endpoint will be routable from all subnets in the VPC.
	// Specify a list of subnet selection objects here to be more specific.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   var vpc vpc
	//
	//
	//   vpc.addGatewayEndpoint(jsii.String("DynamoDbEndpoint"), &gatewayVpcEndpointOptions{
	//   	service: ec2.gatewayVpcEndpointAwsService_DYNAMODB(),
	//   	// Add only to ISOLATED subnets
	//   	subnets: []subnetSelection{
	//   		&subnetSelection{
	//   			subnetType: ec2.subnetType_PRIVATE_ISOLATED,
	//   		},
	//   	},
	//   })
	//
	Subnets *[]*SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

