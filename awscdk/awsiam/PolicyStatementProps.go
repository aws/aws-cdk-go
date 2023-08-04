package awsiam


// Interface for creating a policy statement.
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
type PolicyStatementProps struct {
	// List of actions to add to the statement.
	// Default: - no actions.
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Conditions to add to the statement.
	// Default: - no condition.
	//
	Conditions *map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether to allow or deny the actions in this statement.
	// Default: Effect.ALLOW
	//
	Effect Effect `field:"optional" json:"effect" yaml:"effect"`
	// List of not actions to add to the statement.
	// Default: - no not-actions.
	//
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// List of not principals to add to the statement.
	// Default: - no not principals.
	//
	NotPrincipals *[]IPrincipal `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// NotResource ARNs to add to the statement.
	// Default: - no not-resources.
	//
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// List of principals to add to the statement.
	// Default: - no principals.
	//
	Principals *[]IPrincipal `field:"optional" json:"principals" yaml:"principals"`
	// Resource ARNs to add to the statement.
	// Default: - no resources.
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The Sid (statement ID) is an optional identifier that you provide for the policy statement.
	//
	// You can assign a Sid value to each statement in a
	// statement array. In services that let you specify an ID element, such as
	// SQS and SNS, the Sid value is just a sub-ID of the policy document's ID. In
	// IAM, the Sid value must be unique within a JSON policy.
	// Default: - no sid.
	//
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

