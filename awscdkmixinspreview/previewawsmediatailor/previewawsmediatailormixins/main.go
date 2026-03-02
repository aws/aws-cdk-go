package previewawsmediatailormixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelMixinProps",
		reflect.TypeOf((*CfnChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPolicyMixinProps",
		reflect.TypeOf((*CfnChannelPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPolicyPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.DashPlaylistSettingsProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_DashPlaylistSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.HlsPlaylistSettingsProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_HlsPlaylistSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.LogConfigurationForChannelProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_LogConfigurationForChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.RequestOutputItemProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_RequestOutputItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.SlateSourceProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_SlateSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin.TimeShiftConfigurationProperty",
		reflect.TypeOf((*CfnChannelPropsMixin_TimeShiftConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnLiveSourceMixinProps",
		reflect.TypeOf((*CfnLiveSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnLiveSourcePropsMixin",
		reflect.TypeOf((*CfnLiveSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLiveSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnLiveSourcePropsMixin.HttpPackageConfigurationProperty",
		reflect.TypeOf((*CfnLiveSourcePropsMixin_HttpPackageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogs",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsDestProps",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsFirehoseProps",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsLogGroupProps",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat.S3",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_S3_JSON,
			"PLAIN": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_S3_PLAIN,
			"W3C": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_S3_W3C,
			"PARQUET": CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsRecordFields",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"ORIGINID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ORIGINID,
			"CUSTOMERID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_CUSTOMERID,
			"EVENTTYPE": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_EVENTTYPE,
			"SESSIONTYPE": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_SESSIONTYPE,
			"REQUESTID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_REQUESTID,
			"EVENTDESCRIPTION": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_EVENTDESCRIPTION,
			"SESSIONID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_SESSIONID,
			"VODVASTRESPONSETIMEOFFSET": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_VODVASTRESPONSETIMEOFFSET,
			"ADSREQUESTURL": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ADSREQUESTURL,
			"ADMARKERS": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ADMARKERS,
			"SEGMENTATIONUPID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_SEGMENTATIONUPID,
			"SEGMENTATIONTYPEID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_SEGMENTATIONTYPEID,
			"ORIGINALTARGETURL": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ORIGINALTARGETURL,
			"UPDATEDTARGETURL": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_UPDATEDTARGETURL,
			"RAWADSREQUEST": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_RAWADSREQUEST,
			"RAWADSREQUESTINDEX": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_RAWADSREQUESTINDEX,
			"RAWADSRESPONSE": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_RAWADSRESPONSE,
			"RAWADSRESPONSEINDEX": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_RAWADSRESPONSEINDEX,
			"AVAIL": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_AVAIL,
			"VASTRESPONSE": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_VASTRESPONSE,
			"VASTAD": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_VASTAD,
			"VODCREATIVEOFFSETS": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_VODCREATIVEOFFSETS,
			"REQUESTHEADERS": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_REQUESTHEADERS,
			"BEACONINFO": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_BEACONINFO,
			"ADBREAKCORRELATIONID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ADBREAKCORRELATIONID,
			"PREFETCHSCHEDULENAME": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_PREFETCHSCHEDULENAME,
			"EVENTTIMESTAMP": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_EVENTTIMESTAMP,
			"AWSACCOUNTID": CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_AWSACCOUNTID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsS3Props",
		reflect.TypeOf((*CfnPlaybackConfigurationAdDecisionServerLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		reflect.TypeOf((*CfnPlaybackConfigurationLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlaybackConfigurationLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogs",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsDestProps",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsFirehoseProps",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsLogGroupProps",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationManifestServiceLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat.S3",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_S3_JSON,
			"PLAIN": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_S3_PLAIN,
			"W3C": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_S3_W3C,
			"PARQUET": CfnPlaybackConfigurationManifestServiceLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsRecordFields",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"CUSTOMERID": CfnPlaybackConfigurationManifestServiceLogsRecordFields_CUSTOMERID,
			"EVENTTYPE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_EVENTTYPE,
			"SESSIONTYPE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_SESSIONTYPE,
			"REQUESTID": CfnPlaybackConfigurationManifestServiceLogsRecordFields_REQUESTID,
			"EVENTDESCRIPTION": CfnPlaybackConfigurationManifestServiceLogsRecordFields_EVENTDESCRIPTION,
			"SESSIONID": CfnPlaybackConfigurationManifestServiceLogsRecordFields_SESSIONID,
			"ORIGINREQUESTURL": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ORIGINREQUESTURL,
			"MEDIATAILORPATH": CfnPlaybackConfigurationManifestServiceLogsRecordFields_MEDIATAILORPATH,
			"RESPONSEBODY": CfnPlaybackConfigurationManifestServiceLogsRecordFields_RESPONSEBODY,
			"REQUESTNEXTTOKEN": CfnPlaybackConfigurationManifestServiceLogsRecordFields_REQUESTNEXTTOKEN,
			"ASSETPATH": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ASSETPATH,
			"ORIGINFULLURL": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ORIGINFULLURL,
			"ORIGINPREFIXURL": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ORIGINPREFIXURL,
			"ADDITIONALINFO": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ADDITIONALINFO,
			"CAUSE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_CAUSE,
			"RESPONSE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_RESPONSE,
			"HTTPCODE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_HTTPCODE,
			"ERRORMESSAGE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ERRORMESSAGE,
			"ADADSRESPONSE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ADADSRESPONSE,
			"ADADSRAWRESPONSE": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ADADSRAWRESPONSE,
			"ADADSREQUEST": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ADADSREQUEST,
			"ADNUMNEWAVAILS": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ADNUMNEWAVAILS,
			"GENERATEDMEDIAPLAYLIST": CfnPlaybackConfigurationManifestServiceLogsRecordFields_GENERATEDMEDIAPLAYLIST,
			"REQUESTSTARTTIME": CfnPlaybackConfigurationManifestServiceLogsRecordFields_REQUESTSTARTTIME,
			"REQUESTENDTIME": CfnPlaybackConfigurationManifestServiceLogsRecordFields_REQUESTENDTIME,
			"REQUESTSTARTTIMEEPOCHMILLIS": CfnPlaybackConfigurationManifestServiceLogsRecordFields_REQUESTSTARTTIMEEPOCHMILLIS,
			"REQUESTENDTIMEEPOCHMILLIS": CfnPlaybackConfigurationManifestServiceLogsRecordFields_REQUESTENDTIMEEPOCHMILLIS,
			"EVENTTIMESTAMP": CfnPlaybackConfigurationManifestServiceLogsRecordFields_EVENTTIMESTAMP,
			"AWSACCOUNTID": CfnPlaybackConfigurationManifestServiceLogsRecordFields_AWSACCOUNTID,
			"ORIGINID": CfnPlaybackConfigurationManifestServiceLogsRecordFields_ORIGINID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsS3Props",
		reflect.TypeOf((*CfnPlaybackConfigurationManifestServiceLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationMixinProps",
		reflect.TypeOf((*CfnPlaybackConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlaybackConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdConditioningConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdConditioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdDecisionServerConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdDecisionServerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdMarkerPassthroughProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdMarkerPassthroughProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AdsInteractionLogProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AdsInteractionLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.AvailSuppressionProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_AvailSuppressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.BumperProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_BumperProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.CdnConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_CdnConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.DashConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_DashConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.HlsConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_HlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.HttpRequestProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_HttpRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.LivePreRollConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_LivePreRollConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.LogConfigurationProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.ManifestProcessingRulesProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_ManifestProcessingRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin.ManifestServiceInteractionLogProperty",
		reflect.TypeOf((*CfnPlaybackConfigurationPropsMixin_ManifestServiceInteractionLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogs",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationTranscodeLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsDestProps",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsFirehoseProps",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsLogGroupProps",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnPlaybackConfigurationTranscodeLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPlaybackConfigurationTranscodeLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnPlaybackConfigurationTranscodeLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnPlaybackConfigurationTranscodeLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnPlaybackConfigurationTranscodeLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnPlaybackConfigurationTranscodeLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat.S3",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPlaybackConfigurationTranscodeLogsOutputFormat_S3_JSON,
			"PLAIN": CfnPlaybackConfigurationTranscodeLogsOutputFormat_S3_PLAIN,
			"W3C": CfnPlaybackConfigurationTranscodeLogsOutputFormat_S3_W3C,
			"PARQUET": CfnPlaybackConfigurationTranscodeLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsRecordFields",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"EVENTTYPE": CfnPlaybackConfigurationTranscodeLogsRecordFields_EVENTTYPE,
			"EVENTDESCRIPTION": CfnPlaybackConfigurationTranscodeLogsRecordFields_EVENTDESCRIPTION,
			"SESSIONID": CfnPlaybackConfigurationTranscodeLogsRecordFields_SESSIONID,
			"CREATIVEUNIQUEID": CfnPlaybackConfigurationTranscodeLogsRecordFields_CREATIVEUNIQUEID,
			"PROFILENAME": CfnPlaybackConfigurationTranscodeLogsRecordFields_PROFILENAME,
			"ADURI": CfnPlaybackConfigurationTranscodeLogsRecordFields_ADURI,
			"TRANSCODEREQUESTID": CfnPlaybackConfigurationTranscodeLogsRecordFields_TRANSCODEREQUESTID,
			"CACHESTATUS": CfnPlaybackConfigurationTranscodeLogsRecordFields_CACHESTATUS,
			"EVENTTIMESTAMP": CfnPlaybackConfigurationTranscodeLogsRecordFields_EVENTTIMESTAMP,
			"AWSACCOUNTID": CfnPlaybackConfigurationTranscodeLogsRecordFields_AWSACCOUNTID,
			"ORIGINID": CfnPlaybackConfigurationTranscodeLogsRecordFields_ORIGINID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsS3Props",
		reflect.TypeOf((*CfnPlaybackConfigurationTranscodeLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationMixinProps",
		reflect.TypeOf((*CfnSourceLocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin",
		reflect.TypeOf((*CfnSourceLocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceLocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.AccessConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_AccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.DefaultSegmentDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_DefaultSegmentDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.HttpConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_HttpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.SecretsManagerAccessTokenConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_SecretsManagerAccessTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnSourceLocationPropsMixin.SegmentDeliveryConfigurationProperty",
		reflect.TypeOf((*CfnSourceLocationPropsMixin_SegmentDeliveryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnVodSourceMixinProps",
		reflect.TypeOf((*CfnVodSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnVodSourcePropsMixin",
		reflect.TypeOf((*CfnVodSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVodSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnVodSourcePropsMixin.HttpPackageConfigurationProperty",
		reflect.TypeOf((*CfnVodSourcePropsMixin_HttpPackageConfigurationProperty)(nil)).Elem(),
	)
}
