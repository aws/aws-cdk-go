package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Resources types that are supported by AWS Config.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
//   	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
//   	ConfigurationChanges: jsii.Boolean(true),
//   	LambdaFunction: evalComplianceFn,
//   	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
//   })
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html
//
type ResourceType interface {
	// Valid value of resource type.
	ComplianceResourceType() *string
}

// The jsii proxy struct for ResourceType
type jsiiProxy_ResourceType struct {
	_ byte // padding
}

func (j *jsiiProxy_ResourceType) ComplianceResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceResourceType",
		&returns,
	)
	return returns
}


// A custom resource type to support future cases.
func ResourceType_Of(type_ *string) ResourceType {
	_init_.Initialize()

	if err := validateResourceType_OfParameters(type_); err != nil {
		panic(err)
	}
	var returns ResourceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.ResourceType",
		"of",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func ResourceType_ACM_CERTIFICATE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ACM_CERTIFICATE",
		&returns,
	)
	return returns
}

func ResourceType_AMAZON_MQ_BROKER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"AMAZON_MQ_BROKER",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAY_REST_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APIGATEWAY_REST_API",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAY_STAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APIGATEWAY_STAGE",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAYV2_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APIGATEWAYV2_API",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAYV2_STAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APIGATEWAYV2_STAGE",
		&returns,
	)
	return returns
}

func ResourceType_APPCONFIG_APPLICATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APPCONFIG_APPLICATION",
		&returns,
	)
	return returns
}

func ResourceType_APPCONFIG_CONFIGURATION_PROFILE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APPCONFIG_CONFIGURATION_PROFILE",
		&returns,
	)
	return returns
}

func ResourceType_APPCONFIG_ENVIRONMENT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APPCONFIG_ENVIRONMENT",
		&returns,
	)
	return returns
}

func ResourceType_APPSYNC_GRAPHQL_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"APPSYNC_GRAPHQL_API",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"AUTO_SCALING_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_LAUNCH_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"AUTO_SCALING_LAUNCH_CONFIGURATION",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"AUTO_SCALING_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_SCHEDULED_ACTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"AUTO_SCALING_SCHEDULED_ACTION",
		&returns,
	)
	return returns
}

func ResourceType_BACKUP_BACKUP_PLAN() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BACKUP_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ResourceType_BACKUP_BACKUP_SELECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BACKUP_BACKUP_SELECTION",
		&returns,
	)
	return returns
}

func ResourceType_BACKUP_BACKUP_VAULT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BACKUP_BACKUP_VAULT",
		&returns,
	)
	return returns
}

func ResourceType_BACKUP_RECOVERY_POINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BACKUP_RECOVERY_POINT",
		&returns,
	)
	return returns
}

func ResourceType_BACKUP_REPORT_PLAN() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BACKUP_REPORT_PLAN",
		&returns,
	)
	return returns
}

func ResourceType_BATCH_COMPUTE_ENVIRONMENT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BATCH_COMPUTE_ENVIRONMENT",
		&returns,
	)
	return returns
}

func ResourceType_BATCH_JOB_QUEUE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"BATCH_JOB_QUEUE",
		&returns,
	)
	return returns
}

func ResourceType_CLOUD9_ENVIRONMENT_EC2() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUD9_ENVIRONMENT_EC2",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFORMATION_STACK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUDFORMATION_STACK",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFRONT_DISTRIBUTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUDFRONT_DISTRIBUTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFRONT_STREAMING_DISTRIBUTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUDFRONT_STREAMING_DISTRIBUTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDTRAIL_TRAIL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUDTRAIL_TRAIL",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDWATCH_ALARM() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUDWATCH_ALARM",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDWATCH_RUM_APP_MONITOR() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CLOUDWATCH_RUM_APP_MONITOR",
		&returns,
	)
	return returns
}

func ResourceType_CODEBUILD_PROJECT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CODEBUILD_PROJECT",
		&returns,
	)
	return returns
}

func ResourceType_CODEDEPLOY_APPLICATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CODEDEPLOY_APPLICATION",
		&returns,
	)
	return returns
}

func ResourceType_CODEDEPLOY_DEPLOYMENT_CONFIG() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CODEDEPLOY_DEPLOYMENT_CONFIG",
		&returns,
	)
	return returns
}

func ResourceType_CODEDEPLOY_DEPLOYMENT_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CODEDEPLOY_DEPLOYMENT_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_CODEPIPELINE_PIPELINE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CODEPIPELINE_PIPELINE",
		&returns,
	)
	return returns
}

