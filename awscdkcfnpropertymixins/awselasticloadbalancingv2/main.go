package awselasticloadbalancingv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerCertificateMixinProps",
		reflect.TypeOf((*CfnListenerCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerCertificatePropsMixin",
		reflect.TypeOf((*CfnListenerCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerCertificatePropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnListenerCertificatePropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerMixinProps",
		reflect.TypeOf((*CfnListenerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin",
		reflect.TypeOf((*CfnListenerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.AuthenticateCognitoConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_AuthenticateCognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.AuthenticateOidcConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_AuthenticateOidcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.FixedResponseConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_FixedResponseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.ForwardConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_ForwardConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.JwtValidationActionAdditionalClaimProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_JwtValidationActionAdditionalClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.JwtValidationConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_JwtValidationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.ListenerAttributeProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_ListenerAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.MutualAuthenticationProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_MutualAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.RedirectConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_RedirectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.TargetGroupStickinessConfigProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_TargetGroupStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerPropsMixin.TargetGroupTupleProperty",
		reflect.TypeOf((*CfnListenerPropsMixin_TargetGroupTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRuleMixinProps",
		reflect.TypeOf((*CfnListenerRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin",
		reflect.TypeOf((*CfnListenerRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.ActionProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.AuthenticateCognitoConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_AuthenticateCognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.AuthenticateOidcConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_AuthenticateOidcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.FixedResponseConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_FixedResponseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.ForwardConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_ForwardConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.HostHeaderConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_HostHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.HttpHeaderConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_HttpHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.HttpRequestMethodConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_HttpRequestMethodConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.JwtValidationActionAdditionalClaimProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_JwtValidationActionAdditionalClaimProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.JwtValidationConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_JwtValidationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.PathPatternConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_PathPatternConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.QueryStringConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_QueryStringConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.QueryStringKeyValueProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_QueryStringKeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.RedirectConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RedirectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.RewriteConfigObjectProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RewriteConfigObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.RewriteConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RewriteConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.RuleConditionProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.SourceIpConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_SourceIpConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.TargetGroupStickinessConfigProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_TargetGroupStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.TargetGroupTupleProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_TargetGroupTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnListenerRulePropsMixin.TransformProperty",
		reflect.TypeOf((*CfnListenerRulePropsMixin_TransformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin.LoadBalancerAttributeProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_LoadBalancerAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin.MinimumLoadBalancerCapacityProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_MinimumLoadBalancerCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnLoadBalancerPropsMixin.SubnetMappingProperty",
		reflect.TypeOf((*CfnLoadBalancerPropsMixin_SubnetMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTargetGroupMixinProps",
		reflect.TypeOf((*CfnTargetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTargetGroupPropsMixin",
		reflect.TypeOf((*CfnTargetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTargetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTargetGroupPropsMixin.MatcherProperty",
		reflect.TypeOf((*CfnTargetGroupPropsMixin_MatcherProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTargetGroupPropsMixin.TargetDescriptionProperty",
		reflect.TypeOf((*CfnTargetGroupPropsMixin_TargetDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTargetGroupPropsMixin.TargetGroupAttributeProperty",
		reflect.TypeOf((*CfnTargetGroupPropsMixin_TargetGroupAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreMixinProps",
		reflect.TypeOf((*CfnTrustStoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStorePropsMixin",
		reflect.TypeOf((*CfnTrustStorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationMixinProps",
		reflect.TypeOf((*CfnTrustStoreRevocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin",
		reflect.TypeOf((*CfnTrustStoreRevocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStoreRevocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin.RevocationContentProperty",
		reflect.TypeOf((*CfnTrustStoreRevocationPropsMixin_RevocationContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin.TrustStoreRevocationProperty",
		reflect.TypeOf((*CfnTrustStoreRevocationPropsMixin_TrustStoreRevocationProperty)(nil)).Elem(),
	)
}
