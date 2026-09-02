package previewawscloudfrontmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogs",
		reflect.TypeOf((*CfnDistributionAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnDistributionAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsDestProps",
		reflect.TypeOf((*CfnDistributionAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnDistributionAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnDistributionAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsOutputFormat",
		reflect.TypeOf((*CfnDistributionAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnDistributionAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnDistributionAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnDistributionAccessLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnDistributionAccessLogsOutputFormat_Firehose_JSON,
			"RAW": CfnDistributionAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnDistributionAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnDistributionAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnDistributionAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnDistributionAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnDistributionAccessLogsOutputFormat_S3_PLAIN,
			"JSON": CfnDistributionAccessLogsOutputFormat_S3_JSON,
			"W3C": CfnDistributionAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnDistributionAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsRecordFields",
		reflect.TypeOf((*CfnDistributionAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnDistributionAccessLogsRecordFields_TIMESTAMP,
			"DATE": CfnDistributionAccessLogsRecordFields_DATE,
			"TIME": CfnDistributionAccessLogsRecordFields_TIME,
			"X_EDGE_LOCATION": CfnDistributionAccessLogsRecordFields_X_EDGE_LOCATION,
			"SC_BYTES": CfnDistributionAccessLogsRecordFields_SC_BYTES,
			"C_IP": CfnDistributionAccessLogsRecordFields_C_IP,
			"CS_METHOD": CfnDistributionAccessLogsRecordFields_CS_METHOD,
			"CS_HOST_": CfnDistributionAccessLogsRecordFields_CS_HOST_,
			"CS_URI_STEM": CfnDistributionAccessLogsRecordFields_CS_URI_STEM,
			"SC_STATUS": CfnDistributionAccessLogsRecordFields_SC_STATUS,
			"CS_REFERER_": CfnDistributionAccessLogsRecordFields_CS_REFERER_,
			"CS_USER_AGENT_": CfnDistributionAccessLogsRecordFields_CS_USER_AGENT_,
			"CS_URI_QUERY": CfnDistributionAccessLogsRecordFields_CS_URI_QUERY,
			"CS_COOKIE_": CfnDistributionAccessLogsRecordFields_CS_COOKIE_,
			"X_EDGE_RESULT_TYPE": CfnDistributionAccessLogsRecordFields_X_EDGE_RESULT_TYPE,
			"X_EDGE_REQUEST_ID": CfnDistributionAccessLogsRecordFields_X_EDGE_REQUEST_ID,
			"X_HOST_HEADER": CfnDistributionAccessLogsRecordFields_X_HOST_HEADER,
			"CS_PROTOCOL": CfnDistributionAccessLogsRecordFields_CS_PROTOCOL,
			"CS_BYTES": CfnDistributionAccessLogsRecordFields_CS_BYTES,
			"TIME_TAKEN": CfnDistributionAccessLogsRecordFields_TIME_TAKEN,
			"X_FORWARDED_FOR": CfnDistributionAccessLogsRecordFields_X_FORWARDED_FOR,
			"SSL_PROTOCOL": CfnDistributionAccessLogsRecordFields_SSL_PROTOCOL,
			"SSL_CIPHER": CfnDistributionAccessLogsRecordFields_SSL_CIPHER,
			"X_EDGE_RESPONSE_RESULT_TYPE": CfnDistributionAccessLogsRecordFields_X_EDGE_RESPONSE_RESULT_TYPE,
			"CS_PROTOCOL_VERSION": CfnDistributionAccessLogsRecordFields_CS_PROTOCOL_VERSION,
			"FLE_STATUS": CfnDistributionAccessLogsRecordFields_FLE_STATUS,
			"FLE_ENCRYPTED_FIELDS": CfnDistributionAccessLogsRecordFields_FLE_ENCRYPTED_FIELDS,
			"C_PORT": CfnDistributionAccessLogsRecordFields_C_PORT,
			"TIME_TO_FIRST_BYTE": CfnDistributionAccessLogsRecordFields_TIME_TO_FIRST_BYTE,
			"X_EDGE_DETAILED_RESULT_TYPE": CfnDistributionAccessLogsRecordFields_X_EDGE_DETAILED_RESULT_TYPE,
			"SC_CONTENT_TYPE": CfnDistributionAccessLogsRecordFields_SC_CONTENT_TYPE,
			"SC_CONTENT_LEN": CfnDistributionAccessLogsRecordFields_SC_CONTENT_LEN,
			"SC_RANGE_START": CfnDistributionAccessLogsRecordFields_SC_RANGE_START,
			"SC_RANGE_END": CfnDistributionAccessLogsRecordFields_SC_RANGE_END,
			"TIMESTAMP_MS_": CfnDistributionAccessLogsRecordFields_TIMESTAMP_MS_,
			"ORIGIN_FBL": CfnDistributionAccessLogsRecordFields_ORIGIN_FBL,
			"ORIGIN_LBL": CfnDistributionAccessLogsRecordFields_ORIGIN_LBL,
			"ASN": CfnDistributionAccessLogsRecordFields_ASN,
			"C_COUNTRY": CfnDistributionAccessLogsRecordFields_C_COUNTRY,
			"CACHE_BEHAVIOR_PATH_PATTERN": CfnDistributionAccessLogsRecordFields_CACHE_BEHAVIOR_PATH_PATTERN,
			"DISTRIBUTION_TENANT_ID": CfnDistributionAccessLogsRecordFields_DISTRIBUTION_TENANT_ID,
			"CONNECTION_ID": CfnDistributionAccessLogsRecordFields_CONNECTION_ID,
			"VIEWER_REQUEST_LOG_DATA": CfnDistributionAccessLogsRecordFields_VIEWER_REQUEST_LOG_DATA,
			"VIEWER_RESPONSE_LOG_DATA": CfnDistributionAccessLogsRecordFields_VIEWER_RESPONSE_LOG_DATA,
			"CMCD_ENCODED_BITRATE": CfnDistributionAccessLogsRecordFields_CMCD_ENCODED_BITRATE,
			"CMCD_BUFFER_LENGTH": CfnDistributionAccessLogsRecordFields_CMCD_BUFFER_LENGTH,
			"CMCD_BUFFER_STARVATION": CfnDistributionAccessLogsRecordFields_CMCD_BUFFER_STARVATION,
			"CMCD_CONTENT_ID": CfnDistributionAccessLogsRecordFields_CMCD_CONTENT_ID,
			"CMCD_OBJECT_DURATION": CfnDistributionAccessLogsRecordFields_CMCD_OBJECT_DURATION,
			"CMCD_DEADLINE": CfnDistributionAccessLogsRecordFields_CMCD_DEADLINE,
			"CMCD_MEASURED_THROUGHPUT": CfnDistributionAccessLogsRecordFields_CMCD_MEASURED_THROUGHPUT,
			"CMCD_NEXT_OBJECT_REQUEST": CfnDistributionAccessLogsRecordFields_CMCD_NEXT_OBJECT_REQUEST,
			"CMCD_NEXT_RANGE_REQUEST": CfnDistributionAccessLogsRecordFields_CMCD_NEXT_RANGE_REQUEST,
			"CMCD_OBJECT_TYPE": CfnDistributionAccessLogsRecordFields_CMCD_OBJECT_TYPE,
			"CMCD_PLAYBACK_RATE": CfnDistributionAccessLogsRecordFields_CMCD_PLAYBACK_RATE,
			"CMCD_REQUESTED_MAXIMUM_THROUGHPUT": CfnDistributionAccessLogsRecordFields_CMCD_REQUESTED_MAXIMUM_THROUGHPUT,
			"CMCD_STREAMING_FORMAT": CfnDistributionAccessLogsRecordFields_CMCD_STREAMING_FORMAT,
			"CMCD_SESSION_ID": CfnDistributionAccessLogsRecordFields_CMCD_SESSION_ID,
			"CMCD_STREAM_TYPE": CfnDistributionAccessLogsRecordFields_CMCD_STREAM_TYPE,
			"CMCD_STARTUP": CfnDistributionAccessLogsRecordFields_CMCD_STARTUP,
			"CMCD_TOP_BITRATE": CfnDistributionAccessLogsRecordFields_CMCD_TOP_BITRATE,
			"CMCD_VERSION": CfnDistributionAccessLogsRecordFields_CMCD_VERSION,
			"S_IP": CfnDistributionAccessLogsRecordFields_S_IP,
			"PRIMARY_DISTRIBUTION_ID": CfnDistributionAccessLogsRecordFields_PRIMARY_DISTRIBUTION_ID,
			"PRIMARY_DISTRIBUTION_DNS_NAME": CfnDistributionAccessLogsRecordFields_PRIMARY_DISTRIBUTION_DNS_NAME,
			"R_HOST": CfnDistributionAccessLogsRecordFields_R_HOST,
			"SR_REASON": CfnDistributionAccessLogsRecordFields_SR_REASON,
			"X_EDGE_MQCS": CfnDistributionAccessLogsRecordFields_X_EDGE_MQCS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsS3Props",
		reflect.TypeOf((*CfnDistributionAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogs",
		reflect.TypeOf((*CfnDistributionConnectionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnDistributionConnectionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsDestProps",
		reflect.TypeOf((*CfnDistributionConnectionLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsFirehoseProps",
		reflect.TypeOf((*CfnDistributionConnectionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsLogGroupProps",
		reflect.TypeOf((*CfnDistributionConnectionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsOutputFormat",
		reflect.TypeOf((*CfnDistributionConnectionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnDistributionConnectionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnDistributionConnectionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnDistributionConnectionLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnDistributionConnectionLogsOutputFormat_Firehose_JSON,
			"RAW": CfnDistributionConnectionLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnDistributionConnectionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnDistributionConnectionLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnDistributionConnectionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnDistributionConnectionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnDistributionConnectionLogsOutputFormat_S3_PLAIN,
			"JSON": CfnDistributionConnectionLogsOutputFormat_S3_JSON,
			"W3C": CfnDistributionConnectionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnDistributionConnectionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsRecordFields",
		reflect.TypeOf((*CfnDistributionConnectionLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnDistributionConnectionLogsRecordFields_TIMESTAMP,
			"CONNECTIONSTATUS": CfnDistributionConnectionLogsRecordFields_CONNECTIONSTATUS,
			"CLIENTIP": CfnDistributionConnectionLogsRecordFields_CLIENTIP,
			"CLIENTPORT": CfnDistributionConnectionLogsRecordFields_CLIENTPORT,
			"SERVERIP": CfnDistributionConnectionLogsRecordFields_SERVERIP,
			"DISTRIBUTIONTENANTID": CfnDistributionConnectionLogsRecordFields_DISTRIBUTIONTENANTID,
			"TLSPROTOCOL": CfnDistributionConnectionLogsRecordFields_TLSPROTOCOL,
			"TLSCIPHER": CfnDistributionConnectionLogsRecordFields_TLSCIPHER,
			"TLSHANDSHAKEDURATION": CfnDistributionConnectionLogsRecordFields_TLSHANDSHAKEDURATION,
			"TLSSNI": CfnDistributionConnectionLogsRecordFields_TLSSNI,
			"CLIENTLEAFCERTSERIALNUMBER": CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTSERIALNUMBER,
			"CLIENTLEAFCERTSUBJECT": CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTSUBJECT,
			"CLIENTLEAFCERTISSUER": CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTISSUER,
			"CLIENTLEAFCERTVALIDITY": CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTVALIDITY,
			"CONNECTIONLOGCUSTOMDATA": CfnDistributionConnectionLogsRecordFields_CONNECTIONLOGCUSTOMDATA,
			"EVENTTIMESTAMP": CfnDistributionConnectionLogsRecordFields_EVENTTIMESTAMP,
			"DISTRIBUTIONID": CfnDistributionConnectionLogsRecordFields_DISTRIBUTIONID,
			"CONNECTIONID": CfnDistributionConnectionLogsRecordFields_CONNECTIONID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsS3Props",
		reflect.TypeOf((*CfnDistributionConnectionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionLogsMixin",
		reflect.TypeOf((*CfnDistributionLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
