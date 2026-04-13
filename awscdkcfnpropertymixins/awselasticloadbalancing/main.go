package awselasticloadbalancing

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.AccessLoggingPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_AccessLoggingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.AppCookieStickinessPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_AppCookieStickinessPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.ConnectionDrainingPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_ConnectionDrainingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.ConnectionSettingsProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_ConnectionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.HealthCheckProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_HealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.LBCookieStickinessPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_LBCookieStickinessPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.ListenersProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_ListenersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.PoliciesProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_PoliciesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancing.CfnLoadBalancerPropsMixin.SourceSecurityGroupProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_SourceSecurityGroupProperty)(nil)).Elem(),
	)
}
