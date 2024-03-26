package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An AWS service for an interface VPC endpoint.
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
type InterfaceVpcEndpointAwsService interface {
	IInterfaceVpcEndpointService
	// The name of the service.
	//
	// e.g. com.amazonaws.us-east-1.ecs
	Name() *string
	// The port of the service.
	Port() *float64
	// Whether Private DNS is supported by default.
	PrivateDnsDefault() *bool
	// The short name of the service.
	//
	// e.g. ecs
	ShortName() *string
}

// The jsii proxy struct for InterfaceVpcEndpointAwsService
type jsiiProxy_InterfaceVpcEndpointAwsService struct {
	jsiiProxy_IInterfaceVpcEndpointService
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) PrivateDnsDefault() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privateDnsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpointAwsService) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}


func NewInterfaceVpcEndpointAwsService(name *string, prefix *string, port *float64) InterfaceVpcEndpointAwsService {
	_init_.Initialize()

	if err := validateNewInterfaceVpcEndpointAwsServiceParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_InterfaceVpcEndpointAwsService{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		[]interface{}{name, prefix, port},
		&j,
	)

	return &j
}

func NewInterfaceVpcEndpointAwsService_Override(i InterfaceVpcEndpointAwsService, name *string, prefix *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		[]interface{}{name, prefix, port},
		i,
	)
}

