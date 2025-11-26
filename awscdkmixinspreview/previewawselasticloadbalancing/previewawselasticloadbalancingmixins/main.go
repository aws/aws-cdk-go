package previewawselasticloadbalancingmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.AccessLoggingPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_AccessLoggingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.AppCookieStickinessPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_AppCookieStickinessPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.ConnectionDrainingPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_ConnectionDrainingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.ConnectionSettingsProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_ConnectionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.HealthCheckProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_HealthCheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.LBCookieStickinessPolicyProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_LBCookieStickinessPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.ListenersProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_ListenersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancing.mixins.CfnLoadBalancerPropsMixin.PoliciesProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_PoliciesProperty)(nil)).Elem(),
	)
}
