package previewawsosismixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelineLogsMixin",
		reflect.TypeOf((*CfnPipelineLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelineLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelineMixinProps",
		reflect.TypeOf((*CfnPipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogs",
		reflect.TypeOf((*CfnPipelinePipelineLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPipelinePipelineLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsFirehoseProps",
		reflect.TypeOf((*CfnPipelinePipelineLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsLogGroupProps",
		reflect.TypeOf((*CfnPipelinePipelineLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsOutputFormat",
		reflect.TypeOf((*CfnPipelinePipelineLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnPipelinePipelineLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnPipelinePipelineLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPipelinePipelineLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnPipelinePipelineLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnPipelinePipelineLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnPipelinePipelineLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnPipelinePipelineLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnPipelinePipelineLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsOutputFormat.S3",
		reflect.TypeOf((*CfnPipelinePipelineLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPipelinePipelineLogsOutputFormat_S3_JSON,
			"PLAIN": CfnPipelinePipelineLogsOutputFormat_S3_PLAIN,
			"W3C": CfnPipelinePipelineLogsOutputFormat_S3_W3C,
			"PARQUET": CfnPipelinePipelineLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePipelineLogsS3Props",
		reflect.TypeOf((*CfnPipelinePipelineLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin",
		reflect.TypeOf((*CfnPipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.BufferOptionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_BufferOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.CloudWatchLogDestinationProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_CloudWatchLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.EncryptionAtRestOptionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_EncryptionAtRestOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.LogPublishingOptionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_LogPublishingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.ResourcePolicyProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_ResourcePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.VpcAttachmentOptionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_VpcAttachmentOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.VpcEndpointProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_VpcEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin.VpcOptionsProperty",
		reflect.TypeOf((*CfnPipelinePropsMixin_VpcOptionsProperty)(nil)).Elem(),
	)
}
