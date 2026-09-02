package awssupportauthz

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitMixinProps",
		reflect.TypeOf((*CfnSupportPermitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin",
		reflect.TypeOf((*CfnSupportPermitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSupportPermitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin.ActionSetProperty",
		reflect.TypeOf((*CfnSupportPermitPropsMixin_ActionSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnSupportPermitPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin.PermitProperty",
		reflect.TypeOf((*CfnSupportPermitPropsMixin_PermitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin.ResourceSetProperty",
		reflect.TypeOf((*CfnSupportPermitPropsMixin_ResourceSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin.SigningKeyInfoProperty",
		reflect.TypeOf((*CfnSupportPermitPropsMixin_SigningKeyInfoProperty)(nil)).Elem(),
	)
}