func ResourceType_CONFIG_CONFORMANCE_PACK_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CONFIG_CONFORMANCE_PACK_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_CONFIG_RESOURCE_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"CONFIG_RESOURCE_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_EFS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_EFS",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_FSX_LUSTRE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_FSX_LUSTRE",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_FSX_WINDOWS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_FSX_WINDOWS",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_HDFS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_HDFS",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_NFS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_NFS",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_OBJECT_STORAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_OBJECT_STORAGE",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_S3() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_S3",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_LOCATION_SMB() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_LOCATION_SMB",
		&returns,
	)
	return returns
}

func ResourceType_DATASYNC_TASK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DATASYNC_TASK",
		&returns,
	)
	return returns
}

func ResourceType_DMS_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DMS_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_DMS_REPLICATION_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DMS_REPLICATION_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_DYNAMODB_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"DYNAMODB_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EBS_VOLUME() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EBS_VOLUME",
		&returns,
	)
	return returns
}

func ResourceType_EC2_CUSTOMER_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_CUSTOMER_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_EGRESS_ONLY_INTERNET_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_EGRESS_ONLY_INTERNET_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_EIP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_EIP",
		&returns,
	)
	return returns
}

func ResourceType_EC2_FLOW_LOG() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_FLOW_LOG",
		&returns,
	)
	return returns
}

func ResourceType_EC2_HOST() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_HOST",
		&returns,
	)
	return returns
}

func ResourceType_EC2_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_INTERNET_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_INTERNET_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_LAUNCH_TEMPLATE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_LAUNCH_TEMPLATE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NAT_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_NAT_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NETWORK_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_NETWORK_ACL",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NETWORK_INSIGHTS_ACCESS_SCOPE_ANALYSIS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_NETWORK_INSIGHTS_ACCESS_SCOPE_ANALYSIS",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NETWORK_INTERFACE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_NETWORK_INTERFACE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_REGISTERED_HA_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_REGISTERED_HA_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_ROUTE_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_ROUTE_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_EC2_SUBNET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_SUBNET",
		&returns,
	)
	return returns
}

func ResourceType_EC2_TRANSIT_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_TRANSIT_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_TRANSIT_GATEWAY_ATTACHMENT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_TRANSIT_GATEWAY_ATTACHMENT",
		&returns,
	)
	return returns
}

func ResourceType_EC2_TRANSIT_GATEWAY_ROUTE_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_TRANSIT_GATEWAY_ROUTE_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_VPC",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_ENDPOINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_VPC_ENDPOINT",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_ENDPOINT_SERVICE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_VPC_ENDPOINT_SERVICE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_PEERING_CONNECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_VPC_PEERING_CONNECTION",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPN_CONNECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_VPN_CONNECTION",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPN_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EC2_VPN_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_ECR_PUBLIC_REPOSITORY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ECR_PUBLIC_REPOSITORY",
		&returns,
	)
	return returns
}

func ResourceType_ECR_REGISTRY_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ECR_REGISTRY_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_ECR_REPOSITORY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ECR_REPOSITORY",
		&returns,
	)
	return returns
}

func ResourceType_ECS_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ECS_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_ECS_SERVICE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ECS_SERVICE",
		&returns,
	)
	return returns
}

func ResourceType_ECS_TASK_DEFINITION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ECS_TASK_DEFINITION",
		&returns,
	)
	return returns
}

func ResourceType_EFS_ACCESS_POINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EFS_ACCESS_POINT",
		&returns,
	)
	return returns
}

func ResourceType_EFS_FILE_SYSTEM() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EFS_FILE_SYSTEM",
		&returns,
	)
	return returns
}

func ResourceType_EKS_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EKS_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_APPLICATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_APPLICATION",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_APPLICATION_VERSION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_APPLICATION_VERSION",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_ENVIRONMENT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_ENVIRONMENT",
		&returns,
	)
	return returns
}

func ResourceType_ELASTICSEARCH_DOMAIN() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELASTICSEARCH_DOMAIN",
		&returns,
	)
	return returns
}

func ResourceType_ELB_LOAD_BALANCER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELB_LOAD_BALANCER",
		&returns,
	)
	return returns
}

func ResourceType_ELBV2_LISTENER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELBV2_LISTENER",
		&returns,
	)
	return returns
}

