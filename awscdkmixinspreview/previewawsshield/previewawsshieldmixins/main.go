package previewawsshieldmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnDRTAccessMixinProps",
		reflect.TypeOf((*CfnDRTAccessMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnDRTAccessPropsMixin",
		reflect.TypeOf((*CfnDRTAccessPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDRTAccessPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProactiveEngagementMixinProps",
		reflect.TypeOf((*CfnProactiveEngagementMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProactiveEngagementPropsMixin",
		reflect.TypeOf((*CfnProactiveEngagementPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProactiveEngagementPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProactiveEngagementPropsMixin.EmergencyContactProperty",
		reflect.TypeOf((*CfnProactiveEngagementPropsMixin_EmergencyContactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogs",
		reflect.TypeOf((*CfnProtectionFlowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnProtectionFlowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsDestProps",
		reflect.TypeOf((*CfnProtectionFlowLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsFirehoseProps",
		reflect.TypeOf((*CfnProtectionFlowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsLogGroupProps",
		reflect.TypeOf((*CfnProtectionFlowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsOutputFormat",
		reflect.TypeOf((*CfnProtectionFlowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnProtectionFlowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnProtectionFlowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnProtectionFlowLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnProtectionFlowLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnProtectionFlowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnProtectionFlowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnProtectionFlowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnProtectionFlowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnProtectionFlowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnProtectionFlowLogsOutputFormat_S3_JSON,
			"PLAIN": CfnProtectionFlowLogsOutputFormat_S3_PLAIN,
			"W3C": CfnProtectionFlowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnProtectionFlowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsRecordFields",
		reflect.TypeOf((*CfnProtectionFlowLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"SRCADDR": CfnProtectionFlowLogsRecordFields_SRCADDR,
			"DSTADDR": CfnProtectionFlowLogsRecordFields_DSTADDR,
			"SRCPORT": CfnProtectionFlowLogsRecordFields_SRCPORT,
			"DSTPORT": CfnProtectionFlowLogsRecordFields_DSTPORT,
			"PROTOCOL": CfnProtectionFlowLogsRecordFields_PROTOCOL,
			"PACKETS": CfnProtectionFlowLogsRecordFields_PACKETS,
			"BYTES": CfnProtectionFlowLogsRecordFields_BYTES,
			"STARTTIME": CfnProtectionFlowLogsRecordFields_STARTTIME,
			"ENDTIME": CfnProtectionFlowLogsRecordFields_ENDTIME,
			"ACTION": CfnProtectionFlowLogsRecordFields_ACTION,
			"LOCATION": CfnProtectionFlowLogsRecordFields_LOCATION,
			"SAMPLING_RATE": CfnProtectionFlowLogsRecordFields_SAMPLING_RATE,
			"TCP_FLAGS": CfnProtectionFlowLogsRecordFields_TCP_FLAGS,
			"SRCCOUNTRY": CfnProtectionFlowLogsRecordFields_SRCCOUNTRY,
			"PROTECTION_ARN": CfnProtectionFlowLogsRecordFields_PROTECTION_ARN,
			"EVENT_TIMESTAMP": CfnProtectionFlowLogsRecordFields_EVENT_TIMESTAMP,
			"VERSION": CfnProtectionFlowLogsRecordFields_VERSION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsS3Props",
		reflect.TypeOf((*CfnProtectionFlowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionGroupMixinProps",
		reflect.TypeOf((*CfnProtectionGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionGroupPropsMixin",
		reflect.TypeOf((*CfnProtectionGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProtectionGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionLogsMixin",
		reflect.TypeOf((*CfnProtectionLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProtectionLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionMixinProps",
		reflect.TypeOf((*CfnProtectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin",
		reflect.TypeOf((*CfnProtectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProtectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnProtectionPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin.ApplicationLayerAutomaticResponseConfigurationProperty",
		reflect.TypeOf((*CfnProtectionPropsMixin_ApplicationLayerAutomaticResponseConfigurationProperty)(nil)).Elem(),
	)
}
