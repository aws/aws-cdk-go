package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Managed rules that are supported by AWS Config.
//
// Example:
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   config.NewManagedRule(this, jsii.String("AccessKeysRotated"), &ManagedRuleProps{
//   	Identifier: config.ManagedRuleIdentifiers_ACCESS_KEYS_ROTATED(),
//   	InputParameters: map[string]interface{}{
//   		"maxAccessKeyAge": jsii.Number(60),
//   	},
//
//   	// default is 24 hours
//   	MaximumExecutionFrequency: config.MaximumExecutionFrequency_TWELVE_HOURS,
//   })
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
//
type ManagedRuleIdentifiers interface {
}

// The jsii proxy struct for ManagedRuleIdentifiers
type jsiiProxy_ManagedRuleIdentifiers struct {
	_ byte // padding
}

func ManagedRuleIdentifiers_ACCESS_KEYS_ROTATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ACCESS_KEYS_ROTATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ACCOUNT_PART_OF_ORGANIZATIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ACCOUNT_PART_OF_ORGANIZATIONS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ACM_CERTIFICATE_EXPIRATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ACM_CERTIFICATE_EXPIRATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_DESYNC_MODE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ALB_DESYNC_MODE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_HTTP_DROP_INVALID_HEADER_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ALB_HTTP_DROP_INVALID_HEADER_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_HTTP_TO_HTTPS_REDIRECTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ALB_HTTP_TO_HTTPS_REDIRECTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_WAF_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ALB_WAF_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_ASSOCIATED_WITH_WAF() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GW_ASSOCIATED_WITH_WAF",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_CACHE_ENABLED_AND_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GW_CACHE_ENABLED_AND_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_ENDPOINT_TYPE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GW_ENDPOINT_TYPE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_EXECUTION_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GW_EXECUTION_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_SSL_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GW_SSL_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_XRAY_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GW_XRAY_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GWV2_ACCESS_LOGS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GWV2_ACCESS_LOGS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GWV2_AUTHORIZATION_TYPE_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"API_GWV2_AUTHORIZATION_TYPE_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_APPROVED_AMIS_BY_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"APPROVED_AMIS_BY_ID",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_APPROVED_AMIS_BY_TAG() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"APPROVED_AMIS_BY_TAG",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AURORA_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AURORA_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AURORA_MYSQL_BACKTRACKING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AURORA_MYSQL_BACKTRACKING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AURORA_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AURORA_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_CAPACITY_REBALANCING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_CAPACITY_REBALANCING",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_GROUP_ELB_HEALTHCHECK_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_GROUP_ELB_HEALTHCHECK_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_LAUNCH_CONFIG_HOP_LIMIT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_LAUNCH_CONFIG_HOP_LIMIT",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_LAUNCH_CONFIG_PUBLIC_IP_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_LAUNCH_CONFIG_PUBLIC_IP_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_LAUNCH_TEMPLATE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_LAUNCH_TEMPLATE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_LAUNCHCONFIG_REQUIRES_IMDSV2() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_LAUNCHCONFIG_REQUIRES_IMDSV2",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_MULTIPLE_AZ() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_MULTIPLE_AZ",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_MULTIPLE_INSTANCE_TYPES() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_MULTIPLE_INSTANCE_TYPES",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_BACKUP_PLAN_MIN_FREQUENCY_AND_MIN_RETENTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"BACKUP_PLAN_MIN_FREQUENCY_AND_MIN_RETENTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_BACKUP_RECOVERY_POINT_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"BACKUP_RECOVERY_POINT_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_BACKUP_RECOVERY_POINT_MANUAL_DELETION_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"BACKUP_RECOVERY_POINT_MANUAL_DELETION_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_BACKUP_RECOVERY_POINT_MINIMUM_RETENTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"BACKUP_RECOVERY_POINT_MINIMUM_RETENTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_BEANSTALK_ENHANCED_HEALTH_REPORTING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"BEANSTALK_ENHANCED_HEALTH_REPORTING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLB_DESYNC_MODE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLB_DESYNC_MODE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLB_MULTIPLE_AZ() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLB_MULTIPLE_AZ",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_CLOUD_WATCH_LOGS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_CLOUD_WATCH_LOGS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_LOG_FILE_VALIDATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_LOG_FILE_VALIDATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFORMATION_STACK_DRIFT_DETECTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFORMATION_STACK_DRIFT_DETECTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFORMATION_STACK_NOTIFICATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFORMATION_STACK_NOTIFICATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_ACCESSLOGS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_ACCESSLOGS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_ASSOCIATED_WITH_WAF() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_ASSOCIATED_WITH_WAF",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_CUSTOM_SSL_CERTIFICATE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_CUSTOM_SSL_CERTIFICATE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_DEFAULT_ROOT_OBJECT_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_DEFAULT_ROOT_OBJECT_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_NO_DEPRECATED_SSL_PROTOCOLS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_NO_DEPRECATED_SSL_PROTOCOLS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_ORIGIN_ACCESS_IDENTITY_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_ORIGIN_ACCESS_IDENTITY_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_ORIGIN_FAILOVER_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_ORIGIN_FAILOVER_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_SECURITY_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_SECURITY_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_SNI_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_SNI_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_TRAFFIC_TO_ORIGIN_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_TRAFFIC_TO_ORIGIN_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_VIEWER_POLICY_HTTPS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_VIEWER_POLICY_HTTPS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDTRAIL_MULTI_REGION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDTRAIL_MULTI_REGION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDTRAIL_S3_DATAEVENTS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDTRAIL_S3_DATAEVENTS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDTRAIL_SECURITY_TRAIL_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDTRAIL_SECURITY_TRAIL_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_ACTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_ACTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_ACTION_ENABLED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_ACTION_ENABLED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_RESOURCE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_RESOURCE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_SETTINGS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_SETTINGS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_LOG_GROUP_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_LOG_GROUP_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CMK_BACKING_KEY_ROTATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CMK_BACKING_KEY_ROTATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_ARTIFACT_ENCRYPTION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_ARTIFACT_ENCRYPTION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_ENVIRONMENT_PRIVILEGED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_ENVIRONMENT_PRIVILEGED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_ENVVAR_AWSCRED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_ENVVAR_AWSCRED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_S3_LOGS_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_S3_LOGS_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_SOURCE_REPO_URL_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_SOURCE_REPO_URL_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEDEPLOY_AUTO_ROLLBACK_MONITOR_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEDEPLOY_AUTO_ROLLBACK_MONITOR_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEDEPLOY_EC2_MINIMUM_HEALTHY_HOSTS_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEDEPLOY_EC2_MINIMUM_HEALTHY_HOSTS_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEDEPLOY_LAMBDA_ALLATONCE_TRAFFIC_SHIFT_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEDEPLOY_LAMBDA_ALLATONCE_TRAFFIC_SHIFT_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEPIPELINE_DEPLOYMENT_COUNT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEPIPELINE_DEPLOYMENT_COUNT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEPIPELINE_REGION_FANOUT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CODEPIPELINE_REGION_FANOUT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CW_LOGGROUP_RETENTION_PERIOD_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"CW_LOGGROUP_RETENTION_PERIOD_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DAX_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DAX_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DMS_REPLICATION_NOT_PUBLIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DMS_REPLICATION_NOT_PUBLIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_AUTOSCALING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_AUTOSCALING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_PITR_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_PITR_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_TABLE_ENCRYPTED_KMS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_TABLE_ENCRYPTED_KMS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_TABLE_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_TABLE_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_THROUGHPUT_LIMIT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_THROUGHPUT_LIMIT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_ENCRYPTED_VOLUMES() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EBS_ENCRYPTED_VOLUMES",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EBS_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_OPTIMIZED_INSTANCE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EBS_OPTIMIZED_INSTANCE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EBS_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_SNAPSHOT_PUBLIC_RESTORABLE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EBS_SNAPSHOT_PUBLIC_RESTORABLE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_DESIRED_INSTANCE_TENANCY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_DESIRED_INSTANCE_TENANCY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_DESIRED_INSTANCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_DESIRED_INSTANCE_TYPE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_EBS_ENCRYPTION_BY_DEFAULT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_EBS_ENCRYPTION_BY_DEFAULT",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_IMDSV2_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_IMDSV2_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_DETAILED_MONITORING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_DETAILED_MONITORING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_MANAGED_BY_SSM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_MANAGED_BY_SSM",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_MULTIPLE_ENI_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_MULTIPLE_ENI_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_NO_PUBLIC_IP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_NO_PUBLIC_IP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_PROFILE_ATTACHED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_PROFILE_ATTACHED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCES_IN_VPC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCES_IN_VPC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_APPLICATIONS_BLOCKED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_APPLICATIONS_BLOCKED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_APPLICATIONS_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_APPLICATIONS_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_ASSOCIATION_COMPLIANCE_STATUS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_ASSOCIATION_COMPLIANCE_STATUS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_INVENTORY_BLOCKED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_INVENTORY_BLOCKED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_PATCH_COMPLIANCE_STATUS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_PATCH_COMPLIANCE_STATUS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_PLATFORM_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_PLATFORM_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_NO_AMAZON_KEY_PAIR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_NO_AMAZON_KEY_PAIR",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_PARAVIRTUAL_INSTANCE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_PARAVIRTUAL_INSTANCE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUP_ATTACHED_TO_ENI() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUP_ATTACHED_TO_ENI",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUP_ATTACHED_TO_ENI_PERIODIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUP_ATTACHED_TO_ENI_PERIODIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUPS_RESTRICTED_INCOMING_TRAFFIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUPS_RESTRICTED_INCOMING_TRAFFIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_STOPPED_INSTANCE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_STOPPED_INSTANCE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_TOKEN_HOP_LIMIT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_TOKEN_HOP_LIMIT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_TRANSIT_GATEWAY_AUTO_VPC_ATTACH_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_TRANSIT_GATEWAY_AUTO_VPC_ATTACH_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_VOLUME_IECS_TASK_DEFINITION_USER_FOR_HOST_MODE_CHECKNUSE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_VOLUME_IECS_TASK_DEFINITION_USER_FOR_HOST_MODE_CHECKNUSE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_VOLUME_INUSE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EC2_VOLUME_INUSE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECR_PRIVATE_IMAGE_SCANNING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECR_PRIVATE_IMAGE_SCANNING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECR_PRIVATE_LIFECYCLE_POLICY_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECR_PRIVATE_LIFECYCLE_POLICY_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECR_PRIVATE_TAG_IMMUTABILITY_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECR_PRIVATE_TAG_IMMUTABILITY_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_AWSVPC_NETWORKING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_AWSVPC_NETWORKING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_CONTAINER_INSIGHTS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_CONTAINER_INSIGHTS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_CONTAINERS_NONPRIVILEGED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_CONTAINERS_NONPRIVILEGED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_CONTAINERS_READONLY_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_CONTAINERS_READONLY_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_FARGATE_LATEST_PLATFORM_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_FARGATE_LATEST_PLATFORM_VERSION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_NO_ENVIRONMENT_SECRETS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_NO_ENVIRONMENT_SECRETS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_TASK_DEFINITION_LOG_CONFIGURATION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_TASK_DEFINITION_LOG_CONFIGURATION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_TASK_DEFINITION_MEMORY_HARD_LIMIT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_TASK_DEFINITION_MEMORY_HARD_LIMIT",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_TASK_DEFINITION_NONROOT_USER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_TASK_DEFINITION_NONROOT_USER",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ECS_TASK_DEFINITION_PID_MODE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ECS_TASK_DEFINITION_PID_MODE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_ACCESS_POINT_ENFORCE_ROOT_DIRECTORY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EFS_ACCESS_POINT_ENFORCE_ROOT_DIRECTORY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_ACCESS_POINT_ENFORCE_USER_IDENTITY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EFS_ACCESS_POINT_ENFORCE_USER_IDENTITY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_ENCRYPTED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EFS_ENCRYPTED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EFS_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EFS_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EFS_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EIP_ATTACHED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EIP_ATTACHED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EKS_CLUSTER_OLDEST_SUPPORTED_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EKS_CLUSTER_OLDEST_SUPPORTED_VERSION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EKS_CLUSTER_SUPPORTED_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EKS_CLUSTER_SUPPORTED_VERSION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EKS_ENDPOINT_NO_PUBLIC_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EKS_ENDPOINT_NO_PUBLIC_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EKS_SECRETS_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EKS_SECRETS_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTIC_BEANSTALK_MANAGED_UPDATES_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELASTIC_BEANSTALK_MANAGED_UPDATES_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICACHE_REDIS_CLUSTER_AUTOMATIC_BACKUP_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELASTICACHE_REDIS_CLUSTER_AUTOMATIC_BACKUP_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICSEARCH_ENCRYPTED_AT_REST() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELASTICSEARCH_ENCRYPTED_AT_REST",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICSEARCH_IN_VPC_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELASTICSEARCH_IN_VPC_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICSEARCH_NODE_TO_NODE_ENCRYPTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELASTICSEARCH_NODE_TO_NODE_ENCRYPTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_ACM_CERTIFICATE_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_ACM_CERTIFICATE_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_CROSS_ZONE_LOAD_BALANCING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_CROSS_ZONE_LOAD_BALANCING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_CUSTOM_SECURITY_POLICY_SSL_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_CUSTOM_SECURITY_POLICY_SSL_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_DELETION_PROTECTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_DELETION_PROTECTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_PREDEFINED_SECURITY_POLICY_SSL_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_PREDEFINED_SECURITY_POLICY_SSL_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_TLS_HTTPS_LISTENERS_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELB_TLS_HTTPS_LISTENERS_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELBV2_ACM_CERTIFICATE_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELBV2_ACM_CERTIFICATE_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELBV2_MULTIPLE_AZ() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ELBV2_MULTIPLE_AZ",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EMR_KERBEROS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EMR_KERBEROS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EMR_MASTER_NO_PUBLIC_IP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"EMR_MASTER_NO_PUBLIC_IP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SECURITY_GROUP_AUDIT_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FMS_SECURITY_GROUP_AUDIT_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SECURITY_GROUP_CONTENT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FMS_SECURITY_GROUP_CONTENT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SECURITY_GROUP_RESOURCE_ASSOCIATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FMS_SECURITY_GROUP_RESOURCE_ASSOCIATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SHIELD_RESOURCE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FMS_SHIELD_RESOURCE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_WEBACL_RESOURCE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FMS_WEBACL_RESOURCE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_WEBACL_RULEGROUP_ASSOCIATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FMS_WEBACL_RULEGROUP_ASSOCIATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FSX_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FSX_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FSX_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"FSX_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_GUARDDUTY_ENABLED_CENTRALIZED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"GUARDDUTY_ENABLED_CENTRALIZED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_GUARDDUTY_NON_ARCHIVED_FINDINGS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"GUARDDUTY_NON_ARCHIVED_FINDINGS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_CUSTOMER_POLICY_BLOCKED_KMS_ACTIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_CUSTOMER_POLICY_BLOCKED_KMS_ACTIONS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_GROUP_HAS_USERS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_GROUP_HAS_USERS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_INLINE_POLICY_BLOCKED_KMS_ACTIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_INLINE_POLICY_BLOCKED_KMS_ACTIONS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_NO_INLINE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_NO_INLINE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_PASSWORD_POLICY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_PASSWORD_POLICY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_BLOCKED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_BLOCKED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_IN_USE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_IN_USE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_NO_STATEMENTS_WITH_ADMIN_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_NO_STATEMENTS_WITH_ADMIN_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_NO_STATEMENTS_WITH_FULL_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_NO_STATEMENTS_WITH_FULL_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_ROLE_MANAGED_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_ROLE_MANAGED_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_ROOT_ACCESS_KEY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_ROOT_ACCESS_KEY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_GROUP_MEMBERSHIP_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_GROUP_MEMBERSHIP_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_MFA_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_MFA_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_NO_POLICIES_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_NO_POLICIES_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_UNUSED_CREDENTIALS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_UNUSED_CREDENTIALS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_INTERNET_GATEWAY_AUTHORIZED_VPC_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"INTERNET_GATEWAY_AUTHORIZED_VPC_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_KINESIS_STREAM_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"KINESIS_STREAM_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_KMS_CMK_NOT_SCHEDULED_FOR_DELETION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"KMS_CMK_NOT_SCHEDULED_FOR_DELETION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_CONCURRENCY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_CONCURRENCY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_DLQ_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_DLQ_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_FUNCTION_PUBLIC_ACCESS_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_FUNCTION_PUBLIC_ACCESS_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_FUNCTION_SETTINGS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_FUNCTION_SETTINGS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_INSIDE_VPC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_INSIDE_VPC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_VPC_MULTI_AZ_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_VPC_MULTI_AZ_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_MFA_ENABLED_FOR_IAM_CONSOLE_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"MFA_ENABLED_FOR_IAM_CONSOLE_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_NACL_NO_UNRESTRICTED_SSH_RDP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"NACL_NO_UNRESTRICTED_SSH_RDP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_NETFW_POLICY_DEFAULT_ACTION_FRAGMENT_PACKETS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"NETFW_POLICY_DEFAULT_ACTION_FRAGMENT_PACKETS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_NETFW_POLICY_DEFAULT_ACTION_FULL_PACKETS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"NETFW_POLICY_DEFAULT_ACTION_FULL_PACKETS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_NETFW_POLICY_RULE_GROUP_ASSOCIATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"NETFW_POLICY_RULE_GROUP_ASSOCIATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_NETFW_STATELESS_RULE_GROUP_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"NETFW_STATELESS_RULE_GROUP_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_NLB_CROSS_ZONE_LOAD_BALANCING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"NLB_CROSS_ZONE_LOAD_BALANCING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_ACCESS_CONTROL_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_ACCESS_CONTROL_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_AUDIT_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_AUDIT_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_DATA_NODE_FAULT_TOLERANCE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_DATA_NODE_FAULT_TOLERANCE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_ENCRYPTED_AT_REST() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_ENCRYPTED_AT_REST",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_HTTPS_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_HTTPS_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_IN_VPC_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_IN_VPC_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_LOGS_TO_CLOUDWATCH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_LOGS_TO_CLOUDWATCH",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_OPENSEARCH_NODE_TO_NODE_ENCRYPTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"OPENSEARCH_NODE_TO_NODE_ENCRYPTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_AUTOMATIC_MINOR_VERSION_UPGRADE_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_AUTOMATIC_MINOR_VERSION_UPGRADE_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_CLUSTER_DEFAULT_ADMIN_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_CLUSTER_DEFAULT_ADMIN_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_CLUSTER_DELETION_PROTECTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_CLUSTER_DELETION_PROTECTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_CLUSTER_IAM_AUTHENTICATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_CLUSTER_IAM_AUTHENTICATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_CLUSTER_MULTI_AZ_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_CLUSTER_MULTI_AZ_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_DB_INSTANCE_BACKUP_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_DB_INSTANCE_BACKUP_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_DB_SECURITY_GROUP_NOT_ALLOWED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_DB_SECURITY_GROUP_NOT_ALLOWED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_ENHANCED_MONITORING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_ENHANCED_MONITORING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_DEFAULT_ADMIN_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_DEFAULT_ADMIN_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_DELETION_PROTECTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_DELETION_PROTECTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_IAM_AUTHENTICATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_IAM_AUTHENTICATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_PUBLIC_ACCESS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_PUBLIC_ACCESS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_MULTI_AZ_SUPPORT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_MULTI_AZ_SUPPORT",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_SNAPSHOT_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_SNAPSHOT_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_SNAPSHOTS_PUBLIC_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_SNAPSHOTS_PUBLIC_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_STORAGE_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"RDS_STORAGE_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_AUDIT_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_AUDIT_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_BACKUP_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_BACKUP_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_CONFIGURATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_CONFIGURATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_KMS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_KMS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_MAINTENANCE_SETTINGS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_MAINTENANCE_SETTINGS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_PUBLIC_ACCESS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_PUBLIC_ACCESS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_DEFAULT_ADMIN_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_DEFAULT_ADMIN_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_DEFAULT_DB_NAME_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_DEFAULT_DB_NAME_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_ENHANCED_VPC_ROUTING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_ENHANCED_VPC_ROUTING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_REQUIRE_TLS_SSL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_REQUIRE_TLS_SSL",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REQUIRED_TAGS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"REQUIRED_TAGS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ROOT_ACCOUNT_HARDWARE_MFA_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ROOT_ACCOUNT_HARDWARE_MFA_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ROOT_ACCOUNT_MFA_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"ROOT_ACCOUNT_MFA_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_ACCOUNT_LEVEL_PUBLIC_ACCESS_BLOCKS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_ACCOUNT_LEVEL_PUBLIC_ACCESS_BLOCKS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_ACCOUNT_LEVEL_PUBLIC_ACCESS_BLOCKS_PERIODIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_ACCOUNT_LEVEL_PUBLIC_ACCESS_BLOCKS_PERIODIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_ACL_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_ACL_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_BLOCKED_ACTIONS_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_BLOCKED_ACTIONS_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_DEFAULT_LOCK_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_DEFAULT_LOCK_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_LEVEL_PUBLIC_ACCESS_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_LEVEL_PUBLIC_ACCESS_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_POLICY_GRANTEE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_POLICY_GRANTEE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_POLICY_NOT_MORE_PERMISSIVE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_POLICY_NOT_MORE_PERMISSIVE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_PUBLIC_READ_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_PUBLIC_READ_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_PUBLIC_WRITE_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_PUBLIC_WRITE_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_REPLICATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_REPLICATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_SERVER_SIDE_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_SERVER_SIDE_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_SSL_REQUESTS_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_SSL_REQUESTS_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_VERSIONING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_VERSIONING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_DEFAULT_ENCRYPTION_KMS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_DEFAULT_ENCRYPTION_KMS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_EVENT_NOTIFICATIONS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_EVENT_NOTIFICATIONS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_LIFECYCLE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_LIFECYCLE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_VERSION_LIFECYCLE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"S3_VERSION_LIFECYCLE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SAGEMAKER_ENDPOINT_CONFIGURATION_KMS_KEY_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SAGEMAKER_ENDPOINT_CONFIGURATION_KMS_KEY_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SAGEMAKER_NOTEBOOK_INSTANCE_KMS_KEY_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SAGEMAKER_NOTEBOOK_INSTANCE_KMS_KEY_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SAGEMAKER_NOTEBOOK_NO_DIRECT_INTERNET_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SAGEMAKER_NOTEBOOK_NO_DIRECT_INTERNET_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_ROTATION_ENABLED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_ROTATION_ENABLED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_SCHEDULED_ROTATION_SUCCESS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_SCHEDULED_ROTATION_SUCCESS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_SECRET_PERIODIC_ROTATION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_SECRET_PERIODIC_ROTATION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_SECRET_UNUSED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_SECRET_UNUSED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_USING_CMK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_USING_CMK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECURITYHUB_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SECURITYHUB_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SERVICE_VPC_ENDPOINT_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SERVICE_VPC_ENDPOINT_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SHIELD_ADVANCED_ENABLED_AUTO_RENEW() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SHIELD_ADVANCED_ENABLED_AUTO_RENEW",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SHIELD_DRT_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SHIELD_DRT_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SNS_ENCRYPTED_KMS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SNS_ENCRYPTED_KMS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SNS_TOPIC_MESSAGE_DELIVERY_NOTIFICATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SNS_TOPIC_MESSAGE_DELIVERY_NOTIFICATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SSM_DOCUMENT_NOT_PUBLIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SSM_DOCUMENT_NOT_PUBLIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_STORAGEGATEWAY_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"STORAGEGATEWAY_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SUBNET_AUTO_ASSIGN_PUBLIC_IP_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"SUBNET_AUTO_ASSIGN_PUBLIC_IP_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VIRTUALMACHINE_LAST_BACKUP_RECOVERY_POINT_CREATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VIRTUALMACHINE_LAST_BACKUP_RECOVERY_POINT_CREATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VIRTUALMACHINE_RESOURCES_PROTECTED_BY_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VIRTUALMACHINE_RESOURCES_PROTECTED_BY_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_DEFAULT_SECURITY_GROUP_CLOSED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VPC_DEFAULT_SECURITY_GROUP_CLOSED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_FLOW_LOGS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VPC_FLOW_LOGS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_NETWORK_ACL_UNUSED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VPC_NETWORK_ACL_UNUSED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_PEERING_DNS_RESOLUTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VPC_PEERING_DNS_RESOLUTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_SG_OPEN_ONLY_TO_AUTHORIZED_PORTS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VPC_SG_OPEN_ONLY_TO_AUTHORIZED_PORTS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_VPN_2_TUNNELS_UP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"VPC_VPN_2_TUNNELS_UP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_CLASSIC_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_CLASSIC_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_GLOBAL_RULE_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_GLOBAL_RULE_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_GLOBAL_RULEGROUP_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_GLOBAL_RULEGROUP_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_GLOBAL_WEBACL_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_GLOBAL_WEBACL_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_REGIONAL_RULE_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_REGIONAL_RULE_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_REGIONAL_RULEGROUP_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_REGIONAL_RULEGROUP_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_REGIONAL_WEBACL_NOT_EMPTY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAF_REGIONAL_WEBACL_NOT_EMPTY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAFV2_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_config.ManagedRuleIdentifiers",
		"WAFV2_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

