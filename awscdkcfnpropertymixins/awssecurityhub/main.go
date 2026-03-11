package awssecurityhub

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAggregatorV2MixinProps",
		reflect.TypeOf((*CfnAggregatorV2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAggregatorV2PropsMixin",
		reflect.TypeOf((*CfnAggregatorV2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAggregatorV2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleMixinProps",
		reflect.TypeOf((*CfnAutomationRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin",
		reflect.TypeOf((*CfnAutomationRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutomationRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.AutomationRulesActionProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_AutomationRulesActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.AutomationRulesFindingFieldsUpdateProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_AutomationRulesFindingFieldsUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.AutomationRulesFindingFiltersProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_AutomationRulesFindingFiltersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.DateFilterProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_DateFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.DateRangeProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_DateRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.MapFilterProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_MapFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.NoteUpdateProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_NoteUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.NumberFilterProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_NumberFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.RelatedFindingProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_RelatedFindingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.SeverityUpdateProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_SeverityUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.StringFilterProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_StringFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRulePropsMixin.WorkflowUpdateProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_WorkflowUpdateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2MixinProps",
		reflect.TypeOf((*CfnAutomationRuleV2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutomationRuleV2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.AutomationRulesActionV2Property",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_AutomationRulesActionV2Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.AutomationRulesFindingFieldsUpdateV2Property",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_AutomationRulesFindingFieldsUpdateV2Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.BooleanFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_BooleanFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.CompositeFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_CompositeFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.CriteriaProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_CriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.DateFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_DateFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.DateRangeProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_DateRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.ExternalIntegrationConfigurationProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_ExternalIntegrationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.MapFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_MapFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.NumberFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_NumberFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.OcsfBooleanFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_OcsfBooleanFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.OcsfDateFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_OcsfDateFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.OcsfFindingFiltersProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_OcsfFindingFiltersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.OcsfMapFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_OcsfMapFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.OcsfNumberFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_OcsfNumberFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.OcsfStringFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_OcsfStringFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnAutomationRuleV2PropsMixin.StringFilterProperty",
		reflect.TypeOf((*CfnAutomationRuleV2PropsMixin_StringFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyMixinProps",
		reflect.TypeOf((*CfnConfigurationPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin.ParameterConfigurationProperty",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin_ParameterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin.ParameterValueProperty",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin_ParameterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin.PolicyProperty",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin_PolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin.SecurityControlCustomParameterProperty",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin_SecurityControlCustomParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin.SecurityControlsConfigurationProperty",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin_SecurityControlsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConfigurationPolicyPropsMixin.SecurityHubPolicyProperty",
		reflect.TypeOf((*CfnConfigurationPolicyPropsMixin_SecurityHubPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConnectorV2MixinProps",
		reflect.TypeOf((*CfnConnectorV2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConnectorV2PropsMixin",
		reflect.TypeOf((*CfnConnectorV2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectorV2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConnectorV2PropsMixin.JiraCloudProviderConfigurationProperty",
		reflect.TypeOf((*CfnConnectorV2PropsMixin_JiraCloudProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConnectorV2PropsMixin.ProviderProperty",
		reflect.TypeOf((*CfnConnectorV2PropsMixin_ProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnConnectorV2PropsMixin.ServiceNowProviderConfigurationProperty",
		reflect.TypeOf((*CfnConnectorV2PropsMixin_ServiceNowProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnDelegatedAdminMixinProps",
		reflect.TypeOf((*CfnDelegatedAdminMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnDelegatedAdminPropsMixin",
		reflect.TypeOf((*CfnDelegatedAdminPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDelegatedAdminPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnFindingAggregatorMixinProps",
		reflect.TypeOf((*CfnFindingAggregatorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnFindingAggregatorPropsMixin",
		reflect.TypeOf((*CfnFindingAggregatorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFindingAggregatorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnHubMixinProps",
		reflect.TypeOf((*CfnHubMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnHubPropsMixin",
		reflect.TypeOf((*CfnHubPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHubPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnHubV2MixinProps",
		reflect.TypeOf((*CfnHubV2MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnHubV2PropsMixin",
		reflect.TypeOf((*CfnHubV2PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHubV2PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightMixinProps",
		reflect.TypeOf((*CfnInsightMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin",
		reflect.TypeOf((*CfnInsightPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInsightPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.AwsSecurityFindingFiltersProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_AwsSecurityFindingFiltersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.BooleanFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_BooleanFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.DateFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_DateFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.DateRangeProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_DateRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.IpFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_IpFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.KeywordFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_KeywordFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.MapFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_MapFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.NumberFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_NumberFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnInsightPropsMixin.StringFilterProperty",
		reflect.TypeOf((*CfnInsightPropsMixin_StringFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnOrganizationConfigurationMixinProps",
		reflect.TypeOf((*CfnOrganizationConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnOrganizationConfigurationPropsMixin",
		reflect.TypeOf((*CfnOrganizationConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnPolicyAssociationMixinProps",
		reflect.TypeOf((*CfnPolicyAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnPolicyAssociationPropsMixin",
		reflect.TypeOf((*CfnPolicyAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPolicyAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnProductSubscriptionMixinProps",
		reflect.TypeOf((*CfnProductSubscriptionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnProductSubscriptionPropsMixin",
		reflect.TypeOf((*CfnProductSubscriptionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProductSubscriptionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnSecurityControlMixinProps",
		reflect.TypeOf((*CfnSecurityControlMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnSecurityControlPropsMixin",
		reflect.TypeOf((*CfnSecurityControlPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityControlPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnSecurityControlPropsMixin.ParameterConfigurationProperty",
		reflect.TypeOf((*CfnSecurityControlPropsMixin_ParameterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnSecurityControlPropsMixin.ParameterValueProperty",
		reflect.TypeOf((*CfnSecurityControlPropsMixin_ParameterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnStandardMixinProps",
		reflect.TypeOf((*CfnStandardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnStandardPropsMixin",
		reflect.TypeOf((*CfnStandardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStandardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnStandardPropsMixin.StandardsControlProperty",
		reflect.TypeOf((*CfnStandardPropsMixin_StandardsControlProperty)(nil)).Elem(),
	)
}
