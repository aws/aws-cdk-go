package previewawsssmguiconnectmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesMixinProps",
		reflect.TypeOf((*CfnPreferencesMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin",
		reflect.TypeOf((*CfnPreferencesPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPreferencesPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin.ConnectionRecordingPreferencesProperty",
		reflect.TypeOf((*CfnPreferencesPropsMixin_ConnectionRecordingPreferencesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin.RecordingDestinationsProperty",
		reflect.TypeOf((*CfnPreferencesPropsMixin_RecordingDestinationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmguiconnect.mixins.CfnPreferencesPropsMixin.S3BucketProperty",
		reflect.TypeOf((*CfnPreferencesPropsMixin_S3BucketProperty)(nil)).Elem(),
	)
}
