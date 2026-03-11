package awsnetworkfirewall

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallMixinProps",
		reflect.TypeOf((*CfnFirewallMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyMixinProps",
		reflect.TypeOf((*CfnFirewallPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFirewallPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.ActionDefinitionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_ActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.CustomActionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_CustomActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.FirewallPolicyProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_FirewallPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.FlowTimeoutsProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_FlowTimeoutsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.IPSetProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_IPSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.PolicyVariablesProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_PolicyVariablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.PublishMetricActionProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_PublishMetricActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.StatefulEngineOptionsProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatefulEngineOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.StatefulRuleGroupOverrideProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatefulRuleGroupOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.StatefulRuleGroupReferenceProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatefulRuleGroupReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPolicyPropsMixin.StatelessRuleGroupReferenceProperty",
		reflect.TypeOf((*CfnFirewallPolicyPropsMixin_StatelessRuleGroupReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPropsMixin",
		reflect.TypeOf((*CfnFirewallPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFirewallPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPropsMixin.AvailabilityZoneMappingProperty",
		reflect.TypeOf((*CfnFirewallPropsMixin_AvailabilityZoneMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnFirewallPropsMixin.SubnetMappingProperty",
		reflect.TypeOf((*CfnFirewallPropsMixin_SubnetMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnLoggingConfigurationMixinProps",
		reflect.TypeOf((*CfnLoggingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnLoggingConfigurationPropsMixin",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoggingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnLoggingConfigurationPropsMixin.LogDestinationConfigProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_LogDestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnLoggingConfigurationPropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnLoggingConfigurationPropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupMixinProps",
		reflect.TypeOf((*CfnRuleGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin",
		reflect.TypeOf((*CfnRuleGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuleGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.ActionDefinitionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_ActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.AddressProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.CustomActionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_CustomActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.HeaderProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_HeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.IPSetProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_IPSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.IPSetReferenceProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_IPSetReferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.MatchAttributesProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_MatchAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.PortRangeProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.PortSetProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_PortSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.PublishMetricActionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_PublishMetricActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.ReferenceSetsProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_ReferenceSetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.RuleDefinitionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.RuleGroupProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.RuleOptionProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.RuleVariablesProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RuleVariablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.RulesSourceListProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RulesSourceListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.RulesSourceProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_RulesSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.StatefulRuleOptionsProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatefulRuleOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.StatefulRuleProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatefulRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.StatelessRuleProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatelessRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.StatelessRulesAndCustomActionsProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_StatelessRulesAndCustomActionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.SummaryConfigurationProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_SummaryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnRuleGroupPropsMixin.TCPFlagFieldProperty",
		reflect.TypeOf((*CfnRuleGroupPropsMixin_TCPFlagFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationMixinProps",
		reflect.TypeOf((*CfnTLSInspectionConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTLSInspectionConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.AddressProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.CheckCertificateRevocationStatusProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_CheckCertificateRevocationStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.PortRangeProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.ServerCertificateConfigurationProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_ServerCertificateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.ServerCertificateProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_ServerCertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.ServerCertificateScopeProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_ServerCertificateScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin.TLSInspectionConfigurationProperty",
		reflect.TypeOf((*CfnTLSInspectionConfigurationPropsMixin_TLSInspectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnVpcEndpointAssociationMixinProps",
		reflect.TypeOf((*CfnVpcEndpointAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnVpcEndpointAssociationPropsMixin",
		reflect.TypeOf((*CfnVpcEndpointAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcEndpointAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnVpcEndpointAssociationPropsMixin.SubnetMappingProperty",
		reflect.TypeOf((*CfnVpcEndpointAssociationPropsMixin_SubnetMappingProperty)(nil)).Elem(),
	)
}
