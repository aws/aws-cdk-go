package previewawsdevopsagentmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogs",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAgentSpaceApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsDestProps",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAgentSpaceApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentSpaceApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAgentSpaceApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAgentSpaceApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentSpaceApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAgentSpaceApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentSpaceApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAgentSpaceApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAgentSpaceApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAgentSpaceApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsRecordFields",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL_ACCOUNT_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ACCOUNT_ID,
			"OPTIONAL_AGENT_SPACE_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_AGENT_SPACE_ID,
			"OPTIONAL_LEVEL": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_LEVEL,
			"OPTIONAL_ASSOCIATION_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ASSOCIATION_ID,
			"OPTIONAL_STATUS": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_STATUS,
			"OPTIONAL_WEBHOOK_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_WEBHOOK_ID,
			"OPTIONAL_MCP_ENDPOINT_URL": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_MCP_ENDPOINT_URL,
			"OPTIONAL_SERVICE_TYPE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_SERVICE_TYPE,
			"OPTIONAL_SERVICE_ENDPOINT_URL": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_SERVICE_ENDPOINT_URL,
			"OPTIONAL_SERVICE_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_SERVICE_ID,
			"OPTIONAL_REQUEST_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_REQUEST_ID,
			"OPTIONAL_OPERATION": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_OPERATION,
			"OPTIONAL_TASK_TYPE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_TASK_TYPE,
			"OPTIONAL_TASK_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_TASK_ID,
			"OPTIONAL_REFERENCE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_REFERENCE,
			"OPTIONAL_ERROR_TYPE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ERROR_TYPE,
			"OPTIONAL_ERROR_MESSAGE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ERROR_MESSAGE,
			"OPTIONAL_DETAILS": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_DETAILS,
			"RESOURCE_ARN": CfnAgentSpaceApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnAgentSpaceApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsS3Props",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceLogsMixin",
		reflect.TypeOf((*CfnAgentSpaceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentSpaceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceMixinProps",
		reflect.TypeOf((*CfnAgentSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin",
		reflect.TypeOf((*CfnAgentSpacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentSpacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin.IamAuthConfigurationProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_IamAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin.IdcAuthConfigurationProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_IdcAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin.OperatorAppProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_OperatorAppProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationMixinProps",
		reflect.TypeOf((*CfnAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin",
		reflect.TypeOf((*CfnAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.AWSConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AWSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.AWSResourceProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AWSResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.DynatraceConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_DynatraceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.EventChannelConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_EventChannelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.GitHubConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_GitHubConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.GitLabConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_GitLabConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerDatadogConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerDatadogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerNewRelicConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerNewRelicConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerSplunkConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerSplunkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.ServiceConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_ServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.ServiceNowConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_ServiceNowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SlackChannelProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SlackConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SlackTransmissionTargetProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackTransmissionTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SourceAwsConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SourceAwsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.AdditionalServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AdditionalServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.ApiKeyDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ApiKeyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.BearerTokenDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_BearerTokenDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.DynatraceAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DynatraceAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.DynatraceServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DynatraceServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.GitLabDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_GitLabDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.MCPServerAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.MCPServerDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.MCPServerOAuthClientCredentialsConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerOAuthClientCredentialsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.MCPServerSplunkAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerSplunkAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.MCPServerSplunkDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerSplunkDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.NewRelicApiKeyConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NewRelicApiKeyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.NewRelicAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NewRelicAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.NewRelicServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NewRelicServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.OAuthClientDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_OAuthClientDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.RegisteredDynatraceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredDynatraceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.RegisteredGitLabServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredGitLabServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.RegisteredMCPServerDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredMCPServerDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.RegisteredNewRelicDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredNewRelicDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.RegisteredServiceNowDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredServiceNowDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.ServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.ServiceNowAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceNowAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnServicePropsMixin.ServiceNowServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceNowServiceDetailsProperty)(nil)).Elem(),
	)
}
