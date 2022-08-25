package awsiam


// Interface for creating a policy statement.
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
type PolicyStatementProps struct {
	// List of actions to add to the statement.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Conditions to add to the statement.
	Conditions *map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether to allow or deny the actions in this statement.
	Effect Effect `field:"optional" json:"effect" yaml:"effect"`
	// List of not actions to add to the statement.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// List of not principals to add to the statement.
	NotPrincipals *[]IPrincipal `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// NotResource ARNs to add to the statement.
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// List of principals to add to the statement.
	Principals *[]IPrincipal `field:"optional" json:"principals" yaml:"principals"`
	// Resource ARNs to add to the statement.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The Sid (statement ID) is an optional identifier that you provide for the policy statement.
	//
	// You can assign a Sid value to each statement in a
	// statement array. In services that let you specify an ID element, such as
	// SQS and SNS, the Sid value is just a sub-ID of the policy document's ID. In
	// IAM, the Sid value must be unique within a JSON policy.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

