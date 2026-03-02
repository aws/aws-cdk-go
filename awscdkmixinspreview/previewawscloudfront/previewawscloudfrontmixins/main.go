package previewawscloudfrontmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnAnycastIpListMixinProps",
		reflect.TypeOf((*CfnAnycastIpListMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnAnycastIpListPropsMixin",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnycastIpListPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnAnycastIpListPropsMixin.AnycastIpListProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_AnycastIpListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnAnycastIpListPropsMixin.IpamCidrConfigProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_IpamCidrConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnAnycastIpListPropsMixin.IpamCidrConfigResultProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_IpamCidrConfigResultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnAnycastIpListPropsMixin.TagsProperty",
		reflect.TypeOf((*CfnAnycastIpListPropsMixin_TagsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyMixinProps",
		reflect.TypeOf((*CfnCachePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin",
		reflect.TypeOf((*CfnCachePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCachePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin.CachePolicyConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_CachePolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin.CookiesConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_CookiesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin.HeadersConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_HeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin.ParametersInCacheKeyAndForwardedToOriginProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_ParametersInCacheKeyAndForwardedToOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCachePolicyPropsMixin.QueryStringsConfigProperty",
		reflect.TypeOf((*CfnCachePolicyPropsMixin_QueryStringsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityMixinProps",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityPropsMixin",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityPropsMixin.CloudFrontOriginAccessIdentityConfigProperty",
		reflect.TypeOf((*CfnCloudFrontOriginAccessIdentityPropsMixin_CloudFrontOriginAccessIdentityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnConnectionFunctionMixinProps",
		reflect.TypeOf((*CfnConnectionFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnConnectionFunctionPropsMixin",
		reflect.TypeOf((*CfnConnectionFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnConnectionFunctionPropsMixin.ConnectionFunctionConfigProperty",
		reflect.TypeOf((*CfnConnectionFunctionPropsMixin_ConnectionFunctionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnConnectionFunctionPropsMixin.KeyValueStoreAssociationProperty",
		reflect.TypeOf((*CfnConnectionFunctionPropsMixin_KeyValueStoreAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnConnectionGroupMixinProps",
		reflect.TypeOf((*CfnConnectionGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnConnectionGroupPropsMixin",
		reflect.TypeOf((*CfnConnectionGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectionGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyMixinProps",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContinuousDeploymentPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.ContinuousDeploymentPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_ContinuousDeploymentPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.SessionStickinessConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SessionStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.SingleHeaderConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.SingleHeaderPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleHeaderPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.SingleWeightConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleWeightConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.SingleWeightPolicyConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_SingleWeightPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnContinuousDeploymentPolicyPropsMixin.TrafficConfigProperty",
		reflect.TypeOf((*CfnContinuousDeploymentPolicyPropsMixin_TrafficConfigProperty)(nil)).Elem(),
	)
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
			"JSON": CfnDistributionAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnDistributionAccessLogsOutputFormat_Firehose_PLAIN,
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
			"JSON": CfnDistributionAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnDistributionAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnDistributionAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnDistributionAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsRecordFields",
		reflect.TypeOf((*CfnDistributionAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
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
			"JSON": CfnDistributionConnectionLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnDistributionConnectionLogsOutputFormat_Firehose_PLAIN,
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
			"JSON": CfnDistributionConnectionLogsOutputFormat_S3_JSON,
			"PLAIN": CfnDistributionConnectionLogsOutputFormat_S3_PLAIN,
			"W3C": CfnDistributionConnectionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnDistributionConnectionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsRecordFields",
		reflect.TypeOf((*CfnDistributionConnectionLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
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
			"CONNECTIONID": CfnDistributionConnectionLogsRecordFields_CONNECTIONID,
			"DISTRIBUTIONID": CfnDistributionConnectionLogsRecordFields_DISTRIBUTIONID,
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionMixinProps",
		reflect.TypeOf((*CfnDistributionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin",
		reflect.TypeOf((*CfnDistributionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.CacheBehaviorProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.ConnectionFunctionAssociationProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ConnectionFunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.CookiesProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CookiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.CustomErrorResponseProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CustomErrorResponseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.CustomOriginConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_CustomOriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.DefaultCacheBehaviorProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_DefaultCacheBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.DefinitionProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_DefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.DistributionConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_DistributionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.ForwardedValuesProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ForwardedValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.FunctionAssociationProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_FunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.GeoRestrictionProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_GeoRestrictionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.GrpcConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_GrpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.LambdaFunctionAssociationProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LambdaFunctionAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.LegacyCustomOriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LegacyCustomOriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.LegacyS3OriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LegacyS3OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginCustomHeaderProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginCustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginGroupFailoverCriteriaProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupFailoverCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginGroupMemberProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupMemberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginGroupMembersProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupMembersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginGroupProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginGroupsProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginGroupsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginMtlsConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginMtlsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.OriginShieldProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_OriginShieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.ParameterDefinitionProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ParameterDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.RestrictionsProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_RestrictionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.S3OriginConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_S3OriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.StatusCodesProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_StatusCodesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.StringSchemaProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_StringSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.TenantConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_TenantConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.TrustStoreConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_TrustStoreConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.ViewerCertificateProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ViewerCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.ViewerMtlsConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_ViewerMtlsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionPropsMixin.VpcOriginConfigProperty",
		reflect.TypeOf((*CfnDistributionPropsMixin_VpcOriginConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantMixinProps",
		reflect.TypeOf((*CfnDistributionTenantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionTenantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.CustomizationsProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_CustomizationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.DomainResultProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_DomainResultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.GeoRestrictionCustomizationProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_GeoRestrictionCustomizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.ManagedCertificateRequestProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_ManagedCertificateRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.ParameterProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_ParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionTenantPropsMixin.WebAclCustomizationProperty",
		reflect.TypeOf((*CfnDistributionTenantPropsMixin_WebAclCustomizationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnFunctionMixinProps",
		reflect.TypeOf((*CfnFunctionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnFunctionPropsMixin",
		reflect.TypeOf((*CfnFunctionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunctionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnFunctionPropsMixin.FunctionConfigProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnFunctionPropsMixin.FunctionMetadataProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_FunctionMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnFunctionPropsMixin.KeyValueStoreAssociationProperty",
		reflect.TypeOf((*CfnFunctionPropsMixin_KeyValueStoreAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnKeyGroupMixinProps",
		reflect.TypeOf((*CfnKeyGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnKeyGroupPropsMixin",
		reflect.TypeOf((*CfnKeyGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnKeyGroupPropsMixin.KeyGroupConfigProperty",
		reflect.TypeOf((*CfnKeyGroupPropsMixin_KeyGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnKeyValueStoreMixinProps",
		reflect.TypeOf((*CfnKeyValueStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnKeyValueStorePropsMixin",
		reflect.TypeOf((*CfnKeyValueStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyValueStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnKeyValueStorePropsMixin.ImportSourceProperty",
		reflect.TypeOf((*CfnKeyValueStorePropsMixin_ImportSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnMonitoringSubscriptionMixinProps",
		reflect.TypeOf((*CfnMonitoringSubscriptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnMonitoringSubscriptionPropsMixin",
		reflect.TypeOf((*CfnMonitoringSubscriptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitoringSubscriptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnMonitoringSubscriptionPropsMixin.MonitoringSubscriptionProperty",
		reflect.TypeOf((*CfnMonitoringSubscriptionPropsMixin_MonitoringSubscriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnMonitoringSubscriptionPropsMixin.RealtimeMetricsSubscriptionConfigProperty",
		reflect.TypeOf((*CfnMonitoringSubscriptionPropsMixin_RealtimeMetricsSubscriptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlMixinProps",
		reflect.TypeOf((*CfnOriginAccessControlMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlPropsMixin",
		reflect.TypeOf((*CfnOriginAccessControlPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginAccessControlPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlPropsMixin.OriginAccessControlConfigProperty",
		reflect.TypeOf((*CfnOriginAccessControlPropsMixin_OriginAccessControlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyMixinProps",
		reflect.TypeOf((*CfnOriginRequestPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginRequestPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin.CookiesConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_CookiesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin.HeadersConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_HeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin.OriginRequestPolicyConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_OriginRequestPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginRequestPolicyPropsMixin.QueryStringsConfigProperty",
		reflect.TypeOf((*CfnOriginRequestPolicyPropsMixin_QueryStringsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnPublicKeyMixinProps",
		reflect.TypeOf((*CfnPublicKeyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnPublicKeyPropsMixin",
		reflect.TypeOf((*CfnPublicKeyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPublicKeyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnPublicKeyPropsMixin.PublicKeyConfigProperty",
		reflect.TypeOf((*CfnPublicKeyPropsMixin_PublicKeyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnRealtimeLogConfigMixinProps",
		reflect.TypeOf((*CfnRealtimeLogConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnRealtimeLogConfigPropsMixin",
		reflect.TypeOf((*CfnRealtimeLogConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRealtimeLogConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnRealtimeLogConfigPropsMixin.EndPointProperty",
		reflect.TypeOf((*CfnRealtimeLogConfigPropsMixin_EndPointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnRealtimeLogConfigPropsMixin.KinesisStreamConfigProperty",
		reflect.TypeOf((*CfnRealtimeLogConfigPropsMixin_KinesisStreamConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyMixinProps",
		reflect.TypeOf((*CfnResponseHeadersPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResponseHeadersPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.AccessControlAllowHeadersProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlAllowHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.AccessControlAllowMethodsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlAllowMethodsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.AccessControlAllowOriginsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlAllowOriginsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.AccessControlExposeHeadersProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_AccessControlExposeHeadersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.ContentSecurityPolicyProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ContentSecurityPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.ContentTypeOptionsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ContentTypeOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.CorsConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_CorsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.CustomHeaderProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_CustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.CustomHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_CustomHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.FrameOptionsProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_FrameOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.ReferrerPolicyProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ReferrerPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.RemoveHeaderProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_RemoveHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.RemoveHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_RemoveHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.ResponseHeadersPolicyConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ResponseHeadersPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.SecurityHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_SecurityHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.ServerTimingHeadersConfigProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_ServerTimingHeadersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.StrictTransportSecurityProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_StrictTransportSecurityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnResponseHeadersPolicyPropsMixin.XSSProtectionProperty",
		reflect.TypeOf((*CfnResponseHeadersPolicyPropsMixin_XSSProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnStreamingDistributionMixinProps",
		reflect.TypeOf((*CfnStreamingDistributionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnStreamingDistributionPropsMixin",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStreamingDistributionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnStreamingDistributionPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnStreamingDistributionPropsMixin.S3OriginProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_S3OriginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnStreamingDistributionPropsMixin.StreamingDistributionConfigProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_StreamingDistributionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnStreamingDistributionPropsMixin.TrustedSignersProperty",
		reflect.TypeOf((*CfnStreamingDistributionPropsMixin_TrustedSignersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnTrustStoreMixinProps",
		reflect.TypeOf((*CfnTrustStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnTrustStorePropsMixin",
		reflect.TypeOf((*CfnTrustStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnTrustStorePropsMixin.CaCertificatesBundleS3LocationProperty",
		reflect.TypeOf((*CfnTrustStorePropsMixin_CaCertificatesBundleS3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnTrustStorePropsMixin.CaCertificatesBundleSourceProperty",
		reflect.TypeOf((*CfnTrustStorePropsMixin_CaCertificatesBundleSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnVpcOriginMixinProps",
		reflect.TypeOf((*CfnVpcOriginMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnVpcOriginPropsMixin",
		reflect.TypeOf((*CfnVpcOriginPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcOriginPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnVpcOriginPropsMixin.VpcOriginEndpointConfigProperty",
		reflect.TypeOf((*CfnVpcOriginPropsMixin_VpcOriginEndpointConfigProperty)(nil)).Elem(),
	)
}
