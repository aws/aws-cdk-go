package previewawspcsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterLogsMixin",
		reflect.TypeOf((*CfnClusterLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogs",
		reflect.TypeOf((*CfnClusterPcsJobcompLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsJobcompLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsDestProps",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsJobcompLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsJobcompLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterPcsJobcompLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterPcsJobcompLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterPcsJobcompLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterPcsJobcompLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsJobcompLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterPcsJobcompLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterPcsJobcompLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterPcsJobcompLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsRecordFields",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ID": CfnClusterPcsJobcompLogsRecordFields_RESOURCE_ID,
			"RESOURCE_TYPE": CfnClusterPcsJobcompLogsRecordFields_RESOURCE_TYPE,
			"EVENT_TIMESTAMP": CfnClusterPcsJobcompLogsRecordFields_EVENT_TIMESTAMP,
			"SCHEDULER_TYPE": CfnClusterPcsJobcompLogsRecordFields_SCHEDULER_TYPE,
			"SCHEDULER_MAJOR_VERSION": CfnClusterPcsJobcompLogsRecordFields_SCHEDULER_MAJOR_VERSION,
			"FIELDS": CfnClusterPcsJobcompLogsRecordFields_FIELDS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsJobcompLogsS3Props",
		reflect.TypeOf((*CfnClusterPcsJobcompLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogs",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsSchedulerLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsDestProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterPcsSchedulerLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsSchedulerLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterPcsSchedulerLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterPcsSchedulerLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterPcsSchedulerLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterPcsSchedulerLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterPcsSchedulerLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterPcsSchedulerLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterPcsSchedulerLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterPcsSchedulerLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsRecordFields",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ID": CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_ID,
			"RESOURCE_TYPE": CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_TYPE,
			"EVENT_TIMESTAMP": CfnClusterPcsSchedulerLogsRecordFields_EVENT_TIMESTAMP,
			"LOG_LEVEL": CfnClusterPcsSchedulerLogsRecordFields_LOG_LEVEL,
			"LOG_NAME": CfnClusterPcsSchedulerLogsRecordFields_LOG_NAME,
			"SCHEDULER_TYPE": CfnClusterPcsSchedulerLogsRecordFields_SCHEDULER_TYPE,
			"SCHEDULER_MAJOR_VERSION": CfnClusterPcsSchedulerLogsRecordFields_SCHEDULER_MAJOR_VERSION,
			"SCHEDULER_PATCH_VERSION": CfnClusterPcsSchedulerLogsRecordFields_SCHEDULER_PATCH_VERSION,
			"NODE_TYPE": CfnClusterPcsSchedulerLogsRecordFields_NODE_TYPE,
			"MESSAGE": CfnClusterPcsSchedulerLogsRecordFields_MESSAGE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsS3Props",
		reflect.TypeOf((*CfnClusterPcsSchedulerLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.AccountingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AccountingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.AuthKeyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AuthKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.EndpointProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.JwtAuthProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JwtAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.JwtKeyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JwtKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.NetworkingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_NetworkingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SchedulerProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SchedulerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SlurmRestProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupMixinProps",
		reflect.TypeOf((*CfnComputeNodeGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputeNodeGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.CustomLaunchTemplateProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_CustomLaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.InstanceConfigProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_InstanceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.ScalingConfigurationProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_ScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.SpotOptionsProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SpotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueueMixinProps",
		reflect.TypeOf((*CfnQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin",
		reflect.TypeOf((*CfnQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.ComputeNodeGroupConfigurationProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_ComputeNodeGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
}
