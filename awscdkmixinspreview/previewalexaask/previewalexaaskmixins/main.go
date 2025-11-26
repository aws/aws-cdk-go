package previewalexaaskmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillMixinProps",
		reflect.TypeOf((*CfnSkillMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin",
		reflect.TypeOf((*CfnSkillPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSkillPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnSkillPropsMixin_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin.OverridesProperty",
		reflect.TypeOf((*CfnSkillPropsMixin_OverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin.SkillPackageProperty",
		reflect.TypeOf((*CfnSkillPropsMixin_SkillPackageProperty)(nil)).Elem(),
	)
}
