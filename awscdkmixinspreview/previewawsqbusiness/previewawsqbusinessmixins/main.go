package previewawsqbusinessmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogs",
		reflect.TypeOf((*CfnApplicationEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationEventLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsDestProps",
		reflect.TypeOf((*CfnApplicationEventLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsFirehoseProps",
		reflect.TypeOf((*CfnApplicationEventLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsLogGroupProps",
		reflect.TypeOf((*CfnApplicationEventLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsOutputFormat",
		reflect.TypeOf((*CfnApplicationEventLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnApplicationEventLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnApplicationEventLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationEventLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnApplicationEventLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnApplicationEventLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnApplicationEventLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnApplicationEventLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnApplicationEventLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsOutputFormat.S3",
		reflect.TypeOf((*CfnApplicationEventLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationEventLogsOutputFormat_S3_JSON,
			"PLAIN": CfnApplicationEventLogsOutputFormat_S3_PLAIN,
			"W3C": CfnApplicationEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnApplicationEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsRecordFields",
		reflect.TypeOf((*CfnApplicationEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"APPLICATION_ID": CfnApplicationEventLogsRecordFields_APPLICATION_ID,
			"EVENT_TIMESTAMP": CfnApplicationEventLogsRecordFields_EVENT_TIMESTAMP,
			"LOG_TYPE": CfnApplicationEventLogsRecordFields_LOG_TYPE,
			"ACCOUNT_ID": CfnApplicationEventLogsRecordFields_ACCOUNT_ID,
			"CONVERSATION_ID": CfnApplicationEventLogsRecordFields_CONVERSATION_ID,
			"SYSTEM_MESSAGE_ID": CfnApplicationEventLogsRecordFields_SYSTEM_MESSAGE_ID,
			"USER_MESSAGE_ID": CfnApplicationEventLogsRecordFields_USER_MESSAGE_ID,
			"USER_MESSAGE": CfnApplicationEventLogsRecordFields_USER_MESSAGE,
			"SYSTEM_MESSAGE": CfnApplicationEventLogsRecordFields_SYSTEM_MESSAGE,
			"OUTPUT_METRICS_IS_MESSAGE_BLOCKED": CfnApplicationEventLogsRecordFields_OUTPUT_METRICS_IS_MESSAGE_BLOCKED,
			"OUTPUT_METRICS_IS_MESSAGE_WITH_NO_ANSWER": CfnApplicationEventLogsRecordFields_OUTPUT_METRICS_IS_MESSAGE_WITH_NO_ANSWER,
			"COMMENT": CfnApplicationEventLogsRecordFields_COMMENT,
			"USEFULNESS_REASON": CfnApplicationEventLogsRecordFields_USEFULNESS_REASON,
			"USEFULNESS": CfnApplicationEventLogsRecordFields_USEFULNESS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsS3Props",
		reflect.TypeOf((*CfnApplicationEventLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationLogsMixin",
		reflect.TypeOf((*CfnApplicationLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogs",
		reflect.TypeOf((*CfnApplicationSyncJobLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnApplicationSyncJobLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsDestProps",
		reflect.TypeOf((*CfnApplicationSyncJobLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsFirehoseProps",
		reflect.TypeOf((*CfnApplicationSyncJobLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsLogGroupProps",
		reflect.TypeOf((*CfnApplicationSyncJobLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsOutputFormat",
		reflect.TypeOf((*CfnApplicationSyncJobLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnApplicationSyncJobLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnApplicationSyncJobLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationSyncJobLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnApplicationSyncJobLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnApplicationSyncJobLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnApplicationSyncJobLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnApplicationSyncJobLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnApplicationSyncJobLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsOutputFormat.S3",
		reflect.TypeOf((*CfnApplicationSyncJobLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnApplicationSyncJobLogsOutputFormat_S3_JSON,
			"PLAIN": CfnApplicationSyncJobLogsOutputFormat_S3_PLAIN,
			"W3C": CfnApplicationSyncJobLogsOutputFormat_S3_W3C,
			"PARQUET": CfnApplicationSyncJobLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsRecordFields",
		reflect.TypeOf((*CfnApplicationSyncJobLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"AWSACCOUNTID": CfnApplicationSyncJobLogsRecordFields_AWSACCOUNTID,
			"DATASOURCEID": CfnApplicationSyncJobLogsRecordFields_DATASOURCEID,
			"SYNCJOBID": CfnApplicationSyncJobLogsRecordFields_SYNCJOBID,
			"INDEXID": CfnApplicationSyncJobLogsRecordFields_INDEXID,
			"MEMBERNAME": CfnApplicationSyncJobLogsRecordFields_MEMBERNAME,
			"MEMBERGLOBALNAME": CfnApplicationSyncJobLogsRecordFields_MEMBERGLOBALNAME,
			"DOCUMENTID": CfnApplicationSyncJobLogsRecordFields_DOCUMENTID,
			"DOCUMENTTITLE": CfnApplicationSyncJobLogsRecordFields_DOCUMENTTITLE,
			"SOURCEURI": CfnApplicationSyncJobLogsRecordFields_SOURCEURI,
			"ACL": CfnApplicationSyncJobLogsRecordFields_ACL,
			"METADATA": CfnApplicationSyncJobLogsRecordFields_METADATA,
			"TYPE": CfnApplicationSyncJobLogsRecordFields_TYPE,
			"CRAWLACTION": CfnApplicationSyncJobLogsRecordFields_CRAWLACTION,
			"CRAWLSTATUS": CfnApplicationSyncJobLogsRecordFields_CRAWLSTATUS,
			"SYNCSTATUS": CfnApplicationSyncJobLogsRecordFields_SYNCSTATUS,
			"INDEXSTATUS": CfnApplicationSyncJobLogsRecordFields_INDEXSTATUS,
			"CONNECTORDOCUMENTSTATUS": CfnApplicationSyncJobLogsRecordFields_CONNECTORDOCUMENTSTATUS,
			"MEMBERTYPE": CfnApplicationSyncJobLogsRecordFields_MEMBERTYPE,
			"ISMEMBERFEDERATED": CfnApplicationSyncJobLogsRecordFields_ISMEMBERFEDERATED,
			"MESSAGE": CfnApplicationSyncJobLogsRecordFields_MESSAGE,
			"JSONMESSAGE": CfnApplicationSyncJobLogsRecordFields_JSONMESSAGE,
			"GROUPNAME": CfnApplicationSyncJobLogsRecordFields_GROUPNAME,
			"GROUPGLOBALNAME": CfnApplicationSyncJobLogsRecordFields_GROUPGLOBALNAME,
			"ISGROUPFEDERATED": CfnApplicationSyncJobLogsRecordFields_ISGROUPFEDERATED,
			"HASHEDDOCUMENTID": CfnApplicationSyncJobLogsRecordFields_HASHEDDOCUMENTID,
			"CONTENTTYPE": CfnApplicationSyncJobLogsRecordFields_CONTENTTYPE,
			"INTERMEDIATESTATUS": CfnApplicationSyncJobLogsRecordFields_INTERMEDIATESTATUS,
			"ERRORMSG": CfnApplicationSyncJobLogsRecordFields_ERRORMSG,
			"APPLICATIONID": CfnApplicationSyncJobLogsRecordFields_APPLICATIONID,
			"EVENT_TIMESTAMP": CfnApplicationSyncJobLogsRecordFields_EVENT_TIMESTAMP,
			"FEATURENAME": CfnApplicationSyncJobLogsRecordFields_FEATURENAME,
			"LOGLEVEL": CfnApplicationSyncJobLogsRecordFields_LOGLEVEL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsS3Props",
		reflect.TypeOf((*CfnApplicationSyncJobLogsS3Props)(nil)).Elem(),
	)
}