func ResourceType_ELBV2_LOAD_BALANCER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ELBV2_LOAD_BALANCER",
		&returns,
	)
	return returns
}

func ResourceType_EMR_SECURITY_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EMR_SECURITY_CONFIGURATION",
		&returns,
	)
	return returns
}

func ResourceType_EVENTBRIDGE_API_DESTINATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTBRIDGE_API_DESTINATION",
		&returns,
	)
	return returns
}

func ResourceType_EVENTBRIDGE_ARCHIVE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTBRIDGE_ARCHIVE",
		&returns,
	)
	return returns
}

func ResourceType_EVENTBRIDGE_ENDPOINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTBRIDGE_ENDPOINT",
		&returns,
	)
	return returns
}

func ResourceType_EVENTBRIDGE_EVENTBUS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTBRIDGE_EVENTBUS",
		&returns,
	)
	return returns
}

func ResourceType_EVENTSCHEMAS_DISCOVERER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTSCHEMAS_DISCOVERER",
		&returns,
	)
	return returns
}

func ResourceType_EVENTSCHEMAS_REGISTRY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTSCHEMAS_REGISTRY",
		&returns,
	)
	return returns
}

func ResourceType_EVENTSCHEMAS_REGISTRY_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"EVENTSCHEMAS_REGISTRY_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_FIS_EXPERIMENT_TEMPLATE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"FIS_EXPERIMENT_TEMPLATE",
		&returns,
	)
	return returns
}

func ResourceType_FRAUDDETECTOR_ENTITY_TYPE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"FRAUDDETECTOR_ENTITY_TYPE",
		&returns,
	)
	return returns
}

func ResourceType_FRAUDDETECTOR_LABEL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"FRAUDDETECTOR_LABEL",
		&returns,
	)
	return returns
}

func ResourceType_FRAUDDETECTOR_OUTCOME() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"FRAUDDETECTOR_OUTCOME",
		&returns,
	)
	return returns
}

func ResourceType_FRAUDDETECTOR_VARIABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"FRAUDDETECTOR_VARIABLE",
		&returns,
	)
	return returns
}

func ResourceType_GLOBALACCELERATOR_ACCELERATOR() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GLOBALACCELERATOR_ACCELERATOR",
		&returns,
	)
	return returns
}

func ResourceType_GLOBALACCELERATOR_ENDPOINT_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GLOBALACCELERATOR_ENDPOINT_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_GLOBALACCELERATOR_LISTENER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GLOBALACCELERATOR_LISTENER",
		&returns,
	)
	return returns
}

func ResourceType_GLUE_CLASSIFIER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GLUE_CLASSIFIER",
		&returns,
	)
	return returns
}

func ResourceType_GLUE_JOB() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GLUE_JOB",
		&returns,
	)
	return returns
}

func ResourceType_GUARDDUTY_DETECTOR() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GUARDDUTY_DETECTOR",
		&returns,
	)
	return returns
}

func ResourceType_GUARDDUTY_FILTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GUARDDUTY_FILTER",
		&returns,
	)
	return returns
}

func ResourceType_GUARDDUTY_IP_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GUARDDUTY_IP_SET",
		&returns,
	)
	return returns
}

func ResourceType_GUARDDUTY_THREAT_INTEL_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"GUARDDUTY_THREAT_INTEL_SET",
		&returns,
	)
	return returns
}

func ResourceType_IAM_ACCESSANALYZER_ANALYZER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IAM_ACCESSANALYZER_ANALYZER",
		&returns,
	)
	return returns
}

func ResourceType_IAM_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IAM_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_IAM_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IAM_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_IAM_ROLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IAM_ROLE",
		&returns,
	)
	return returns
}

func ResourceType_IAM_USER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IAM_USER",
		&returns,
	)
	return returns
}

func ResourceType_IMAGEBUILDER_CONTAINER_RECIPE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IMAGEBUILDER_CONTAINER_RECIPE",
		&returns,
	)
	return returns
}

func ResourceType_IMAGEBUILDER_DISTRIBUTION_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IMAGEBUILDER_DISTRIBUTION_CONFIGURATION",
		&returns,
	)
	return returns
}

func ResourceType_IMAGEBUILDER_INFRASTRUCTURE_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IMAGEBUILDER_INFRASTRUCTURE_CONFIGURATION",
		&returns,
	)
	return returns
}

func ResourceType_IOT_ANALYTICS_DATASTORE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_ANALYTICS_DATASTORE",
		&returns,
	)
	return returns
}

