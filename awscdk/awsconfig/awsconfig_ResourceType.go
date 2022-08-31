package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Resources types that are supported by AWS Config.
//
// Example:
//   var evalComplianceFn function
//   sshRule := config.NewManagedRule(this, jsii.String("SSH"), &managedRuleProps{
//   	identifier: config.managedRuleIdentifiers_EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED(),
//   	ruleScope: config.ruleScope.fromResource(config.resourceType_EC2_SECURITY_GROUP(), jsii.String("sg-1234567890abcdefgh")),
//   })
//   customRule := config.NewCustomRule(this, jsii.String("Lambda"), &customRuleProps{
//   	lambdaFunction: evalComplianceFn,
//   	configurationChanges: jsii.Boolean(true),
//   	ruleScope: config.*ruleScope.fromResources([]*resourceType{
//   		config.*resourceType_CLOUDFORMATION_STACK(),
//   		config.*resourceType_S3_BUCKET(),
//   	}),
//   })
//
//   tagRule := config.NewCustomRule(this, jsii.String("CostCenterTagRule"), &customRuleProps{
//   	lambdaFunction: evalComplianceFn,
//   	configurationChanges: jsii.Boolean(true),
//   	ruleScope: config.*ruleScope.fromTag(jsii.String("Cost Center"), jsii.String("MyApp")),
//   })
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html
//
// Experimental.
type ResourceType interface {
	// Valid value of resource type.
	// Experimental.
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
// Experimental.
func ResourceType_Of(type_ *string) ResourceType {
	_init_.Initialize()

	var returns ResourceType

	_jsii_.StaticInvoke(
		"monocdk.aws_config.ResourceType",
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
		"monocdk.aws_config.ResourceType",
		"ACM_CERTIFICATE",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAY_REST_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAY_REST_API",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAY_STAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAY_STAGE",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAYV2_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAYV2_API",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAYV2_STAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAYV2_STAGE",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_LAUNCH_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_LAUNCH_CONFIGURATION",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_SCHEDULED_ACTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_SCHEDULED_ACTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFORMATION_STACK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDFORMATION_STACK",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFRONT_DISTRIBUTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDFRONT_DISTRIBUTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFRONT_STREAMING_DISTRIBUTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDFRONT_STREAMING_DISTRIBUTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDTRAIL_TRAIL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDTRAIL_TRAIL",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDWATCH_ALARM() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDWATCH_ALARM",
		&returns,
	)
	return returns
}

func ResourceType_CODEBUILD_PROJECT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CODEBUILD_PROJECT",
		&returns,
	)
	return returns
}

func ResourceType_CODEPIPELINE_PIPELINE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CODEPIPELINE_PIPELINE",
		&returns,
	)
	return returns
}

func ResourceType_DYNAMODB_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"DYNAMODB_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EBS_VOLUME() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EBS_VOLUME",
		&returns,
	)
	return returns
}

func ResourceType_EC2_CUSTOMER_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_CUSTOMER_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_EGRESS_ONLY_INTERNET_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_EGRESS_ONLY_INTERNET_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_EIP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_EIP",
		&returns,
	)
	return returns
}

func ResourceType_EC2_FLOW_LOG() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_FLOW_LOG",
		&returns,
	)
	return returns
}

func ResourceType_EC2_HOST() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_HOST",
		&returns,
	)
	return returns
}

func ResourceType_EC2_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_INTERNET_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_INTERNET_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NAT_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_NAT_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NETWORK_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_NETWORK_ACL",
		&returns,
	)
	return returns
}

func ResourceType_EC2_ROUTE_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_ROUTE_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_EC2_SUBNET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_SUBNET",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_ENDPOINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC_ENDPOINT",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_ENDPOINT_SERVICE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC_ENDPOINT_SERVICE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_PEERING_CONNECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC_PEERING_CONNECTION",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPN_CONNECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPN_CONNECTION",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPN_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPN_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_APPLICATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_APPLICATION",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_APPLICATION_VERSION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_APPLICATION_VERSION",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_ENVIRONMENT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_ENVIRONMENT",
		&returns,
	)
	return returns
}

func ResourceType_ELASTICSEARCH_DOMAIN() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTICSEARCH_DOMAIN",
		&returns,
	)
	return returns
}

func ResourceType_ELB_LOAD_BALANCER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELB_LOAD_BALANCER",
		&returns,
	)
	return returns
}

func ResourceType_ELBV2_LOAD_BALANCER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELBV2_LOAD_BALANCER",
		&returns,
	)
	return returns
}

func ResourceType_IAM_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_IAM_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_IAM_ROLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_ROLE",
		&returns,
	)
	return returns
}

func ResourceType_IAM_USER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_USER",
		&returns,
	)
	return returns
}

func ResourceType_KMS_KEY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"KMS_KEY",
		&returns,
	)
	return returns
}

func ResourceType_LAMBDA_FUNCTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"LAMBDA_FUNCTION",
		&returns,
	)
	return returns
}

func ResourceType_QLDB_LEDGER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"QLDB_LEDGER",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_CLUSTER_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_CLUSTER_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_RDS_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_PARAMETER_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_PARAMETER_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_S3_ACCOUNT_PUBLIC_ACCESS_BLOCK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"S3_ACCOUNT_PUBLIC_ACCESS_BLOCK",
		&returns,
	)
	return returns
}

func ResourceType_S3_BUCKET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"S3_BUCKET",
		&returns,
	)
	return returns
}

func ResourceType_SECRETS_MANAGER_SECRET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SECRETS_MANAGER_SECRET",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_CLOUDFORMATION_PRODUCT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SERVICE_CATALOG_CLOUDFORMATION_PRODUCT",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_CLOUDFORMATION_PROVISIONED_PRODUCT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SERVICE_CATALOG_CLOUDFORMATION_PROVISIONED_PRODUCT",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_PORTFOLIO() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SERVICE_CATALOG_PORTFOLIO",
		&returns,
	)
	return returns
}

func ResourceType_SHIELD_PROTECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SHIELD_PROTECTION",
		&returns,
	)
	return returns
}

func ResourceType_SHIELD_REGIONAL_PROTECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SHIELD_REGIONAL_PROTECTION",
		&returns,
	)
	return returns
}

func ResourceType_SNS_TOPIC() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SNS_TOPIC",
		&returns,
	)
	return returns
}

func ResourceType_SQS_QUEUE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SQS_QUEUE",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_ASSOCIATION_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_ASSOCIATION_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_FILE_DATA() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_FILE_DATA",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_MANAGED_INSTANCE_INVENTORY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_MANAGED_INSTANCE_INVENTORY",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_PATCH_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_PATCH_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RATE_BASED_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_RATE_BASED_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RATE_BASED_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_RATE_BASED_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAF_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_MANAGED_RULE_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAFV2_MANAGED_RULE_SET",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAFV2_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAFV2_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_XRAY_ENCRYPTION_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"XRAY_ENCRYPTION_CONFIGURATION",
		&returns,
	)
	return returns
}

