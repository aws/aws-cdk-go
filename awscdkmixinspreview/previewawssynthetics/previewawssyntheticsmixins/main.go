package previewawssyntheticsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryMixinProps",
		reflect.TypeOf((*CfnCanaryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin",
		reflect.TypeOf((*CfnCanaryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCanaryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.ArtifactConfigProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_ArtifactConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.BaseScreenshotProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_BaseScreenshotProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.BrowserConfigProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_BrowserConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.CodeProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.DependencyProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_DependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.RetryConfigProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_RetryConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.RunConfigProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_RunConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.S3EncryptionProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_S3EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.VPCConfigProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_VPCConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin.VisualReferenceProperty",
		reflect.TypeOf((*CfnCanaryPropsMixin_VisualReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnGroupMixinProps",
		reflect.TypeOf((*CfnGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnGroupPropsMixin",
		reflect.TypeOf((*CfnGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
