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
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
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
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPipelinePipelineLogs{}
		},
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
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
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
