package previewawss3mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketLogsMixin",
		reflect.TypeOf((*CfnBucketLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBucketLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogs",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnBucketS3ServerAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsDestProps",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsOutputFormat",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnBucketS3ServerAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBucketS3ServerAccessLogsOutputFormat_Firehose_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBucketS3ServerAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBucketS3ServerAccessLogsOutputFormat_S3_JSON,
			"PARQUET": CfnBucketS3ServerAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsRecordFields",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"BUCKET_NAME": CfnBucketS3ServerAccessLogsRecordFields_BUCKET_NAME,
			"SCHEMA_VERSION_ID": CfnBucketS3ServerAccessLogsRecordFields_SCHEMA_VERSION_ID,
			"BUCKET_ARN": CfnBucketS3ServerAccessLogsRecordFields_BUCKET_ARN,
			"REQUEST_TIME": CfnBucketS3ServerAccessLogsRecordFields_REQUEST_TIME,
			"BUCKET_OWNER_ID": CfnBucketS3ServerAccessLogsRecordFields_BUCKET_OWNER_ID,
			"REMOTE_IP": CfnBucketS3ServerAccessLogsRecordFields_REMOTE_IP,
			"REQUESTER": CfnBucketS3ServerAccessLogsRecordFields_REQUESTER,
			"REQUEST_ID": CfnBucketS3ServerAccessLogsRecordFields_REQUEST_ID,
			"OPERATION": CfnBucketS3ServerAccessLogsRecordFields_OPERATION,
			"KEY_NAME": CfnBucketS3ServerAccessLogsRecordFields_KEY_NAME,
			"REQUEST_URI": CfnBucketS3ServerAccessLogsRecordFields_REQUEST_URI,
			"HTTP_STATUS": CfnBucketS3ServerAccessLogsRecordFields_HTTP_STATUS,
			"ERROR_CODE": CfnBucketS3ServerAccessLogsRecordFields_ERROR_CODE,
			"BYTES_SENT_SIZE": CfnBucketS3ServerAccessLogsRecordFields_BYTES_SENT_SIZE,
			"OBJECT_SIZE": CfnBucketS3ServerAccessLogsRecordFields_OBJECT_SIZE,
			"TOTAL_DURATION": CfnBucketS3ServerAccessLogsRecordFields_TOTAL_DURATION,
			"TURN_AROUND_DURATION": CfnBucketS3ServerAccessLogsRecordFields_TURN_AROUND_DURATION,
			"REFERER": CfnBucketS3ServerAccessLogsRecordFields_REFERER,
			"USER_AGENT": CfnBucketS3ServerAccessLogsRecordFields_USER_AGENT,
			"VERSION_ID": CfnBucketS3ServerAccessLogsRecordFields_VERSION_ID,
			"HOST_ID": CfnBucketS3ServerAccessLogsRecordFields_HOST_ID,
			"SIGNATURE_VERSION": CfnBucketS3ServerAccessLogsRecordFields_SIGNATURE_VERSION,
			"CIPHER_SUITE": CfnBucketS3ServerAccessLogsRecordFields_CIPHER_SUITE,
			"AUTHENTICATION_TYPE": CfnBucketS3ServerAccessLogsRecordFields_AUTHENTICATION_TYPE,
			"HOST_HEADER": CfnBucketS3ServerAccessLogsRecordFields_HOST_HEADER,
			"TLS_VERSION": CfnBucketS3ServerAccessLogsRecordFields_TLS_VERSION,
			"ACCESS_POINT_ARN": CfnBucketS3ServerAccessLogsRecordFields_ACCESS_POINT_ARN,
			"ACL_REQUIRED": CfnBucketS3ServerAccessLogsRecordFields_ACL_REQUIRED,
			"SOURCE_REGION": CfnBucketS3ServerAccessLogsRecordFields_SOURCE_REGION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsS3Props",
		reflect.TypeOf((*CfnBucketS3ServerAccessLogsS3Props)(nil)).Elem(),
	)
}
