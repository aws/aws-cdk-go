package previewawselasticloadbalancingv2mixins

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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerLogsMixin",
		reflect.TypeOf((*CfnLoadBalancerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerMixinProps",
		reflect.TypeOf((*CfnLoadBalancerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogs",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerNlbAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsDestProps",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerNlbAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsRecordFields",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TYPE": CfnLoadBalancerNlbAccessLogsRecordFields_TYPE,
			"VERSION": CfnLoadBalancerNlbAccessLogsRecordFields_VERSION,
			"TIME": CfnLoadBalancerNlbAccessLogsRecordFields_TIME,
			"ELB": CfnLoadBalancerNlbAccessLogsRecordFields_ELB,
			"LISTENER": CfnLoadBalancerNlbAccessLogsRecordFields_LISTENER,
			"CLIENT_PORT": CfnLoadBalancerNlbAccessLogsRecordFields_CLIENT_PORT,
			"DESTINATION_PORT": CfnLoadBalancerNlbAccessLogsRecordFields_DESTINATION_PORT,
			"CONNECTION_TIME": CfnLoadBalancerNlbAccessLogsRecordFields_CONNECTION_TIME,
			"TLS_HANDSHAKE_TIME": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_HANDSHAKE_TIME,
			"RECEIVED_BYTES": CfnLoadBalancerNlbAccessLogsRecordFields_RECEIVED_BYTES,
			"SENT_BYTES": CfnLoadBalancerNlbAccessLogsRecordFields_SENT_BYTES,
			"INCOMING_TLS_ALERT": CfnLoadBalancerNlbAccessLogsRecordFields_INCOMING_TLS_ALERT,
			"CHOSEN_CERT_ARN": CfnLoadBalancerNlbAccessLogsRecordFields_CHOSEN_CERT_ARN,
			"CHOSEN_CERT_SERIAL": CfnLoadBalancerNlbAccessLogsRecordFields_CHOSEN_CERT_SERIAL,
			"TLS_CIPHER": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_CIPHER,
			"TLS_PROTOCOL_VERSION": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_PROTOCOL_VERSION,
			"TLS_KEYEXCHANGE": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_KEYEXCHANGE,
			"DOMAIN_NAME": CfnLoadBalancerNlbAccessLogsRecordFields_DOMAIN_NAME,
			"ALPN_FE_PROTOCOL": CfnLoadBalancerNlbAccessLogsRecordFields_ALPN_FE_PROTOCOL,
			"ALPN_BE_PROTOCOL": CfnLoadBalancerNlbAccessLogsRecordFields_ALPN_BE_PROTOCOL,
			"ALPN_CLIENT_PREFERENCE_LIST": CfnLoadBalancerNlbAccessLogsRecordFields_ALPN_CLIENT_PREFERENCE_LIST,
			"TLS_CONNECTION_CREATION_TIME": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_CONNECTION_CREATION_TIME,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsS3Props",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsS3Props)(nil)).Elem(),
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
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
