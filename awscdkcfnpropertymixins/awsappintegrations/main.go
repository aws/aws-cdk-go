package awsappintegrations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationPropsMixin.ApplicationConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationPropsMixin.ApplicationSourceConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationPropsMixin.ContactHandlingProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ContactHandlingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationPropsMixin.ExternalUrlConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ExternalUrlConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnApplicationPropsMixin.IframeConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_IframeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnDataIntegrationMixinProps",
		reflect.TypeOf((*CfnDataIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnDataIntegrationPropsMixin",
		reflect.TypeOf((*CfnDataIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnDataIntegrationPropsMixin.FileConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationPropsMixin_FileConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnDataIntegrationPropsMixin.ScheduleConfigProperty",
		reflect.TypeOf((*CfnDataIntegrationPropsMixin_ScheduleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationMixinProps",
		reflect.TypeOf((*CfnEventIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationPropsMixin",
		reflect.TypeOf((*CfnEventIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationPropsMixin.EventFilterProperty",
		reflect.TypeOf((*CfnEventIntegrationPropsMixin_EventFilterProperty)(nil)).Elem(),
	)
}
