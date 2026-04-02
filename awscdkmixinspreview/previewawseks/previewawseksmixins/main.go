package previewawseksmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityAckLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityAckLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityAckLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityAckLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityAckLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityAckLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityAckLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityAckLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityAckLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityAckLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityAckLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityAckLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityAckLogsRecordFields_LEVEL,
			"MESSAGE": CfnCapabilityEksCapabilityAckLogsRecordFields_MESSAGE,
			"ERR": CfnCapabilityEksCapabilityAckLogsRecordFields_ERR,
			"ERROR": CfnCapabilityEksCapabilityAckLogsRecordFields_ERROR,
			"CONTROLLER": CfnCapabilityEksCapabilityAckLogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnCapabilityEksCapabilityAckLogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnCapabilityEksCapabilityAckLogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnCapabilityEksCapabilityAckLogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnCapabilityEksCapabilityAckLogsRecordFields_WORKER_COUNT,
			"KIND": CfnCapabilityEksCapabilityAckLogsRecordFields_KIND,
			"RECONCILER_KIND": CfnCapabilityEksCapabilityAckLogsRecordFields_RECONCILER_KIND,
			"NAME": CfnCapabilityEksCapabilityAckLogsRecordFields_NAME,
			"NAMESPACE": CfnCapabilityEksCapabilityAckLogsRecordFields_NAMESPACE,
			"AWS_SERVICE": CfnCapabilityEksCapabilityAckLogsRecordFields_AWS_SERVICE,
			"VERSION": CfnCapabilityEksCapabilityAckLogsRecordFields_VERSION,
			"TYPE": CfnCapabilityEksCapabilityAckLogsRecordFields_TYPE,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityAckLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityAckLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityAckLogsS3Props)(nil)).Elem(),
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
			"MESSAGE": CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_MESSAGE,
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
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_LEVEL,
			"MESSAGE": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_MESSAGE,
			"ERROR": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_ERROR,
			"CONTROLLER": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_CONTROLLER,
			"CONTROLLERGROUP": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_CONTROLLERGROUP,
			"CONTROLLERKIND": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_CONTROLLERKIND,
			"RECONCILEID": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_RECONCILEID,
			"WORKER_COUNT": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_WORKER_COUNT,
			"NAME": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_NAME,
			"NAMESPACE": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_NAMESPACE,
			"APPLICATION": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_APPLICATION,
			"APP_NAMESPACE": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_APP_NAMESPACE,
			"PROJECT": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_PROJECT,
			"APPLICATIONSET": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_APPLICATIONSET,
			"ACTION": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_ACTION,
			"COUNT": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_COUNT,
			"REQUEUEAFTER": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_REQUEUEAFTER,
			"TYPE": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_TYPE,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdApplicationsetLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdCommitserverLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_LEVEL,
			"MESSAGE": CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_MESSAGE,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdCommitserverLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdReposerverLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_LEVEL,
			"MESSAGE": CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_MESSAGE,
			"ERROR": CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_ERROR,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdReposerverLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogs",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsDestProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsLogGroupProps",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsRecordFields",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STREAM": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_STREAM,
			"LEVEL": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_LEVEL,
			"MESSAGE": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_MESSAGE,
			"ERROR": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_ERROR,
			"NAME": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_NAME,
			"NAMESPACE": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_NAMESPACE,
			"REASON": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_REASON,
			"RESOURCE_ARN": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsS3Props",
		reflect.TypeOf((*CfnCapabilityEksCapabilityArgocdServerLogsS3Props)(nil)).Elem(),
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