func ResourceType_IOT_AUTHORIZER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_AUTHORIZER",
		&returns,
	)
	return returns
}

func ResourceType_IOT_DIMENSION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_DIMENSION",
		&returns,
	)
	return returns
}

func ResourceType_IOT_EVENTS_ALARM_MODEL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_EVENTS_ALARM_MODEL",
		&returns,
	)
	return returns
}

func ResourceType_IOT_EVENTS_DETECTOR_MODEL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_EVENTS_DETECTOR_MODEL",
		&returns,
	)
	return returns
}

func ResourceType_IOT_EVENTS_INPUT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_EVENTS_INPUT",
		&returns,
	)
	return returns
}

func ResourceType_IOT_ROLE_ALIAS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_ROLE_ALIAS",
		&returns,
	)
	return returns
}

func ResourceType_IOT_SECURITY_PROFILE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"IOT_SECURITY_PROFILE",
		&returns,
	)
	return returns
}

func ResourceType_KINESIS_STREAM() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"KINESIS_STREAM",
		&returns,
	)
	return returns
}

func ResourceType_KINESIS_STREAM_CONSUMER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"KINESIS_STREAM_CONSUMER",
		&returns,
	)
	return returns
}

func ResourceType_KMS_KEY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"KMS_KEY",
		&returns,
	)
	return returns
}

func ResourceType_LAMBDA_FUNCTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"LAMBDA_FUNCTION",
		&returns,
	)
	return returns
}

func ResourceType_LIGHTSAIL_BUCKET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"LIGHTSAIL_BUCKET",
		&returns,
	)
	return returns
}

func ResourceType_LIGHTSAIL_CERTIFICATE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"LIGHTSAIL_CERTIFICATE",
		&returns,
	)
	return returns
}

func ResourceType_LIGHTSAIL_DISK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"LIGHTSAIL_DISK",
		&returns,
	)
	return returns
}

func ResourceType_LIGHTSAIL_STATIC_IP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"LIGHTSAIL_STATIC_IP",
		&returns,
	)
	return returns
}

func ResourceType_MEDIAPACKAGE_PACKAGING_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"MEDIAPACKAGE_PACKAGING_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_MSK_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"MSK_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_NETWORK_FIREWALL_FIREWALL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"NETWORK_FIREWALL_FIREWALL",
		&returns,
	)
	return returns
}

func ResourceType_NETWORK_FIREWALL_FIREWALL_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"NETWORK_FIREWALL_FIREWALL_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_NETWORK_FIREWALL_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"NETWORK_FIREWALL_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_OPENSEARCH_DOMAIN() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"OPENSEARCH_DOMAIN",
		&returns,
	)
	return returns
}

func ResourceType_QLDB_LEDGER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"QLDB_LEDGER",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_DB_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_CLUSTER_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_DB_CLUSTER_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_DB_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_DB_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_DB_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_DB_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_RDS_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RDS_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"REDSHIFT_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_PARAMETER_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_PARAMETER_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"REDSHIFT_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_RESILIENCEHUB_RESILIENCY_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"RESILIENCEHUB_RESILIENCY_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_HEALTH_CHECK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_HEALTH_CHECK",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_HOSTED_ZONE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_HOSTED_ZONE",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_RECOVERY_READINESS_CELL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_RECOVERY_READINESS_CELL",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_RECOVERY_READINESS_READINESS_CHECK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_RECOVERY_READINESS_READINESS_CHECK",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_RECOVERY_READINESS_RECOVERY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_RECOVERY_READINESS_RECOVERY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_RESOLVER_RESOLVER_ENDPOINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_RESOLVER_RESOLVER_ENDPOINT",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_RESOLVER_RESOLVER_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_RESOLVER_RESOLVER_RULE",
		&returns,
	)
	return returns
}

func ResourceType_ROUTE53_RESOLVER_RESOLVER_RULE_ASSOCIATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"ROUTE53_RESOLVER_RESOLVER_RULE_ASSOCIATION",
		&returns,
	)
	return returns
}

func ResourceType_S3_ACCOUNT_PUBLIC_ACCESS_BLOCK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"S3_ACCOUNT_PUBLIC_ACCESS_BLOCK",
		&returns,
	)
	return returns
}

func ResourceType_S3_BUCKET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"S3_BUCKET",
		&returns,
	)
	return returns
}

