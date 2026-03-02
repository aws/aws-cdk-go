package previewawsnetworkfirewallmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogs",
		reflect.TypeOf((*CfnFirewallAlertLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallAlertLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallAlertLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallAlertLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallAlertLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallAlertLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFirewallAlertLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFirewallAlertLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallAlertLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallAlertLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallAlertLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFirewallAlertLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFirewallAlertLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallAlertLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsS3Props",
		reflect.TypeOf((*CfnFirewallAlertLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogs",
		reflect.TypeOf((*CfnFirewallFlowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallFlowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallFlowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallFlowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallFlowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallFlowLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFirewallFlowLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFirewallFlowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallFlowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallFlowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallFlowLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFirewallFlowLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFirewallFlowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallFlowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsS3Props",
		reflect.TypeOf((*CfnFirewallFlowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		reflect.TypeOf((*CfnFirewallLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFirewallLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallMixinProps",
		reflect.TypeOf((*CfnFirewallMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyMixinProps",
		reflect.TypeOf((*CfnFirewallPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFirewallPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.ActionDefinitionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_ActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.CustomActionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_CustomActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.FirewallPolicyProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_FirewallPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.FlowTimeoutsProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_FlowTimeoutsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.IPSetProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_IPSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.PolicyVariablesProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_PolicyVariablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.PublishMetricActionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_PublishMetricActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.StatefulEngineOptionsProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatefulEngineOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.StatefulRuleGroupOverrideProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatefulRuleGroupOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.StatefulRuleGroupReferenceProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatefulRuleGroupReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPolicyPropsMixin.StatelessRuleGroupReferenceProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatelessRuleGroupReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin",
		reflect.TypeOf((*CfnFirewallPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFirewallPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin.AvailabilityZoneMappingProperty",
		reflect.TypeOf((*CfnFirewallPropsMixin_AvailabilityZoneMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallPropsMixin.SubnetMappingProperty",
		reflect.TypeOf((*CfnFirewallPropsMixin_SubnetMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogs",
		reflect.TypeOf((*CfnFirewallTlsLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallTlsLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallTlsLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallTlsLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallTlsLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallTlsLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFirewallTlsLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFirewallTlsLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallTlsLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallTlsLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallTlsLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFirewallTlsLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFirewallTlsLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallTlsLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsS3Props",
		reflect.TypeOf((*CfnFirewallTlsLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationMixinProps",
		reflect.TypeOf((*CfnLoggingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoggingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin.LogDestinationConfigProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_LogDestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnLoggingConfigurationPropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupMixinProps",
		reflect.TypeOf((*CfnRuleGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin",
		reflect.TypeOf((*CfnRuleGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuleGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.ActionDefinitionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_ActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.AddressProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.CustomActionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_CustomActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.HeaderProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_HeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.IPSetProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_IPSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.IPSetReferenceProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_IPSetReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.MatchAttributesProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_MatchAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.PortRangeProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.PortSetProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_PortSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.PublishMetricActionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_PublishMetricActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.ReferenceSetsProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_ReferenceSetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.RuleDefinitionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.RuleGroupProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.RuleOptionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.RuleVariablesProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleVariablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.RulesSourceListProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RulesSourceListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.RulesSourceProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RulesSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.StatefulRuleOptionsProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatefulRuleOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.StatefulRuleProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatefulRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.StatelessRuleProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatelessRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.StatelessRulesAndCustomActionsProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatelessRulesAndCustomActionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.SummaryConfigurationProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_SummaryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnRuleGroupPropsMixin.TCPFlagFieldProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_TCPFlagFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationMixinProps",
		reflect.TypeOf((*CfnTLSInspectionConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTLSInspectionConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.AddressProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.CheckCertificateRevocationStatusProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_CheckCertificateRevocationStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.PortRangeProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.ServerCertificateConfigurationProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_ServerCertificateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.ServerCertificateProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_ServerCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.ServerCertificateScopeProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_ServerCertificateScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin.TLSInspectionConfigurationProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_TLSInspectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnVpcEndpointAssociationMixinProps",
		reflect.TypeOf((*CfnVpcEndpointAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnVpcEndpointAssociationPropsMixin",
		reflect.TypeOf((*CfnVpcEndpointAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcEndpointAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnVpcEndpointAssociationPropsMixin.SubnetMappingProperty",
		reflect.TypeOf((*CfnVpcEndpointAssociationPropsMixin_SubnetMappingProperty)(nil)).Elem(),
	)
}
