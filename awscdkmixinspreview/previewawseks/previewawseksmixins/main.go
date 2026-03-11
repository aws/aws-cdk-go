package previewawseksmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3Logs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3Logs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityAckS3Logs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityAckS3LogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityAckS3LogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityAckS3LogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityAckS3LogsRecordFields_LEVEL,
			"MSG": CfnCapabilityEksCapabilityAckS3LogsRecordFields_MSG,
			"ERR": CfnCapabilityEksCapabilityAckS3LogsRecordFields_ERR,
			"ERROR": CfnCapabilityEksCapabilityAckS3LogsRecordFields_ERROR,
			"CONTROLLER": CfnCapabilityEksCapabilityAckS3LogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnCapabilityEksCapabilityAckS3LogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnCapabilityEksCapabilityAckS3LogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnCapabilityEksCapabilityAckS3LogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnCapabilityEksCapabilityAckS3LogsRecordFields_WORKER_COUNT,
			"KIND": CfnCapabilityEksCapabilityAckS3LogsRecordFields_KIND,
			"RECONCILER_KIND": CfnCapabilityEksCapabilityAckS3LogsRecordFields_RECONCILER_KIND,
			"NAME": CfnCapabilityEksCapabilityAckS3LogsRecordFields_NAME,
			"NAMESPACE": CfnCapabilityEksCapabilityAckS3LogsRecordFields_NAMESPACE,
			"AWS_SERVICE": CfnCapabilityEksCapabilityAckS3LogsRecordFields_AWS_SERVICE,
			"VERSION": CfnCapabilityEksCapabilityAckS3LogsRecordFields_VERSION,
			"TYPE": CfnCapabilityEksCapabilityAckS3LogsRecordFields_TYPE,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityAckS3LogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityAckS3LogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckS3LogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_LEVEL,
			"MSG": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_MSG,
			"ERROR": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_ERROR,
			"APPLICATION": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_APPLICATION,
			"APP_NAMESPACE": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_APP_NAMESPACE,
			"PROJECT": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_PROJECT,
			"DEST_SERVER": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_DEST_SERVER,
			"DEST_NAMESPACE": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_DEST_NAMESPACE,
			"DEST_NAME": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_DEST_NAME,
			"COMPARISON_LEVEL": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_COMPARISON_LEVEL,
			"TIME_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_TIME_MS,
			"PATCH_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_PATCH_MS,
			"SETOP_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_SETOP_MS,
			"COMPARE_APP_STATE_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_COMPARE_APP_STATE_MS,
			"REFRESH_APP_CONDITIONS_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_REFRESH_APP_CONDITIONS_MS,
			"PROCESS_REFRESH_APP_CONDITIONS_ERRORS_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_PROCESS_REFRESH_APP_CONDITIONS_ERRORS_MS,
			"COMPARISON_WITH_NOTHING_MS": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_COMPARISON_WITH_NOTHING_MS,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityKroLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityKroLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityKroLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityKroLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityKroLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityKroLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityKroLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityKroLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityKroLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityKroLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityKroLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityKroLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityKroLogsRecordFields_LEVEL,
			"MESSAGE": CfnCapabilityEksCapabilityKroLogsRecordFields_MESSAGE,
			"ERR": CfnCapabilityEksCapabilityKroLogsRecordFields_ERR,
			"ERROR": CfnCapabilityEksCapabilityKroLogsRecordFields_ERROR,
			"CONTROLLER": CfnCapabilityEksCapabilityKroLogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnCapabilityEksCapabilityKroLogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnCapabilityEksCapabilityKroLogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnCapabilityEksCapabilityKroLogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnCapabilityEksCapabilityKroLogsRecordFields_WORKER_COUNT,
			"GVR": CfnCapabilityEksCapabilityKroLogsRecordFields_GVR,
			"CRD": CfnCapabilityEksCapabilityKroLogsRecordFields_CRD,
			"NAME": CfnCapabilityEksCapabilityKroLogsRecordFields_NAME,
			"NAMESPACE": CfnCapabilityEksCapabilityKroLogsRecordFields_NAMESPACE,
			"ITEM": CfnCapabilityEksCapabilityKroLogsRecordFields_ITEM,
			"RESOURCEGRAPHDEFINITION": CfnCapabilityEksCapabilityKroLogsRecordFields_RESOURCEGRAPHDEFINITION,
			"TYPE": CfnCapabilityEksCapabilityKroLogsRecordFields_TYPE,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityKroLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityKroLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityKroLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityLogsMixin",
		reflect.TypeOf((*CfnCapabilityLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapabilityLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogs",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeBlockStorageLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsDestProps",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeBlockStorageLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeBlockStorageLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterAutoModeBlockStorageLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterAutoModeBlockStorageLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterAutoModeBlockStorageLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterAutoModeBlockStorageLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeBlockStorageLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterAutoModeBlockStorageLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterAutoModeBlockStorageLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterAutoModeBlockStorageLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsRecordFields",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"LEVEL": CfnClusterAutoModeBlockStorageLogsRecordFields_LEVEL,
			"STREAM": CfnClusterAutoModeBlockStorageLogsRecordFields_STREAM,
			"MESSAGE": CfnClusterAutoModeBlockStorageLogsRecordFields_MESSAGE,
			"ERR": CfnClusterAutoModeBlockStorageLogsRecordFields_ERR,
			"ERROR": CfnClusterAutoModeBlockStorageLogsRecordFields_ERROR,
			"ENABLED": CfnClusterAutoModeBlockStorageLogsRecordFields_ENABLED,
			"FEATURE": CfnClusterAutoModeBlockStorageLogsRecordFields_FEATURE,
			"RESOURCE_ARN": CfnClusterAutoModeBlockStorageLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnClusterAutoModeBlockStorageLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsS3Props",
		reflect.TypeOf((*CfnClusterAutoModeBlockStorageLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogs",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeComputeLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsDestProps",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsOutputFormat",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeComputeLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeComputeLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterAutoModeComputeLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterAutoModeComputeLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterAutoModeComputeLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterAutoModeComputeLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeComputeLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterAutoModeComputeLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterAutoModeComputeLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterAutoModeComputeLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsRecordFields",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"LEVEL": CfnClusterAutoModeComputeLogsRecordFields_LEVEL,
			"STREAM": CfnClusterAutoModeComputeLogsRecordFields_STREAM,
			"MESSAGE": CfnClusterAutoModeComputeLogsRecordFields_MESSAGE,
			"ERR": CfnClusterAutoModeComputeLogsRecordFields_ERR,
			"ERROR": CfnClusterAutoModeComputeLogsRecordFields_ERROR,
			"CONTROLLER": CfnClusterAutoModeComputeLogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnClusterAutoModeComputeLogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnClusterAutoModeComputeLogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnClusterAutoModeComputeLogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnClusterAutoModeComputeLogsRecordFields_WORKER_COUNT,
			"COUNT": CfnClusterAutoModeComputeLogsRecordFields_COUNT,
			"INSTANCE_TYPE_COUNT": CfnClusterAutoModeComputeLogsRecordFields_INSTANCE_TYPE_COUNT,
			"MODE": CfnClusterAutoModeComputeLogsRecordFields_MODE,
			"OFFERING_COUNT": CfnClusterAutoModeComputeLogsRecordFields_OFFERING_COUNT,
			"RESOURCE_ARN": CfnClusterAutoModeComputeLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnClusterAutoModeComputeLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsS3Props",
		reflect.TypeOf((*CfnClusterAutoModeComputeLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogs",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeIpamLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsDestProps",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsOutputFormat",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeIpamLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeIpamLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterAutoModeIpamLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterAutoModeIpamLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterAutoModeIpamLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterAutoModeIpamLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeIpamLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterAutoModeIpamLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterAutoModeIpamLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterAutoModeIpamLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsRecordFields",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"LEVEL": CfnClusterAutoModeIpamLogsRecordFields_LEVEL,
			"STREAM": CfnClusterAutoModeIpamLogsRecordFields_STREAM,
			"MESSAGE": CfnClusterAutoModeIpamLogsRecordFields_MESSAGE,
			"ERR": CfnClusterAutoModeIpamLogsRecordFields_ERR,
			"ERROR": CfnClusterAutoModeIpamLogsRecordFields_ERROR,
			"CONTROLLER": CfnClusterAutoModeIpamLogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnClusterAutoModeIpamLogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnClusterAutoModeIpamLogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnClusterAutoModeIpamLogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnClusterAutoModeIpamLogsRecordFields_WORKER_COUNT,
			"RESOURCE_ARN": CfnClusterAutoModeIpamLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnClusterAutoModeIpamLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsS3Props",
		reflect.TypeOf((*CfnClusterAutoModeIpamLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogs",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeLoadBalancingLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsDestProps",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsFirehoseProps",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsLogGroupProps",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterAutoModeLoadBalancingLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeLoadBalancingLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnClusterAutoModeLoadBalancingLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnClusterAutoModeLoadBalancingLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterAutoModeLoadBalancingLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnClusterAutoModeLoadBalancingLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnClusterAutoModeLoadBalancingLogsOutputFormat_S3_JSON,
			"PLAIN": CfnClusterAutoModeLoadBalancingLogsOutputFormat_S3_PLAIN,
			"W3C": CfnClusterAutoModeLoadBalancingLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterAutoModeLoadBalancingLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsRecordFields",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"LEVEL": CfnClusterAutoModeLoadBalancingLogsRecordFields_LEVEL,
			"STREAM": CfnClusterAutoModeLoadBalancingLogsRecordFields_STREAM,
			"MESSAGE": CfnClusterAutoModeLoadBalancingLogsRecordFields_MESSAGE,
			"ERR": CfnClusterAutoModeLoadBalancingLogsRecordFields_ERR,
			"ERROR": CfnClusterAutoModeLoadBalancingLogsRecordFields_ERROR,
			"CONTROLLER": CfnClusterAutoModeLoadBalancingLogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnClusterAutoModeLoadBalancingLogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnClusterAutoModeLoadBalancingLogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnClusterAutoModeLoadBalancingLogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnClusterAutoModeLoadBalancingLogsRecordFields_WORKER_COUNT,
			"ATTEMPT": CfnClusterAutoModeLoadBalancingLogsRecordFields_ATTEMPT,
			"DELAY": CfnClusterAutoModeLoadBalancingLogsRecordFields_DELAY,
			"TOTALITEMS": CfnClusterAutoModeLoadBalancingLogsRecordFields_TOTALITEMS,
			"RESOURCE_ARN": CfnClusterAutoModeLoadBalancingLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnClusterAutoModeLoadBalancingLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsS3Props",
		reflect.TypeOf((*CfnClusterAutoModeLoadBalancingLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterLogsMixin",
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
}
