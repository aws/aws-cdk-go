package awsssmguiconnect

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesMixinProps",
		reflect.TypeOf((*CfnPreferencesMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin",
		reflect.TypeOf((*CfnPreferencesPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPreferencesPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin.ConnectionRecordingPreferencesProperty",
		reflect.TypeOf((*CfnPreferencesPropsMixin_ConnectionRecordingPreferencesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin.RecordingDestinationsProperty",
		reflect.TypeOf((*CfnPreferencesPropsMixin_RecordingDestinationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ssmguiconnect.CfnPreferencesPropsMixin.S3BucketProperty",
		reflect.TypeOf((*CfnPreferencesPropsMixin_S3BucketProperty)(nil)).Elem(),
	)
}
