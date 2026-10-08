package awsnetworksecuritymanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnDeploymentMixinProps",
		reflect.TypeOf((*CfnDeploymentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnDeploymentPropsMixin",
		reflect.TypeOf((*CfnDeploymentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnDeploymentPropsMixin.AssociatedPolicyProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_AssociatedPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnDeploymentPropsMixin.AssociatedScopeProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_AssociatedScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnDeploymentPropsMixin.DeploymentConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentPropsMixin_DeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnRuleMixinProps",
		reflect.TypeOf((*CfnRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnRulePropsMixin",
		reflect.TypeOf((*CfnRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnScopeMixinProps",
		reflect.TypeOf((*CfnScopeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnScopePropsMixin",
		reflect.TypeOf((*CfnScopePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScopePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnTemplateMixinProps",
		reflect.TypeOf((*CfnTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnTemplatePropsMixin",
		reflect.TypeOf((*CfnTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networksecuritymanager.CfnTemplatePropsMixin.AssociatedRuleProperty",
		reflect.TypeOf((*CfnTemplatePropsMixin_AssociatedRuleProperty)(nil)).Elem(),
	)
}
