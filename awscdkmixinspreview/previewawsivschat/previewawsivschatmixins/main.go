package previewawsivschatmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnLoggingConfigurationMixinProps",
		reflect.TypeOf((*CfnLoggingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnLoggingConfigurationPropsMixin",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoggingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnLoggingConfigurationPropsMixin.CloudWatchLogsDestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_CloudWatchLogsDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnLoggingConfigurationPropsMixin.DestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnLoggingConfigurationPropsMixin.FirehoseDestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_FirehoseDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnLoggingConfigurationPropsMixin.S3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_S3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogs",
		reflect.TypeOf((*CfnRoomIvsChatLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRoomIvsChatLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsFirehoseProps",
		reflect.TypeOf((*CfnRoomIvsChatLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsLogGroupProps",
		reflect.TypeOf((*CfnRoomIvsChatLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsOutputFormat",
		reflect.TypeOf((*CfnRoomIvsChatLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRoomIvsChatLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRoomIvsChatLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRoomIvsChatLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnRoomIvsChatLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnRoomIvsChatLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRoomIvsChatLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRoomIvsChatLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRoomIvsChatLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRoomIvsChatLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRoomIvsChatLogsOutputFormat_S3_JSON,
			"PLAIN": CfnRoomIvsChatLogsOutputFormat_S3_PLAIN,
			"W3C": CfnRoomIvsChatLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRoomIvsChatLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsS3Props",
		reflect.TypeOf((*CfnRoomIvsChatLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomLogsMixin",
		reflect.TypeOf((*CfnRoomLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRoomLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomMixinProps",
		reflect.TypeOf((*CfnRoomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomPropsMixin",
		reflect.TypeOf((*CfnRoomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRoomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomPropsMixin.MessageReviewHandlerProperty",
		reflect.TypeOf((*CfnRoomPropsMixin_MessageReviewHandlerProperty)(nil)).Elem(),
	)
}
