package awselasticbeanstalk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationPropsMixin.ApplicationResourceLifecycleConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationResourceLifecycleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationPropsMixin.ApplicationVersionLifecycleConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ApplicationVersionLifecycleConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationPropsMixin.MaxAgeRuleProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MaxAgeRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationPropsMixin.MaxCountRuleProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MaxCountRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationVersionMixinProps",
		reflect.TypeOf((*CfnApplicationVersionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationVersionPropsMixin",
		reflect.TypeOf((*CfnApplicationVersionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationVersionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnApplicationVersionPropsMixin.SourceBundleProperty",
		reflect.TypeOf((*CfnApplicationVersionPropsMixin_SourceBundleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnConfigurationTemplateMixinProps",
		reflect.TypeOf((*CfnConfigurationTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnConfigurationTemplatePropsMixin",
		reflect.TypeOf((*CfnConfigurationTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnConfigurationTemplatePropsMixin.ConfigurationOptionSettingProperty",
		reflect.TypeOf((*CfnConfigurationTemplatePropsMixin_ConfigurationOptionSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnConfigurationTemplatePropsMixin.SourceConfigurationProperty",
		reflect.TypeOf((*CfnConfigurationTemplatePropsMixin_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnEnvironmentMixinProps",
		reflect.TypeOf((*CfnEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnEnvironmentPropsMixin",
		reflect.TypeOf((*CfnEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnEnvironmentPropsMixin.OptionSettingProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_OptionSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticbeanstalk.CfnEnvironmentPropsMixin.TierProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_TierProperty)(nil)).Elem(),
	)
}
