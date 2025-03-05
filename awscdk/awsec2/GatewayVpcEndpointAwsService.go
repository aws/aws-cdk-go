package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An AWS service for a gateway VPC endpoint.
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
type GatewayVpcEndpointAwsService interface {
	IGatewayVpcEndpointService
	// The name of the service.
	Name() *string
}

// The jsii proxy struct for GatewayVpcEndpointAwsService
type jsiiProxy_GatewayVpcEndpointAwsService struct {
	jsiiProxy_IGatewayVpcEndpointService
}

func (j *jsiiProxy_GatewayVpcEndpointAwsService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewGatewayVpcEndpointAwsService(name *string, prefix *string) GatewayVpcEndpointAwsService {
	_init_.Initialize()

	if err := validateNewGatewayVpcEndpointAwsServiceParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_GatewayVpcEndpointAwsService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointAwsService",
		[]interface{}{name, prefix},
		&j,
	)

	return &j
}

func NewGatewayVpcEndpointAwsService_Override(g GatewayVpcEndpointAwsService, name *string, prefix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointAwsService",
		[]interface{}{name, prefix},
		g,
	)
}

func GatewayVpcEndpointAwsService_DYNAMODB() GatewayVpcEndpointAwsService {
	_init_.Initialize()
	var returns GatewayVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointAwsService",
		"DYNAMODB",
		&returns,
	)
	return returns
}

func GatewayVpcEndpointAwsService_S3() GatewayVpcEndpointAwsService {
	_init_.Initialize()
	var returns GatewayVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointAwsService",
		"S3",
		&returns,
	)
	return returns
}

func GatewayVpcEndpointAwsService_S3_EXPRESS() GatewayVpcEndpointAwsService {
	_init_.Initialize()
	var returns GatewayVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointAwsService",
		"S3_EXPRESS",
		&returns,
	)
	return returns
}

