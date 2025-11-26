package previewawsrolesanywheremixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnCRLMixinProps",
		reflect.TypeOf((*CfnCRLMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnCRLPropsMixin",
		reflect.TypeOf((*CfnCRLPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCRLPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnProfileMixinProps",
		reflect.TypeOf((*CfnProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnProfilePropsMixin",
		reflect.TypeOf((*CfnProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnProfilePropsMixin.AttributeMappingProperty",
		reflect.TypeOf((*CfnProfilePropsMixin_AttributeMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnProfilePropsMixin.MappingRuleProperty",
		reflect.TypeOf((*CfnProfilePropsMixin_MappingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnTrustAnchorMixinProps",
		reflect.TypeOf((*CfnTrustAnchorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnTrustAnchorPropsMixin",
		reflect.TypeOf((*CfnTrustAnchorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustAnchorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnTrustAnchorPropsMixin.NotificationSettingProperty",
		reflect.TypeOf((*CfnTrustAnchorPropsMixin_NotificationSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnTrustAnchorPropsMixin.SourceDataProperty",
		reflect.TypeOf((*CfnTrustAnchorPropsMixin_SourceDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rolesanywhere.mixins.CfnTrustAnchorPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnTrustAnchorPropsMixin_SourceProperty)(nil)).Elem(),
	)
}
