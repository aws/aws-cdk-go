package previewawsentityresolutionmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowLogsMixin",
		reflect.TypeOf((*CfnIdMappingWorkflowLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdMappingWorkflowLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowMixinProps",
		reflect.TypeOf((*CfnIdMappingWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdMappingWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.IdMappingIncrementalRunConfigProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingIncrementalRunConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.IdMappingRuleBasedPropertiesProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingRuleBasedPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.IdMappingTechniquesProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingTechniquesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.IdMappingWorkflowInputSourceProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingWorkflowInputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.IdMappingWorkflowOutputSourceProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IdMappingWorkflowOutputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.IntermediateSourceConfigurationProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_IntermediateSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.ProviderPropertiesProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_ProviderPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnIdMappingWorkflowPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogs",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnIdMappingWorkflowWorkflowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsDestProps",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsFirehoseProps",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsLogGroupProps",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnIdMappingWorkflowWorkflowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnIdMappingWorkflowWorkflowLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnIdMappingWorkflowWorkflowLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnIdMappingWorkflowWorkflowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnIdMappingWorkflowWorkflowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnIdMappingWorkflowWorkflowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnIdMappingWorkflowWorkflowLogsOutputFormat_S3_JSON,
			"PLAIN": CfnIdMappingWorkflowWorkflowLogsOutputFormat_S3_PLAIN,
			"W3C": CfnIdMappingWorkflowWorkflowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnIdMappingWorkflowWorkflowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsRecordFields",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnIdMappingWorkflowWorkflowLogsRecordFields_RESOURCE_ARN,
			"EVENT_TYPE": CfnIdMappingWorkflowWorkflowLogsRecordFields_EVENT_TYPE,
			"EVENT_TIMESTAMP": CfnIdMappingWorkflowWorkflowLogsRecordFields_EVENT_TIMESTAMP,
			"JOB_ID": CfnIdMappingWorkflowWorkflowLogsRecordFields_JOB_ID,
			"WORKFLOW_NAME": CfnIdMappingWorkflowWorkflowLogsRecordFields_WORKFLOW_NAME,
			"WORKFLOW_START_TIME": CfnIdMappingWorkflowWorkflowLogsRecordFields_WORKFLOW_START_TIME,
			"WORKFLOW_COMPLETE_TIME": CfnIdMappingWorkflowWorkflowLogsRecordFields_WORKFLOW_COMPLETE_TIME,
			"DATA_PROCESSING_PROGRESSION": CfnIdMappingWorkflowWorkflowLogsRecordFields_DATA_PROCESSING_PROGRESSION,
			"ERROR_MESSAGE": CfnIdMappingWorkflowWorkflowLogsRecordFields_ERROR_MESSAGE,
			"MATCH_RULE": CfnIdMappingWorkflowWorkflowLogsRecordFields_MATCH_RULE,
			"MATCH_COUNT": CfnIdMappingWorkflowWorkflowLogsRecordFields_MATCH_COUNT,
			"TOTAL_RECORDS_PROCESSED": CfnIdMappingWorkflowWorkflowLogsRecordFields_TOTAL_RECORDS_PROCESSED,
			"TOTAL_RECORDS_UNPROCESSED": CfnIdMappingWorkflowWorkflowLogsRecordFields_TOTAL_RECORDS_UNPROCESSED,
			"INCREMENTAL_RECORDS_PROCESSED": CfnIdMappingWorkflowWorkflowLogsRecordFields_INCREMENTAL_RECORDS_PROCESSED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsS3Props",
		reflect.TypeOf((*CfnIdMappingWorkflowWorkflowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespaceMixinProps",
		reflect.TypeOf((*CfnIdNamespaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin",
		reflect.TypeOf((*CfnIdNamespacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdNamespacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin.IdNamespaceIdMappingWorkflowPropertiesProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_IdNamespaceIdMappingWorkflowPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin.IdNamespaceInputSourceProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_IdNamespaceInputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin.NamespaceProviderPropertiesProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_NamespaceProviderPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin.NamespaceRuleBasedPropertiesProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_NamespaceRuleBasedPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnIdNamespacePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowLogsMixin",
		reflect.TypeOf((*CfnMatchingWorkflowLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchingWorkflowLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowMixinProps",
		reflect.TypeOf((*CfnMatchingWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMatchingWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.CustomerProfilesIntegrationConfigProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_CustomerProfilesIntegrationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.IncrementalRunConfigProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_IncrementalRunConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.InputSourceProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_InputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.IntermediateSourceConfigurationProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_IntermediateSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.OutputAttributeProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_OutputAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.OutputSourceProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_OutputSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.ProviderPropertiesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_ProviderPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.ResolutionTechniquesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_ResolutionTechniquesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.RuleBasedPropertiesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleBasedPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.RuleConditionPropertiesProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleConditionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.RuleConditionProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnMatchingWorkflowPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogs",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMatchingWorkflowWorkflowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsDestProps",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsFirehoseProps",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsLogGroupProps",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMatchingWorkflowWorkflowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMatchingWorkflowWorkflowLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMatchingWorkflowWorkflowLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMatchingWorkflowWorkflowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMatchingWorkflowWorkflowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMatchingWorkflowWorkflowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMatchingWorkflowWorkflowLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMatchingWorkflowWorkflowLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMatchingWorkflowWorkflowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMatchingWorkflowWorkflowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsRecordFields",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnMatchingWorkflowWorkflowLogsRecordFields_RESOURCE_ARN,
			"EVENT_TYPE": CfnMatchingWorkflowWorkflowLogsRecordFields_EVENT_TYPE,
			"EVENT_TIMESTAMP": CfnMatchingWorkflowWorkflowLogsRecordFields_EVENT_TIMESTAMP,
			"JOB_ID": CfnMatchingWorkflowWorkflowLogsRecordFields_JOB_ID,
			"WORKFLOW_NAME": CfnMatchingWorkflowWorkflowLogsRecordFields_WORKFLOW_NAME,
			"WORKFLOW_START_TIME": CfnMatchingWorkflowWorkflowLogsRecordFields_WORKFLOW_START_TIME,
			"WORKFLOW_COMPLETE_TIME": CfnMatchingWorkflowWorkflowLogsRecordFields_WORKFLOW_COMPLETE_TIME,
			"DATA_PROCESSING_PROGRESSION": CfnMatchingWorkflowWorkflowLogsRecordFields_DATA_PROCESSING_PROGRESSION,
			"ERROR_MESSAGE": CfnMatchingWorkflowWorkflowLogsRecordFields_ERROR_MESSAGE,
			"MATCH_RULE": CfnMatchingWorkflowWorkflowLogsRecordFields_MATCH_RULE,
			"MATCH_COUNT": CfnMatchingWorkflowWorkflowLogsRecordFields_MATCH_COUNT,
			"TOTAL_RECORDS_PROCESSED": CfnMatchingWorkflowWorkflowLogsRecordFields_TOTAL_RECORDS_PROCESSED,
			"TOTAL_RECORDS_UNPROCESSED": CfnMatchingWorkflowWorkflowLogsRecordFields_TOTAL_RECORDS_UNPROCESSED,
			"INCREMENTAL_RECORDS_PROCESSED": CfnMatchingWorkflowWorkflowLogsRecordFields_INCREMENTAL_RECORDS_PROCESSED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsS3Props",
		reflect.TypeOf((*CfnMatchingWorkflowWorkflowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnPolicyStatementMixinProps",
		reflect.TypeOf((*CfnPolicyStatementMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnPolicyStatementPropsMixin",
		reflect.TypeOf((*CfnPolicyStatementPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyStatementPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingMixinProps",
		reflect.TypeOf((*CfnSchemaMappingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingPropsMixin",
		reflect.TypeOf((*CfnSchemaMappingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchemaMappingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingPropsMixin.SchemaInputAttributeProperty",
		reflect.TypeOf((*CfnSchemaMappingPropsMixin_SchemaInputAttributeProperty)(nil)).Elem(),
	)
}
