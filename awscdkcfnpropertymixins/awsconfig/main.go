package awsconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnAggregationAuthorizationMixinProps",
		reflect.TypeOf((*CfnAggregationAuthorizationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnAggregationAuthorizationPropsMixin",
		reflect.TypeOf((*CfnAggregationAuthorizationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAggregationAuthorizationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRuleMixinProps",
		reflect.TypeOf((*CfnConfigRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin",
		reflect.TypeOf((*CfnConfigRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin.ComplianceProperty",
		reflect.TypeOf((*CfnConfigRulePropsMixin_ComplianceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin.CustomPolicyDetailsProperty",
		reflect.TypeOf((*CfnConfigRulePropsMixin_CustomPolicyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin.EvaluationModeConfigurationProperty",
		reflect.TypeOf((*CfnConfigRulePropsMixin_EvaluationModeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin.ScopeProperty",
		reflect.TypeOf((*CfnConfigRulePropsMixin_ScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin.SourceDetailProperty",
		reflect.TypeOf((*CfnConfigRulePropsMixin_SourceDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigRulePropsMixin.SourceProperty",
		reflect.TypeOf((*CfnConfigRulePropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationAggregatorMixinProps",
		reflect.TypeOf((*CfnConfigurationAggregatorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationAggregatorPropsMixin",
		reflect.TypeOf((*CfnConfigurationAggregatorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationAggregatorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationAggregatorPropsMixin.AccountAggregationSourceProperty",
		reflect.TypeOf((*CfnConfigurationAggregatorPropsMixin_AccountAggregationSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationAggregatorPropsMixin.OrganizationAggregationSourceProperty",
		reflect.TypeOf((*CfnConfigurationAggregatorPropsMixin_OrganizationAggregationSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderMixinProps",
		reflect.TypeOf((*CfnConfigurationRecorderMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderPropsMixin",
		reflect.TypeOf((*CfnConfigurationRecorderPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationRecorderPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderPropsMixin.ExclusionByResourceTypesProperty",
		reflect.TypeOf((*CfnConfigurationRecorderPropsMixin_ExclusionByResourceTypesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderPropsMixin.RecordingGroupProperty",
		reflect.TypeOf((*CfnConfigurationRecorderPropsMixin_RecordingGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderPropsMixin.RecordingModeOverrideProperty",
		reflect.TypeOf((*CfnConfigurationRecorderPropsMixin_RecordingModeOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderPropsMixin.RecordingModeProperty",
		reflect.TypeOf((*CfnConfigurationRecorderPropsMixin_RecordingModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConfigurationRecorderPropsMixin.RecordingStrategyProperty",
		reflect.TypeOf((*CfnConfigurationRecorderPropsMixin_RecordingStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackMixinProps",
		reflect.TypeOf((*CfnConformancePackMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin",
		reflect.TypeOf((*CfnConformancePackPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConformancePackPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin.ConformancePackInputParameterProperty",
		reflect.TypeOf((*CfnConformancePackPropsMixin_ConformancePackInputParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin.TemplateSSMDocumentDetailsProperty",
		reflect.TypeOf((*CfnConformancePackPropsMixin_TemplateSSMDocumentDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConnectorMixinProps",
		reflect.TypeOf((*CfnConnectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConnectorPropsMixin",
		reflect.TypeOf((*CfnConnectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConnectorPropsMixin.AzureConnectorConfigurationProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_AzureConnectorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConnectorPropsMixin.ConnectorConfigurationProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_ConnectorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnDeliveryChannelMixinProps",
		reflect.TypeOf((*CfnDeliveryChannelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnDeliveryChannelPropsMixin",
		reflect.TypeOf((*CfnDeliveryChannelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryChannelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnDeliveryChannelPropsMixin.ConfigSnapshotDeliveryPropertiesProperty",
		reflect.TypeOf((*CfnDeliveryChannelPropsMixin_ConfigSnapshotDeliveryPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConfigRuleMixinProps",
		reflect.TypeOf((*CfnOrganizationConfigRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConfigRulePropsMixin",
		reflect.TypeOf((*CfnOrganizationConfigRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationConfigRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConfigRulePropsMixin.OrganizationCustomPolicyRuleMetadataProperty",
		reflect.TypeOf((*CfnOrganizationConfigRulePropsMixin_OrganizationCustomPolicyRuleMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConfigRulePropsMixin.OrganizationCustomRuleMetadataProperty",
		reflect.TypeOf((*CfnOrganizationConfigRulePropsMixin_OrganizationCustomRuleMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConfigRulePropsMixin.OrganizationManagedRuleMetadataProperty",
		reflect.TypeOf((*CfnOrganizationConfigRulePropsMixin_OrganizationManagedRuleMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConformancePackMixinProps",
		reflect.TypeOf((*CfnOrganizationConformancePackMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConformancePackPropsMixin",
		reflect.TypeOf((*CfnOrganizationConformancePackPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationConformancePackPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnOrganizationConformancePackPropsMixin.ConformancePackInputParameterProperty",
		reflect.TypeOf((*CfnOrganizationConformancePackPropsMixin_ConformancePackInputParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationMixinProps",
		reflect.TypeOf((*CfnRemediationConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin",
		reflect.TypeOf((*CfnRemediationConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRemediationConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin.ExecutionControlsProperty",
		reflect.TypeOf((*CfnRemediationConfigurationPropsMixin_ExecutionControlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin.RemediationParameterValueProperty",
		reflect.TypeOf((*CfnRemediationConfigurationPropsMixin_RemediationParameterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin.ResourceValueProperty",
		reflect.TypeOf((*CfnRemediationConfigurationPropsMixin_ResourceValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin.SsmControlsProperty",
		reflect.TypeOf((*CfnRemediationConfigurationPropsMixin_SsmControlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin.StaticValueProperty",
		reflect.TypeOf((*CfnRemediationConfigurationPropsMixin_StaticValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnStoredQueryMixinProps",
		reflect.TypeOf((*CfnStoredQueryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnStoredQueryPropsMixin",
		reflect.TypeOf((*CfnStoredQueryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStoredQueryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
