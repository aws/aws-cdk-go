package previewawsec2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogs",
		reflect.TypeOf((*CfnInstanceConsoleLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnInstanceConsoleLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsDestProps",
		reflect.TypeOf((*CfnInstanceConsoleLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsFirehoseProps",
		reflect.TypeOf((*CfnInstanceConsoleLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsLogGroupProps",
		reflect.TypeOf((*CfnInstanceConsoleLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsOutputFormat",
		reflect.TypeOf((*CfnInstanceConsoleLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnInstanceConsoleLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnInstanceConsoleLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnInstanceConsoleLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnInstanceConsoleLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnInstanceConsoleLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnInstanceConsoleLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnInstanceConsoleLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnInstanceConsoleLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsOutputFormat.S3",
		reflect.TypeOf((*CfnInstanceConsoleLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnInstanceConsoleLogsOutputFormat_S3_JSON,
			"PLAIN": CfnInstanceConsoleLogsOutputFormat_S3_PLAIN,
			"W3C": CfnInstanceConsoleLogsOutputFormat_S3_W3C,
			"PARQUET": CfnInstanceConsoleLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsRecordFields",
		reflect.TypeOf((*CfnInstanceConsoleLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnInstanceConsoleLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnInstanceConsoleLogsRecordFields_EVENT_TIMESTAMP,
			"MESSAGE": CfnInstanceConsoleLogsRecordFields_MESSAGE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsS3Props",
		reflect.TypeOf((*CfnInstanceConsoleLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceLogsMixin",
		reflect.TypeOf((*CfnInstanceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogs",
		reflect.TypeOf((*CfnRouteServerPeerEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRouteServerPeerEventLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsDestProps",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsFirehoseProps",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsLogGroupProps",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRouteServerPeerEventLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRouteServerPeerEventLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnRouteServerPeerEventLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnRouteServerPeerEventLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRouteServerPeerEventLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRouteServerPeerEventLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRouteServerPeerEventLogsOutputFormat_S3_JSON,
			"PLAIN": CfnRouteServerPeerEventLogsOutputFormat_S3_PLAIN,
			"W3C": CfnRouteServerPeerEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRouteServerPeerEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsRecordFields",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"STATUS": CfnRouteServerPeerEventLogsRecordFields_STATUS,
			"MESSAGE": CfnRouteServerPeerEventLogsRecordFields_MESSAGE,
			"RESOURCE_ARN": CfnRouteServerPeerEventLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnRouteServerPeerEventLogsRecordFields_EVENT_TIMESTAMP,
			"TYPE": CfnRouteServerPeerEventLogsRecordFields_TYPE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsS3Props",
		reflect.TypeOf((*CfnRouteServerPeerEventLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerLogsMixin",
		reflect.TypeOf((*CfnRouteServerPeerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouteServerPeerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCLogsMixin",
		reflect.TypeOf((*CfnVPCLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogs",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnVPCRoute53ResolverQueryLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsFirehoseProps",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsLogGroupProps",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnVPCRoute53ResolverQueryLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVPCRoute53ResolverQueryLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnVPCRoute53ResolverQueryLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnVPCRoute53ResolverQueryLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnVPCRoute53ResolverQueryLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnVPCRoute53ResolverQueryLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat.S3",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVPCRoute53ResolverQueryLogsOutputFormat_S3_JSON,
			"PLAIN": CfnVPCRoute53ResolverQueryLogsOutputFormat_S3_PLAIN,
			"W3C": CfnVPCRoute53ResolverQueryLogsOutputFormat_S3_W3C,
			"PARQUET": CfnVPCRoute53ResolverQueryLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsS3Props",
		reflect.TypeOf((*CfnVPCRoute53ResolverQueryLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogs",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnVPNConnectionConnectionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsDestProps",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsFirehoseProps",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsLogGroupProps",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsOutputFormat",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnVPNConnectionConnectionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVPNConnectionConnectionLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnVPNConnectionConnectionLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnVPNConnectionConnectionLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnVPNConnectionConnectionLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnVPNConnectionConnectionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVPNConnectionConnectionLogsOutputFormat_S3_JSON,
			"PLAIN": CfnVPNConnectionConnectionLogsOutputFormat_S3_PLAIN,
			"W3C": CfnVPNConnectionConnectionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnVPNConnectionConnectionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsRecordFields",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnVPNConnectionConnectionLogsRecordFields_TIMESTAMP,
			"DPD_ENABLED": CfnVPNConnectionConnectionLogsRecordFields_DPD_ENABLED,
			"NAT_T_DETECTED": CfnVPNConnectionConnectionLogsRecordFields_NAT_T_DETECTED,
			"IKE_PHASE1_STATE": CfnVPNConnectionConnectionLogsRecordFields_IKE_PHASE1_STATE,
			"IKE_PHASE2_STATE": CfnVPNConnectionConnectionLogsRecordFields_IKE_PHASE2_STATE,
			"EVENT_TIMESTAMP": CfnVPNConnectionConnectionLogsRecordFields_EVENT_TIMESTAMP,
			"DETAILS": CfnVPNConnectionConnectionLogsRecordFields_DETAILS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsS3Props",
		reflect.TypeOf((*CfnVPNConnectionConnectionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogs",
		reflect.TypeOf((*CfnVPNConnectionEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnVPNConnectionEventLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsDestProps",
		reflect.TypeOf((*CfnVPNConnectionEventLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsFirehoseProps",
		reflect.TypeOf((*CfnVPNConnectionEventLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsLogGroupProps",
		reflect.TypeOf((*CfnVPNConnectionEventLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsOutputFormat",
		reflect.TypeOf((*CfnVPNConnectionEventLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnVPNConnectionEventLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnVPNConnectionEventLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVPNConnectionEventLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnVPNConnectionEventLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnVPNConnectionEventLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnVPNConnectionEventLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnVPNConnectionEventLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnVPNConnectionEventLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsOutputFormat.S3",
		reflect.TypeOf((*CfnVPNConnectionEventLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVPNConnectionEventLogsOutputFormat_S3_JSON,
			"PLAIN": CfnVPNConnectionEventLogsOutputFormat_S3_PLAIN,
			"W3C": CfnVPNConnectionEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnVPNConnectionEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsRecordFields",
		reflect.TypeOf((*CfnVPNConnectionEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnVPNConnectionEventLogsRecordFields_TIMESTAMP,
			"TYPE": CfnVPNConnectionEventLogsRecordFields_TYPE,
			"STATUS": CfnVPNConnectionEventLogsRecordFields_STATUS,
			"MESSAGE": CfnVPNConnectionEventLogsRecordFields_MESSAGE,
			"RESOURCE_ID": CfnVPNConnectionEventLogsRecordFields_RESOURCE_ID,
			"EVENT_TIMESTAMP": CfnVPNConnectionEventLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsS3Props",
		reflect.TypeOf((*CfnVPNConnectionEventLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionLogsMixin",
		reflect.TypeOf((*CfnVPNConnectionLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPNConnectionLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceLogsMixin",
		reflect.TypeOf((*CfnVerifiedAccessInstanceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVerifiedAccessInstanceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogs",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnVerifiedAccessInstanceVerifiedAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsS3Props",
		reflect.TypeOf((*CfnVerifiedAccessInstanceVerifiedAccessLogsS3Props)(nil)).Elem(),
	)
}
