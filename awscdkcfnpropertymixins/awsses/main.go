package awsses

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationMixinProps",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationSetEventDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin.CloudWatchDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_CloudWatchDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin.DimensionConfigurationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_DimensionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin.EventBridgeDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_EventBridgeDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin.EventDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_EventDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin.KinesisFirehoseDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_KinesisFirehoseDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetEventDestinationPropsMixin.SnsDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationPropsMixin_SnsDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetMixinProps",
		reflect.TypeOf((*CfnConfigurationSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.ArchivingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ArchivingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.ConditionThresholdProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ConditionThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.DashboardOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_DashboardOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.DeliveryOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_DeliveryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.GuardianOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_GuardianOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.OverallConfidenceThresholdProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_OverallConfidenceThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.ReputationOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ReputationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.SendingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_SendingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.SuppressionOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_SuppressionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.TrackingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_TrackingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.ValidationOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_ValidationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnConfigurationSetPropsMixin.VdmOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSetPropsMixin_VdmOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnContactListMixinProps",
		reflect.TypeOf((*CfnContactListMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnContactListPropsMixin",
		reflect.TypeOf((*CfnContactListPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContactListPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnContactListPropsMixin.TopicProperty",
		reflect.TypeOf((*CfnContactListPropsMixin_TopicProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnCustomVerificationEmailTemplateMixinProps",
		reflect.TypeOf((*CfnCustomVerificationEmailTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnCustomVerificationEmailTemplatePropsMixin",
		reflect.TypeOf((*CfnCustomVerificationEmailTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomVerificationEmailTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnDedicatedIpPoolMixinProps",
		reflect.TypeOf((*CfnDedicatedIpPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnDedicatedIpPoolPropsMixin",
		reflect.TypeOf((*CfnDedicatedIpPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDedicatedIpPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityMixinProps",
		reflect.TypeOf((*CfnEmailIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityPropsMixin",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEmailIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityPropsMixin.ConfigurationSetAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_ConfigurationSetAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityPropsMixin.DkimAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_DkimAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityPropsMixin.DkimSigningAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_DkimSigningAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityPropsMixin.FeedbackAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_FeedbackAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnEmailIdentityPropsMixin.MailFromAttributesProperty",
		reflect.TypeOf((*CfnEmailIdentityPropsMixin_MailFromAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonInstanceMixinProps",
		reflect.TypeOf((*CfnMailManagerAddonInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonInstancePropsMixin",
		reflect.TypeOf((*CfnMailManagerAddonInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerAddonInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonSubscriptionMixinProps",
		reflect.TypeOf((*CfnMailManagerAddonSubscriptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonSubscriptionPropsMixin",
		reflect.TypeOf((*CfnMailManagerAddonSubscriptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddressListMixinProps",
		reflect.TypeOf((*CfnMailManagerAddressListMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddressListPropsMixin",
		reflect.TypeOf((*CfnMailManagerAddressListPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerAddressListPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerArchiveMixinProps",
		reflect.TypeOf((*CfnMailManagerArchiveMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerArchivePropsMixin",
		reflect.TypeOf((*CfnMailManagerArchivePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerArchivePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerArchivePropsMixin.ArchiveRetentionProperty",
		reflect.TypeOf((*CfnMailManagerArchivePropsMixin_ArchiveRetentionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointMixinProps",
		reflect.TypeOf((*CfnMailManagerIngressPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerIngressPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin.IngressPointConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_IngressPointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin.PrivateNetworkConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_PrivateNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin.PublicNetworkConfigurationProperty",
		reflect.TypeOf((*CfnMailManagerIngressPointPropsMixin_PublicNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRelayMixinProps",
		reflect.TypeOf((*CfnMailManagerRelayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRelayPropsMixin",
		reflect.TypeOf((*CfnMailManagerRelayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerRelayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRelayPropsMixin.RelayAuthenticationProperty",
		reflect.TypeOf((*CfnMailManagerRelayPropsMixin_RelayAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetMixinProps",
		reflect.TypeOf((*CfnMailManagerRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.AddHeaderActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_AddHeaderActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.AnalysisProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_AnalysisProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.ArchiveActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_ArchiveActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.DeliverToMailboxActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_DeliverToMailboxActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.DeliverToQBusinessActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_DeliverToQBusinessActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RelayActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RelayActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.ReplaceRecipientActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_ReplaceRecipientActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleBooleanExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleBooleanExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleBooleanToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleBooleanToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleConditionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleDmarcExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleDmarcExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleIpExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleIpExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleIpToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleIpToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleIsInAddressListProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleIsInAddressListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleNumberExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleNumberExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleNumberToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleNumberToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleStringExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleStringExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleStringToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleStringToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleVerdictExpressionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleVerdictExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.RuleVerdictToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_RuleVerdictToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.S3ActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_S3ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.SendActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_SendActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerRuleSetPropsMixin.SnsActionProperty",
		reflect.TypeOf((*CfnMailManagerRuleSetPropsMixin_SnsActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyMixinProps",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressAnalysisProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressAnalysisProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressBooleanExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressBooleanExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressBooleanToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressBooleanToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressIpToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressIpv4ExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpv4ExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressIpv6ExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpv6ExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressIpv6ToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIpv6ToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressIsInAddressListProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressIsInAddressListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressStringExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressStringExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressStringToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressStringToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressTlsProtocolExpressionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressTlsProtocolExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.IngressTlsProtocolToEvaluateProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_IngressTlsProtocolToEvaluateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.PolicyConditionProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_PolicyConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin.PolicyStatementProperty",
		reflect.TypeOf((*CfnMailManagerTrafficPolicyPropsMixin_PolicyStatementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMultiRegionEndpointMixinProps",
		reflect.TypeOf((*CfnMultiRegionEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMultiRegionEndpointPropsMixin",
		reflect.TypeOf((*CfnMultiRegionEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMultiRegionEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMultiRegionEndpointPropsMixin.DetailsProperty",
		reflect.TypeOf((*CfnMultiRegionEndpointPropsMixin_DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMultiRegionEndpointPropsMixin.RouteDetailsItemsProperty",
		reflect.TypeOf((*CfnMultiRegionEndpointPropsMixin_RouteDetailsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptFilterMixinProps",
		reflect.TypeOf((*CfnReceiptFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptFilterPropsMixin",
		reflect.TypeOf((*CfnReceiptFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReceiptFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptFilterPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnReceiptFilterPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptFilterPropsMixin.IpFilterProperty",
		reflect.TypeOf((*CfnReceiptFilterPropsMixin_IpFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRuleMixinProps",
		reflect.TypeOf((*CfnReceiptRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin",
		reflect.TypeOf((*CfnReceiptRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReceiptRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.ActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.AddHeaderActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_AddHeaderActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.BounceActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_BounceActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.ConnectActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_ConnectActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.LambdaActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_LambdaActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.S3ActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_S3ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.SNSActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_SNSActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.StopActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_StopActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRulePropsMixin.WorkmailActionProperty",
		reflect.TypeOf((*CfnReceiptRulePropsMixin_WorkmailActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRuleSetMixinProps",
		reflect.TypeOf((*CfnReceiptRuleSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnReceiptRuleSetPropsMixin",
		reflect.TypeOf((*CfnReceiptRuleSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReceiptRuleSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnTemplateMixinProps",
		reflect.TypeOf((*CfnTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnTemplatePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnTemplatePropsMixin.TemplateProperty",
		reflect.TypeOf((*CfnTemplatePropsMixin_TemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnTenantMixinProps",
		reflect.TypeOf((*CfnTenantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnTenantPropsMixin",
		reflect.TypeOf((*CfnTenantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTenantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnTenantPropsMixin.ResourceAssociationProperty",
		reflect.TypeOf((*CfnTenantPropsMixin_ResourceAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnVdmAttributesMixinProps",
		reflect.TypeOf((*CfnVdmAttributesMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnVdmAttributesPropsMixin",
		reflect.TypeOf((*CfnVdmAttributesPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVdmAttributesPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnVdmAttributesPropsMixin.DashboardAttributesProperty",
		reflect.TypeOf((*CfnVdmAttributesPropsMixin_DashboardAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnVdmAttributesPropsMixin.GuardianAttributesProperty",
		reflect.TypeOf((*CfnVdmAttributesPropsMixin_GuardianAttributesProperty)(nil)).Elem(),
	)
}
