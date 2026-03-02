package previewawsmediapackagev2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogs",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnChannelGroupEgressAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsDestProps",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsOutputFormat",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnChannelGroupEgressAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnChannelGroupEgressAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnChannelGroupEgressAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnChannelGroupEgressAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnChannelGroupEgressAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnChannelGroupEgressAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnChannelGroupEgressAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnChannelGroupEgressAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnChannelGroupEgressAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnChannelGroupEgressAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsRecordFields",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnChannelGroupEgressAccessLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnChannelGroupEgressAccessLogsRecordFields_EVENT_TIMESTAMP,
			"CLIENT_IP": CfnChannelGroupEgressAccessLogsRecordFields_CLIENT_IP,
			"TIME_TO_FIRST_BYTE": CfnChannelGroupEgressAccessLogsRecordFields_TIME_TO_FIRST_BYTE,
			"STATUS_CODE": CfnChannelGroupEgressAccessLogsRecordFields_STATUS_CODE,
			"RECEIVED_BYTES": CfnChannelGroupEgressAccessLogsRecordFields_RECEIVED_BYTES,
			"SENT_BYTES": CfnChannelGroupEgressAccessLogsRecordFields_SENT_BYTES,
			"METHOD": CfnChannelGroupEgressAccessLogsRecordFields_METHOD,
			"REQUEST_URI_BASE": CfnChannelGroupEgressAccessLogsRecordFields_REQUEST_URI_BASE,
			"REQUEST_QUERY_PARAMS": CfnChannelGroupEgressAccessLogsRecordFields_REQUEST_QUERY_PARAMS,
			"PROTOCOL": CfnChannelGroupEgressAccessLogsRecordFields_PROTOCOL,
			"USER_AGENT": CfnChannelGroupEgressAccessLogsRecordFields_USER_AGENT,
			"DOMAIN_NAME": CfnChannelGroupEgressAccessLogsRecordFields_DOMAIN_NAME,
			"REQUEST_ID": CfnChannelGroupEgressAccessLogsRecordFields_REQUEST_ID,
			"ACCOUNT": CfnChannelGroupEgressAccessLogsRecordFields_ACCOUNT,
			"CHANNEL_ID": CfnChannelGroupEgressAccessLogsRecordFields_CHANNEL_ID,
			"CHANNEL_ARN": CfnChannelGroupEgressAccessLogsRecordFields_CHANNEL_ARN,
			"ENDPOINT_ID": CfnChannelGroupEgressAccessLogsRecordFields_ENDPOINT_ID,
			"ENDPOINT_ARN": CfnChannelGroupEgressAccessLogsRecordFields_ENDPOINT_ARN,
			"CHANNEL_GROUP_ID": CfnChannelGroupEgressAccessLogsRecordFields_CHANNEL_GROUP_ID,
			"MANIFEST_NAME": CfnChannelGroupEgressAccessLogsRecordFields_MANIFEST_NAME,
			"MANIFEST_TYPE": CfnChannelGroupEgressAccessLogsRecordFields_MANIFEST_TYPE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsS3Props",
		reflect.TypeOf((*CfnChannelGroupEgressAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogs",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnChannelGroupIngressAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsDestProps",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsOutputFormat",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnChannelGroupIngressAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnChannelGroupIngressAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnChannelGroupIngressAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnChannelGroupIngressAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnChannelGroupIngressAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnChannelGroupIngressAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnChannelGroupIngressAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnChannelGroupIngressAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnChannelGroupIngressAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnChannelGroupIngressAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsRecordFields",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnChannelGroupIngressAccessLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnChannelGroupIngressAccessLogsRecordFields_EVENT_TIMESTAMP,
			"CLIENT_IP": CfnChannelGroupIngressAccessLogsRecordFields_CLIENT_IP,
			"TIME_TO_FIRST_BYTE": CfnChannelGroupIngressAccessLogsRecordFields_TIME_TO_FIRST_BYTE,
			"STATUS_CODE": CfnChannelGroupIngressAccessLogsRecordFields_STATUS_CODE,
			"RECEIVED_BYTES": CfnChannelGroupIngressAccessLogsRecordFields_RECEIVED_BYTES,
			"SENT_BYTES": CfnChannelGroupIngressAccessLogsRecordFields_SENT_BYTES,
			"METHOD": CfnChannelGroupIngressAccessLogsRecordFields_METHOD,
			"REQUEST": CfnChannelGroupIngressAccessLogsRecordFields_REQUEST,
			"PROTOCOL": CfnChannelGroupIngressAccessLogsRecordFields_PROTOCOL,
			"USER_AGENT": CfnChannelGroupIngressAccessLogsRecordFields_USER_AGENT,
			"DOMAIN_NAME": CfnChannelGroupIngressAccessLogsRecordFields_DOMAIN_NAME,
			"REQUEST_ID": CfnChannelGroupIngressAccessLogsRecordFields_REQUEST_ID,
			"ACCOUNT": CfnChannelGroupIngressAccessLogsRecordFields_ACCOUNT,
			"CHANNEL_ID": CfnChannelGroupIngressAccessLogsRecordFields_CHANNEL_ID,
			"CHANNEL_ARN": CfnChannelGroupIngressAccessLogsRecordFields_CHANNEL_ARN,
			"CHANNEL_GROUP_ID": CfnChannelGroupIngressAccessLogsRecordFields_CHANNEL_GROUP_ID,
			"INPUT_TYPE": CfnChannelGroupIngressAccessLogsRecordFields_INPUT_TYPE,
			"INPUT_INDEX": CfnChannelGroupIngressAccessLogsRecordFields_INPUT_INDEX,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsS3Props",
		reflect.TypeOf((*CfnChannelGroupIngressAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		reflect.TypeOf((*CfnChannelGroupLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelGroupLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupMixinProps",
		reflect.TypeOf((*CfnChannelGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupPropsMixin",
		reflect.TypeOf((*CfnChannelGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPolicyMixinProps",
		reflect.TypeOf((*CfnChannelPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPolicyPropsMixin",
		reflect.TypeOf((*CfnChannelPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin",
		reflect.TypeOf((*CfnChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin.IngestEndpointProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_IngestEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin.InputSwitchConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_InputSwitchConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelPropsMixin.OutputHeaderConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_OutputHeaderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointMixinProps",
		reflect.TypeOf((*CfnOriginEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPolicyMixinProps",
		reflect.TypeOf((*CfnOriginEndpointPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPolicyPropsMixin",
		reflect.TypeOf((*CfnOriginEndpointPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginEndpointPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPolicyPropsMixin.CdnAuthConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPolicyPropsMixin_CdnAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOriginEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashBaseUrlProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashBaseUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashDvbFontDownloadProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashDvbFontDownloadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashDvbMetricsReportingProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashDvbMetricsReportingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashDvbSettingsProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashDvbSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashProgramInformationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashProgramInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashSubtitleConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashSubtitleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashTtmlConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashTtmlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.DashUtcTimingProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_DashUtcTimingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.EncryptionContractConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionContractConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.EncryptionMethodProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionMethodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.FilterConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_FilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ForceEndpointErrorConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ForceEndpointErrorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.HlsManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_HlsManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.LowLatencyHlsManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_LowLatencyHlsManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.MssManifestConfigurationProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_MssManifestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ScteDashProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ScteDashProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ScteHlsProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ScteHlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.ScteProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_ScteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.SegmentProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_SegmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.SpekeKeyProviderProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_SpekeKeyProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnOriginEndpointPropsMixin.StartTagProperty",
		reflect.TypeOf((*CfnOriginEndpointPropsMixin_StartTagProperty)(nil)).Elem(),
	)
}