func InterfaceVpcEndpointAwsService_ACCESS_ANALYZER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ACCESS_ANALYZER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ACCOUNT_MANAGEMENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ACCOUNT_MANAGEMENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AIRFLOW_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AIRFLOW_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AIRFLOW_ENV() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AIRFLOW_ENV",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AIRFLOW_OPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AIRFLOW_OPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APIGATEWAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APIGATEWAY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APP_MESH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APP_MESH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APP_MESH_ENVOY_MANAGEMENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APP_MESH_ENVOY_MANAGEMENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APP_MESH_OPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APP_MESH_OPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APP_RUNNER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APP_RUNNER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APP_RUNNER_REQUESTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APP_RUNNER_REQUESTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APP_SYNC() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APP_SYNC",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPCONFIG() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPCONFIG",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPCONFIGDATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPCONFIGDATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPLICATION_AUTOSCALING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPLICATION_AUTOSCALING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPLICATION_MIGRATION_SERVICE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPLICATION_MIGRATION_SERVICE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPSTREAM_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPSTREAM_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_APPSTREAM_STREAMING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"APPSTREAM_STREAMING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ATHENA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ATHENA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AUDIT_MANAGER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AUDIT_MANAGER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AUTOSCALING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AUTOSCALING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_AUTOSCALING_PLANS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"AUTOSCALING_PLANS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_B2B_DATA_INTERCHANGE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"B2B_DATA_INTERCHANGE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BACKUP() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BACKUP",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BACKUP_GATEWAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BACKUP_GATEWAY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BATCH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BATCH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BEDROCK() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BEDROCK",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BEDROCK_AGENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BEDROCK_AGENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BEDROCK_AGENT_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BEDROCK_AGENT_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BEDROCK_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BEDROCK_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BILLING_CONDUCTOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BILLING_CONDUCTOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_BRAKET() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"BRAKET",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLEAN_ROOMS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLEAN_ROOMS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_CONTROL_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_CONTROL_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_CONTROL_API_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_CONTROL_API_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_DIRECTORY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_DIRECTORY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_MAP_DATA_SERVICE_DISCOVERY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_MAP_DATA_SERVICE_DISCOVERY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_MAP_DATA_SERVICE_DISCOVERY_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_MAP_DATA_SERVICE_DISCOVERY_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_MAP_SERVICE_DISCOVERY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_MAP_SERVICE_DISCOVERY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUD_MAP_SERVICE_DISCOVERY_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUD_MAP_SERVICE_DISCOVERY_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDFORMATION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDFORMATION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDHSM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDHSM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDTRAIL() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDTRAIL",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_EVENTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_EVENTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_EVIDENTLY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_EVIDENTLY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_EVIDENTLY_DATAPLANE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_EVIDENTLY_DATAPLANE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_LOGS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_LOGS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_MONITORING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_MONITORING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_NETWORK_MONITOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_NETWORK_MONITOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_RUM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_RUM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_RUM_DATAPLANE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_RUM_DATAPLANE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CLOUDWATCH_SYNTHETICS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CLOUDWATCH_SYNTHETICS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEARTIFACT_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEARTIFACT_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEARTIFACT_REPOSITORIES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEARTIFACT_REPOSITORIES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEBUILD() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEBUILD",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEBUILD_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEBUILD_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECATALYST_GIT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECATALYST_GIT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECATALYST_PACKAGES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECATALYST_PACKAGES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT_GIT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT_GIT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODECOMMIT_GIT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODECOMMIT_GIT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEDEPLOY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEDEPLOY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEDEPLOY_COMMANDS_SECURE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEDEPLOY_COMMANDS_SECURE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEGURU_PROFILER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEGURU_PROFILER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEGURU_REVIEWER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEGURU_REVIEWER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEPIPELINE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEPIPELINE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODESTAR_CONNECTIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODESTAR_CONNECTIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CODEWHISPERER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CODEWHISPERER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_COMPREHEND() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"COMPREHEND",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_COMPREHEND_MEDICAL() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"COMPREHEND_MEDICAL",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONFIG() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONFIG",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONNECT_APP_INTEGRATIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONNECT_APP_INTEGRATIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONNECT_CASES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONNECT_CASES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONNECT_CONNECT_CAMPAIGNS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONNECT_CONNECT_CAMPAIGNS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONNECT_PROFILE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONNECT_PROFILE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONNECT_VOICEID() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONNECT_VOICEID",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_CONNECT_WISDOM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"CONNECT_WISDOM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DATA_EXCHANGE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DATA_EXCHANGE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DATABASE_MIGRATION_SERVICE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DATABASE_MIGRATION_SERVICE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DATABASE_MIGRATION_SERVICE_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DATABASE_MIGRATION_SERVICE_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DATASYNC() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DATASYNC",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DATAZONE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DATAZONE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DEVOPS_GURU() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DEVOPS_GURU",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_DIRECTORY_SERVICE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"DIRECTORY_SERVICE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EBS_DIRECT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EBS_DIRECT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EC2() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EC2",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EC2_MESSAGES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EC2_MESSAGES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECR_DOCKER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECR_DOCKER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECS_AGENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECS_AGENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ECS_TELEMETRY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ECS_TELEMETRY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EKS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EKS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EKS_AUTH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EKS_AUTH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_BEANSTALK() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_BEANSTALK",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_BEANSTALK_HEALTH() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_BEANSTALK_HEALTH",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_DISASTER_RECOVERY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_DISASTER_RECOVERY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_FILESYSTEM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_FILESYSTEM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_FILESYSTEM_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_FILESYSTEM_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_INFERENCE_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_INFERENCE_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTIC_LOAD_BALANCING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTIC_LOAD_BALANCING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTICACHE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTICACHE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELASTICACHE_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELASTICACHE_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ELEMENTAL_MEDIACONNECT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ELEMENTAL_MEDIACONNECT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EMAIL_SMTP() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EMAIL_SMTP",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EMR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EMR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EMR_EKS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EMR_EKS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EMR_SERVERLESS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EMR_SERVERLESS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EMR_WAL() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EMR_WAL",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ENTITY_RESOLUTION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ENTITY_RESOLUTION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_EVENTBRIDGE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"EVENTBRIDGE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FAULT_INJECTION_SIMULATOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FAULT_INJECTION_SIMULATOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FINSPACE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FINSPACE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FINSPACE_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FINSPACE_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FORECAST() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FORECAST",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FORECAST_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FORECAST_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FORECAST_QUERY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FORECAST_QUERY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FORECAST_QUERY_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FORECAST_QUERY_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FRAUD_DETECTOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FRAUD_DETECTOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FSX() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FSX",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_FSX_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"FSX_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GLUE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GLUE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GLUE_DATABREW() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GLUE_DATABREW",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GRAFANA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GRAFANA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GRAFANA_WORKSPACE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GRAFANA_WORKSPACE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GROUNDSTATION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GROUNDSTATION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GUARDDUTY_DATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GUARDDUTY_DATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_GUARDDUTY_DATA_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"GUARDDUTY_DATA_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_HEALTH_IMAGING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"HEALTH_IMAGING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_HEALTH_IMAGING_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"HEALTH_IMAGING_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_HEALTHLAKE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"HEALTHLAKE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IAM_IDENTITY_CENTER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IAM_IDENTITY_CENTER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IAM_ROLES_ANYWHERE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IAM_ROLES_ANYWHERE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IMAGE_BUILDER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IMAGE_BUILDER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_INSPECTOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"INSPECTOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_INSPECTOR_SCAN() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"INSPECTOR_SCAN",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_CORE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_CORE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_CORE_CREDENTIALS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_CORE_CREDENTIALS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_CORE_DEVICE_ADVISOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_CORE_DEVICE_ADVISOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_CORE_FLEETHUB_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_CORE_FLEETHUB_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_CORE_FOR_LORAWAN() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_CORE_FOR_LORAWAN",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_FLEETWISE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_FLEETWISE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_GREENGRASS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_GREENGRASS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_LORAWAN_CUPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_LORAWAN_CUPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_LORAWAN_LNS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_LORAWAN_LNS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_ROBORUNNER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_ROBORUNNER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_SITEWISE_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_SITEWISE_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_SITEWISE_DATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_SITEWISE_DATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_TWINMAKER_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_TWINMAKER_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_IOT_TWINMAKER_DATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"IOT_TWINMAKER_DATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KENDRA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KENDRA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KENDRA_RANKING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KENDRA_RANKING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KEYSPACES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KEYSPACES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KEYSPACES_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KEYSPACES_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KINESIS_FIREHOSE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KINESIS_FIREHOSE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KINESIS_STREAMS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KINESIS_STREAMS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KMS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KMS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_KMS_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"KMS_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LAKE_FORMATION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LAKE_FORMATION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LAMBDA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LAMBDA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LEX_MODELS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LEX_MODELS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LEX_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LEX_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LICENSE_MANAGER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LICENSE_MANAGER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LICENSE_MANAGER_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LICENSE_MANAGER_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LICENSE_MANAGER_USER_SUBSCRIPTIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LICENSE_MANAGER_USER_SUBSCRIPTIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LOOKOUT_EQUIPMENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LOOKOUT_EQUIPMENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LOOKOUT_METRICS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LOOKOUT_METRICS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_LOOKOUT_VISION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"LOOKOUT_VISION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MACIE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MACIE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MAINFRAME_MODERNIZATION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MAINFRAME_MODERNIZATION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MANAGED_BLOCKCHAIN_BITCOIN_MAINNET() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MANAGED_BLOCKCHAIN_BITCOIN_MAINNET",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MANAGED_BLOCKCHAIN_BITCOIN_TESTNET() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MANAGED_BLOCKCHAIN_BITCOIN_TESTNET",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MANAGED_BLOCKCHAIN_QUERY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MANAGED_BLOCKCHAIN_QUERY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MANAGEMENT_CONSOLE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MANAGEMENT_CONSOLE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MANAGEMENT_CONSOLE_SIGNIN() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MANAGEMENT_CONSOLE_SIGNIN",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MEMORY_DB() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MEMORY_DB",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MEMORY_DB_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MEMORY_DB_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MIGRATIONHUB_ORCHESTRATOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MIGRATIONHUB_ORCHESTRATOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MIGRATIONHUB_REFACTOR_SPACES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MIGRATIONHUB_REFACTOR_SPACES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_MIGRATIONHUB_STRATEGY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"MIGRATIONHUB_STRATEGY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_NEPTUNE_ANALYTICS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"NEPTUNE_ANALYTICS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_NIMBLE_STUDIO() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"NIMBLE_STUDIO",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_OMICS_ANALYTICS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"OMICS_ANALYTICS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_OMICS_CONTROL_STORAGE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"OMICS_CONTROL_STORAGE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_OMICS_STORAGE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"OMICS_STORAGE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_OMICS_TAGS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"OMICS_TAGS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_OMICS_WORKFLOWS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"OMICS_WORKFLOWS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ORGANIZATIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ORGANIZATIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ORGANIZATIONS_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ORGANIZATIONS_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PANORAMA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PANORAMA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PAYMENT_CRYPTOGRAPHY_CONTROLPLANE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PAYMENT_CRYPTOGRAPHY_CONTROLPLANE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PAYMENT_CRYTOGRAPHY_DATAPLANE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PAYMENT_CRYTOGRAPHY_DATAPLANE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PERSONALIZE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PERSONALIZE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PERSONALIZE_EVENTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PERSONALIZE_EVENTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PERSONALIZE_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PERSONALIZE_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PINPOINT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PINPOINT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PINPOINT_SMS_VOICE_V2() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PINPOINT_SMS_VOICE_V2",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PINPOINT_V1() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PINPOINT_V1",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_POLLY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"POLLY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PRIVATE_5G() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PRIVATE_5G",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PRIVATE_CERTIFICATE_AUTHORITY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PRIVATE_CERTIFICATE_AUTHORITY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PRIVATE_CERTIFICATE_AUTHORITY_CONNECTOR_AD() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PRIVATE_CERTIFICATE_AUTHORITY_CONNECTOR_AD",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PROMETHEUS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PROMETHEUS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PROMETHEUS_WORKSPACES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PROMETHEUS_WORKSPACES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_PROTON() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"PROTON",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_QLDB() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"QLDB",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_RDS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"RDS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_RDS_DATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"RDS_DATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REDSHIFT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REDSHIFT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REDSHIFT_DATA() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REDSHIFT_DATA",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REDSHIFT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REDSHIFT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REKOGNITION() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REKOGNITION",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REKOGNITION_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REKOGNITION_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REKOGNITION_STREAMING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REKOGNITION_STREAMING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REKOGNITION_STREAMING_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REKOGNITION_STREAMING_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_REPOST_SPACE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"REPOST_SPACE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_ROBOMAKER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"ROBOMAKER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_S3() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"S3",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_S3_MULTI_REGION_ACCESS_POINTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"S3_MULTI_REGION_ACCESS_POINTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_S3_OUTPOSTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"S3_OUTPOSTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_API() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_API",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_FEATURESTORE_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_FEATURESTORE_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_GEOSPATIAL() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_GEOSPATIAL",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_METRICS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_METRICS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_NOTEBOOK() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_NOTEBOOK",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_RUNTIME() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_RUNTIME",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_RUNTIME_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_RUNTIME_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SAGEMAKER_STUDIO() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SAGEMAKER_STUDIO",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SECRETS_MANAGER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SECRETS_MANAGER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SECURITYHUB() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SECURITYHUB",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SERVER_MIGRATION_SERVICE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SERVER_MIGRATION_SERVICE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SERVER_MIGRATION_SERVICE_AWSCONNECTOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SERVER_MIGRATION_SERVICE_AWSCONNECTOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SERVER_MIGRATION_SERVICE_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SERVER_MIGRATION_SERVICE_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SERVICE_CATALOG() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SERVICE_CATALOG",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SERVICE_CATALOG_APPREGISTRY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SERVICE_CATALOG_APPREGISTRY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SIMSPACE_WEAVER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SIMSPACE_WEAVER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SNOW_DEVICE_MANAGEMENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SNOW_DEVICE_MANAGEMENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SNS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SNS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SQS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SQS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SSM() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SSM",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SSM_CONTACTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SSM_CONTACTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SSM_INCIDENTS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SSM_INCIDENTS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SSM_MESSAGES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SSM_MESSAGES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STEP_FUNCTIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STEP_FUNCTIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STEP_FUNCTIONS_SYNC() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STEP_FUNCTIONS_SYNC",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STORAGE_GATEWAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STORAGE_GATEWAY",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_STS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"STS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SUPPLY_CHAIN() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SUPPLY_CHAIN",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SWF() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SWF",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_SWF_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"SWF_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TELCO_NETWORK_BUILDER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TELCO_NETWORK_BUILDER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TEXTRACT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TEXTRACT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TEXTRACT_FIPS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TEXTRACT_FIPS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TIMESTREAM_INFLUXDB() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TIMESTREAM_INFLUXDB",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSCRIBE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSCRIBE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSCRIBE_STREAMING() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSCRIBE_STREAMING",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSFER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSFER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSFER_SERVER() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSFER_SERVER",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRANSLATE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRANSLATE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_TRUSTED_ADVISOR() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"TRUSTED_ADVISOR",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_VERIFIED_PERMISSIONS() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"VERIFIED_PERMISSIONS",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_VPC_LATTICE() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"VPC_LATTICE",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_WORKSPACES() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"WORKSPACES",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_WORKSPACES_THIN_CLIENT() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"WORKSPACES_THIN_CLIENT",
		&returns,
	)
	return returns
}

func InterfaceVpcEndpointAwsService_XRAY() InterfaceVpcEndpointAwsService {
	_init_.Initialize()
	var returns InterfaceVpcEndpointAwsService
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		"XRAY",
		&returns,
	)
	return returns
}

