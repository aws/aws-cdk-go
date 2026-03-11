package awsamplify

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppMixinProps",
		reflect.TypeOf((*CfnAppMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin",
		reflect.TypeOf((*CfnAppPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin.AutoBranchCreationConfigProperty",
		reflect.TypeOf((*CfnAppPropsMixin_AutoBranchCreationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin.BasicAuthConfigProperty",
		reflect.TypeOf((*CfnAppPropsMixin_BasicAuthConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin.CacheConfigProperty",
		reflect.TypeOf((*CfnAppPropsMixin_CacheConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin.CustomRuleProperty",
		reflect.TypeOf((*CfnAppPropsMixin_CustomRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnAppPropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnAppPropsMixin.JobConfigProperty",
		reflect.TypeOf((*CfnAppPropsMixin_JobConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnBranchMixinProps",
		reflect.TypeOf((*CfnBranchMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnBranchPropsMixin",
		reflect.TypeOf((*CfnBranchPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBranchPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnBranchPropsMixin.BackendProperty",
		reflect.TypeOf((*CfnBranchPropsMixin_BackendProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnBranchPropsMixin.BasicAuthConfigProperty",
		reflect.TypeOf((*CfnBranchPropsMixin_BasicAuthConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnBranchPropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnBranchPropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnDomainPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnDomainPropsMixin.CertificateSettingsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CertificateSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_amplify.CfnDomainPropsMixin.SubDomainSettingProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SubDomainSettingProperty)(nil)).Elem(),
	)
}