func ResourceType_SAGEMAKER_CODE_REPOSITORY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SAGEMAKER_CODE_REPOSITORY",
		&returns,
	)
	return returns
}

func ResourceType_SAGEMAKER_MODEL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SAGEMAKER_MODEL",
		&returns,
	)
	return returns
}

func ResourceType_SAGEMAKER_NOTEBOOK_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SAGEMAKER_NOTEBOOK_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_SAGEMAKER_WORKTEAM() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SAGEMAKER_WORKTEAM",
		&returns,
	)
	return returns
}

func ResourceType_SECRETS_MANAGER_SECRET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SECRETS_MANAGER_SECRET",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_CLOUDFORMATION_PRODUCT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SERVICE_CATALOG_CLOUDFORMATION_PRODUCT",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_CLOUDFORMATION_PROVISIONED_PRODUCT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SERVICE_CATALOG_CLOUDFORMATION_PROVISIONED_PRODUCT",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_PORTFOLIO() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SERVICE_CATALOG_PORTFOLIO",
		&returns,
	)
	return returns
}

func ResourceType_SERVICEDISCOVERY_HTTP_NAMESPACE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SERVICEDISCOVERY_HTTP_NAMESPACE",
		&returns,
	)
	return returns
}

func ResourceType_SERVICEDISCOVERY_PUBLIC_DNS_NAMESPACE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SERVICEDISCOVERY_PUBLIC_DNS_NAMESPACE",
		&returns,
	)
	return returns
}

func ResourceType_SERVICEDISCOVERY_SERVICE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SERVICEDISCOVERY_SERVICE",
		&returns,
	)
	return returns
}

func ResourceType_SES_CONFIGURATION_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SES_CONFIGURATION_SET",
		&returns,
	)
	return returns
}

func ResourceType_SES_CONTACT_LIST() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SES_CONTACT_LIST",
		&returns,
	)
	return returns
}

func ResourceType_SES_RECEIPT_FILTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SES_RECEIPT_FILTER",
		&returns,
	)
	return returns
}

func ResourceType_SES_RECEIPT_RECEIPT_RULE_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SES_RECEIPT_RECEIPT_RULE_SET",
		&returns,
	)
	return returns
}

func ResourceType_SES_TEMPLATE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SES_TEMPLATE",
		&returns,
	)
	return returns
}

func ResourceType_SHIELD_PROTECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SHIELD_PROTECTION",
		&returns,
	)
	return returns
}

func ResourceType_SHIELD_REGIONAL_PROTECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SHIELD_REGIONAL_PROTECTION",
		&returns,
	)
	return returns
}

func ResourceType_SNS_TOPIC() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SNS_TOPIC",
		&returns,
	)
	return returns
}

func ResourceType_SQS_QUEUE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SQS_QUEUE",
		&returns,
	)
	return returns
}

func ResourceType_STEPFUNCTIONS_ACTIVITY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"STEPFUNCTIONS_ACTIVITY",
		&returns,
	)
	return returns
}

func ResourceType_STEPFUNCTIONS_STATE_MACHINE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"STEPFUNCTIONS_STATE_MACHINE",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_ASSOCIATION_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SYSTEMS_MANAGER_ASSOCIATION_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_FILE_DATA() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SYSTEMS_MANAGER_FILE_DATA",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_MANAGED_INSTANCE_INVENTORY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SYSTEMS_MANAGER_MANAGED_INSTANCE_INVENTORY",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_PATCH_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"SYSTEMS_MANAGER_PATCH_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_TRANSFER_WORKFLOW() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"TRANSFER_WORKFLOW",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RATE_BASED_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_RATE_BASED_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RATE_BASED_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_REGIONAL_RATE_BASED_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_REGIONAL_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_REGIONAL_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_REGIONAL_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAF_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAF_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_IP_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAFV2_IP_SET",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_MANAGED_RULE_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAFV2_MANAGED_RULE_SET",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_REGEX_PATTERN_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAFV2_REGEX_PATTERN_SET",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAFV2_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WAFV2_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WORKSPACES_CONNECTION_ALIAS() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WORKSPACES_CONNECTION_ALIAS",
		&returns,
	)
	return returns
}

func ResourceType_WORKSPACES_WORKSPACE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"WORKSPACES_WORKSPACE",
		&returns,
	)
	return returns
}

func ResourceType_XRAY_ENCRYPTION_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ResourceType",
		"XRAY_ENCRYPTION_CONFIGURATION",
		&returns,
	)
	return returns
}

