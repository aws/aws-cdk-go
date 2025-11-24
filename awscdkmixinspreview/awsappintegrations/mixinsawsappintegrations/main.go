package mixinsawsappintegrations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationPropsMixin.ApplicationConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationPropsMixin.ApplicationSourceConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationPropsMixin.ContactHandlingProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ContactHandlingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationPropsMixin.ExternalUrlConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ExternalUrlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnApplicationPropsMixin.IframeConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_IframeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationMixinProps",
		reflect.TypeOf((*CfnDataIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin",
		reflect.TypeOf((*CfnDataIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin.FileConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationPropsMixin_FileConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin.ScheduleConfigProperty",
		reflect.TypeOf((*CfnDataIntegrationPropsMixin_ScheduleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnEventIntegrationMixinProps",
		reflect.TypeOf((*CfnEventIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnEventIntegrationPropsMixin",
		reflect.TypeOf((*CfnEventIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnEventIntegrationPropsMixin.EventFilterProperty",
		reflect.TypeOf((*CfnEventIntegrationPropsMixin_EventFilterProperty)(nil)).Elem(),
	)
}
