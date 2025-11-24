package mixinsawselasticloadbalancingv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerCertificateMixinProps",
		reflect.TypeOf((*CfnListenerCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerCertificatePropsMixin",
		reflect.TypeOf((*CfnListenerCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerCertificatePropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnListenerCertificatePropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerMixinProps",
		reflect.TypeOf((*CfnListenerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin",
		reflect.TypeOf((*CfnListenerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.AuthenticateCognitoConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_AuthenticateCognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.AuthenticateOidcConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_AuthenticateOidcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.FixedResponseConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_FixedResponseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.ForwardConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_ForwardConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.JwtValidationActionAdditionalClaimProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_JwtValidationActionAdditionalClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.JwtValidationConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_JwtValidationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.ListenerAttributeProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_ListenerAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.MutualAuthenticationProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_MutualAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.RedirectConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_RedirectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.TargetGroupStickinessConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_TargetGroupStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerPropsMixin.TargetGroupTupleProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_TargetGroupTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRuleMixinProps",
		reflect.TypeOf((*CfnListenerRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin",
		reflect.TypeOf((*CfnListenerRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.ActionProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.AuthenticateCognitoConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_AuthenticateCognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.AuthenticateOidcConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_AuthenticateOidcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.FixedResponseConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_FixedResponseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.ForwardConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_ForwardConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.HostHeaderConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_HostHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.HttpHeaderConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_HttpHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.HttpRequestMethodConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_HttpRequestMethodConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.JwtValidationActionAdditionalClaimProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_JwtValidationActionAdditionalClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.JwtValidationConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_JwtValidationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.PathPatternConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_PathPatternConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.QueryStringConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_QueryStringConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.QueryStringKeyValueProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_QueryStringKeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.RedirectConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RedirectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.RewriteConfigObjectProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RewriteConfigObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.RewriteConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RewriteConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.RuleConditionProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.SourceIpConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_SourceIpConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.TargetGroupStickinessConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_TargetGroupStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.TargetGroupTupleProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_TargetGroupTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnListenerRulePropsMixin.TransformProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_TransformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerPropsMixin.LoadBalancerAttributeProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_LoadBalancerAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerPropsMixin.MinimumLoadBalancerCapacityProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_MinimumLoadBalancerCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerPropsMixin.SubnetMappingProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_SubnetMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTargetGroupMixinProps",
		reflect.TypeOf((*CfnTargetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTargetGroupPropsMixin",
		reflect.TypeOf((*CfnTargetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTargetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTargetGroupPropsMixin.MatcherProperty",
		reflect.TypeOf((*CfnTargetGroupPropsMixin_MatcherProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTargetGroupPropsMixin.TargetDescriptionProperty",
		reflect.TypeOf((*CfnTargetGroupPropsMixin_TargetDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTargetGroupPropsMixin.TargetGroupAttributeProperty",
		reflect.TypeOf((*CfnTargetGroupPropsMixin_TargetGroupAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTrustStoreMixinProps",
		reflect.TypeOf((*CfnTrustStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTrustStorePropsMixin",
		reflect.TypeOf((*CfnTrustStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTrustStoreRevocationMixinProps",
		reflect.TypeOf((*CfnTrustStoreRevocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTrustStoreRevocationPropsMixin",
		reflect.TypeOf((*CfnTrustStoreRevocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStoreRevocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTrustStoreRevocationPropsMixin.RevocationContentProperty",
		reflect.TypeOf((*CfnTrustStoreRevocationPropsMixin_RevocationContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnTrustStoreRevocationPropsMixin.TrustStoreRevocationProperty",
		reflect.TypeOf((*CfnTrustStoreRevocationPropsMixin_TrustStoreRevocationProperty)(nil)).Elem(),
	)
}
