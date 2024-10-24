// The CDK Construct Library for AWS::IoT
package awscdkiotalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.AccountAuditConfiguration",
		reflect.TypeOf((*AccountAuditConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AccountAuditConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAccountAuditConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.AccountAuditConfigurationProps",
		reflect.TypeOf((*AccountAuditConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.ActionConfig",
		reflect.TypeOf((*ActionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-iot-alpha.AuditCheck",
		reflect.TypeOf((*AuditCheck)(nil)).Elem(),
		map[string]interface{}{
			"AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK": AuditCheck_AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK,
			"CA_CERTIFICATE_EXPIRING_CHECK": AuditCheck_CA_CERTIFICATE_EXPIRING_CHECK,
			"CA_CERTIFICATE_KEY_QUALITY_CHECK": AuditCheck_CA_CERTIFICATE_KEY_QUALITY_CHECK,
			"CONFLICTING_CLIENT_IDS_CHECK": AuditCheck_CONFLICTING_CLIENT_IDS_CHECK,
			"DEVICE_CERTIFICATE_EXPIRING_CHECK": AuditCheck_DEVICE_CERTIFICATE_EXPIRING_CHECK,
			"DEVICE_CERTIFICATE_KEY_QUALITY_CHECK": AuditCheck_DEVICE_CERTIFICATE_KEY_QUALITY_CHECK,
			"DEVICE_CERTIFICATE_SHARED_CHECK": AuditCheck_DEVICE_CERTIFICATE_SHARED_CHECK,
			"IOT_POLICY_OVERLY_PERMISSIVE_CHECK": AuditCheck_IOT_POLICY_OVERLY_PERMISSIVE_CHECK,
			"IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK": AuditCheck_IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK,
			"IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK": AuditCheck_IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK,
			"LOGGING_DISABLED_CHECK": AuditCheck_LOGGING_DISABLED_CHECK,
			"REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK": AuditCheck_REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK,
			"REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK": AuditCheck_REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK,
			"UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK": AuditCheck_UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.CheckConfiguration",
		reflect.TypeOf((*CheckConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.DayOfMonth",
		reflect.TypeOf((*DayOfMonth)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "day", GoGetter: "Day"},
		},
		func() interface{} {
			return &jsiiProxy_DayOfMonth{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-iot-alpha.DayOfWeek",
		reflect.TypeOf((*DayOfWeek)(nil)).Elem(),
		map[string]interface{}{
			"SUNDAY": DayOfWeek_SUNDAY,
			"MONDAY": DayOfWeek_MONDAY,
			"TUESDAY": DayOfWeek_TUESDAY,
			"WEDNESDAY": DayOfWeek_WEDNESDAY,
			"THURSDAY": DayOfWeek_THURSDAY,
			"FRIDAY": DayOfWeek_FRIDAY,
			"SATURDAY": DayOfWeek_SATURDAY,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-iot-alpha.Frequency",
		reflect.TypeOf((*Frequency)(nil)).Elem(),
		map[string]interface{}{
			"DAILY": Frequency_DAILY,
			"WEEKLY": Frequency_WEEKLY,
			"BI_WEEKLY": Frequency_BI_WEEKLY,
			"MONTHLY": Frequency_MONTHLY,
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.IAccountAuditConfiguration",
		reflect.TypeOf((*IAccountAuditConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAccountAuditConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.IAction",
		reflect.TypeOf((*IAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IAction{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.ILogging",
		reflect.TypeOf((*ILogging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "logId", GoGetter: "LogId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ILogging{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.IScheduledAudit",
		reflect.TypeOf((*IScheduledAudit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledAuditArn", GoGetter: "ScheduledAuditArn"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledAuditName", GoGetter: "ScheduledAuditName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IScheduledAudit{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-iot-alpha.ITopicRule",
		reflect.TypeOf((*ITopicRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleArn", GoGetter: "TopicRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleName", GoGetter: "TopicRuleName"},
		},
		func() interface{} {
			j := jsiiProxy_ITopicRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.IotSql",
		reflect.TypeOf((*IotSql)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IotSql{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.IotSqlConfig",
		reflect.TypeOf((*IotSqlConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-iot-alpha.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"ERROR": LogLevel_ERROR,
			"WARN": LogLevel_WARN,
			"INFO": LogLevel_INFO,
			"DEBUG": LogLevel_DEBUG,
			"DISABLED": LogLevel_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.Logging",
		reflect.TypeOf((*Logging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logId", GoGetter: "LogId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Logging{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILogging)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.LoggingProps",
		reflect.TypeOf((*LoggingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.ScheduledAudit",
		reflect.TypeOf((*ScheduledAudit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledAuditArn", GoGetter: "ScheduledAuditArn"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledAuditName", GoGetter: "ScheduledAuditName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ScheduledAudit{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScheduledAudit)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.ScheduledAuditAttributes",
		reflect.TypeOf((*ScheduledAuditAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.ScheduledAuditProps",
		reflect.TypeOf((*ScheduledAuditProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		reflect.TypeOf((*TopicRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleArn", GoGetter: "TopicRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicRuleName", GoGetter: "TopicRuleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TopicRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITopicRule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-alpha.TopicRuleProps",
		reflect.TypeOf((*TopicRuleProps)(nil)).Elem(),
	)
}
