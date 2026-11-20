package awscloudwatch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantMixinProps",
		reflect.TypeOf((*CfnAccessGrantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin",
		reflect.TypeOf((*CfnAccessGrantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessGrantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin.AccessGrantPrincipalAttributeProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_AccessGrantPrincipalAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin.AccessGrantPrincipalProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_AccessGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin.ResourceScopeProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_ResourceScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin.RowScopeProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_RowScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessGrantPropsMixin.ScopedActionsProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_ScopedActionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessProfileMixinProps",
		reflect.TypeOf((*CfnAccessProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAccessProfilePropsMixin",
		reflect.TypeOf((*CfnAccessProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmMixinProps",
		reflect.TypeOf((*CfnAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmMuteRuleMixinProps",
		reflect.TypeOf((*CfnAlarmMuteRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmMuteRulePropsMixin",
		reflect.TypeOf((*CfnAlarmMuteRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarmMuteRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmMuteRulePropsMixin.MuteTargetsProperty",
		reflect.TypeOf((*CfnAlarmMuteRulePropsMixin_MuteTargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmMuteRulePropsMixin.RuleProperty",
		reflect.TypeOf((*CfnAlarmMuteRulePropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmMuteRulePropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnAlarmMuteRulePropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin",
		reflect.TypeOf((*CfnAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.ActionConfigurationProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_ActionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.AlarmPromQLCriteriaProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_AlarmPromQLCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.AlarmTagsPropagationProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_AlarmTagsPropagationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.EvaluationCriteriaProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_EvaluationCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.EvaluationWindowProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_EvaluationWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.MetricDataQueryProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.MetricStatProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.WallClockWindowProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_WallClockWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlarmPropsMixin.WarmUpConfigurationProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_WarmUpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertMixinProps",
		reflect.TypeOf((*CfnAlertMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin",
		reflect.TypeOf((*CfnAlertPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlertPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.EvaluationProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_EvaluationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.NoDataProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_NoDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.NotificationRuleProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_NotificationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.QueryProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_QueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.TargetProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.TelemetryRuleProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_TelemetryRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAlertPropsMixin.TriggerProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_TriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.MetricCharacteristicsProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricCharacteristicsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.MetricDataQueryProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.MetricMathAnomalyDetectorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricMathAnomalyDetectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.MetricStatProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.RangeProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnAnomalyDetectorPropsMixin.SingleMetricAnomalyDetectorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_SingleMetricAnomalyDetectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnCompositeAlarmMixinProps",
		reflect.TypeOf((*CfnCompositeAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnCompositeAlarmPropsMixin",
		reflect.TypeOf((*CfnCompositeAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCompositeAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnCompositeAlarmPropsMixin.ActionConfigurationProperty",
		reflect.TypeOf((*CfnCompositeAlarmPropsMixin_ActionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnCompositeAlarmPropsMixin.AlarmTagsPropagationProperty",
		reflect.TypeOf((*CfnCompositeAlarmPropsMixin_AlarmTagsPropagationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDashboardMixinProps",
		reflect.TypeOf((*CfnDashboardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDashboardPropsMixin",
		reflect.TypeOf((*CfnDashboardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationMixinProps",
		reflect.TypeOf((*CfnDomainAccessGrantForOrganizationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin",
		reflect.TypeOf((*CfnDomainAccessGrantForOrganizationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin.OrganizationAccessGrantPrincipalProperty",
		reflect.TypeOf((*CfnDomainAccessGrantForOrganizationPropsMixin_OrganizationAccessGrantPrincipalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin.PrincipalAttributeProperty",
		reflect.TypeOf((*CfnDomainAccessGrantForOrganizationPropsMixin_PrincipalAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationMixinProps",
		reflect.TypeOf((*CfnDomainForOrganizationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin",
		reflect.TypeOf((*CfnDomainForOrganizationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainForOrganizationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin.IdentityCenterConfigurationProperty",
		reflect.TypeOf((*CfnDomainForOrganizationPropsMixin_IdentityCenterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin.IdentityProviderConfigurationProperty",
		reflect.TypeOf((*CfnDomainForOrganizationPropsMixin_IdentityProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainPropsMixin.IdentityCenterConfigurationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdentityCenterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainPropsMixin.IdentityProviderConfigurationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdentityProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnInsightRuleMixinProps",
		reflect.TypeOf((*CfnInsightRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnInsightRulePropsMixin",
		reflect.TypeOf((*CfnInsightRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInsightRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnIntegrationMixinProps",
		reflect.TypeOf((*CfnIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnIntegrationPropsMixin",
		reflect.TypeOf((*CfnIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnIntegrationPropsMixin.ApiKeyCredentialProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_ApiKeyCredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnIntegrationPropsMixin.IntegrationCredentialProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_IntegrationCredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnIntegrationPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnIntegrationPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmMixinProps",
		reflect.TypeOf((*CfnLogAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin",
		reflect.TypeOf((*CfnLogAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLogAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin.ActionConfigurationProperty",
		reflect.TypeOf((*CfnLogAlarmPropsMixin_ActionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin.AlarmTagsPropagationProperty",
		reflect.TypeOf((*CfnLogAlarmPropsMixin_AlarmTagsPropagationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin.ScheduleConfigurationProperty",
		reflect.TypeOf((*CfnLogAlarmPropsMixin_ScheduleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin.ScheduledQueryConfigurationProperty",
		reflect.TypeOf((*CfnLogAlarmPropsMixin_ScheduledQueryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin.WarmUpConfigurationProperty",
		reflect.TypeOf((*CfnLogAlarmPropsMixin_WarmUpConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricDataQueryDefinitionMixinProps",
		reflect.TypeOf((*CfnMetricDataQueryDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricDataQueryDefinitionPropsMixin",
		reflect.TypeOf((*CfnMetricDataQueryDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricStreamMixinProps",
		reflect.TypeOf((*CfnMetricStreamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricStreamPropsMixin",
		reflect.TypeOf((*CfnMetricStreamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMetricStreamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricStreamPropsMixin.MetricStreamFilterProperty",
		reflect.TypeOf((*CfnMetricStreamPropsMixin_MetricStreamFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricStreamPropsMixin.MetricStreamStatisticsConfigurationProperty",
		reflect.TypeOf((*CfnMetricStreamPropsMixin_MetricStreamStatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricStreamPropsMixin.MetricStreamStatisticsMetricProperty",
		reflect.TypeOf((*CfnMetricStreamPropsMixin_MetricStreamStatisticsMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnOTelEnrichmentMixinProps",
		reflect.TypeOf((*CfnOTelEnrichmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnOTelEnrichmentPropsMixin",
		reflect.TypeOf((*CfnOTelEnrichmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOTelEnrichmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnOTelEnrichmentPropsMixin.OTelEnrichmentMetricSelectorProperty",
		reflect.TypeOf((*CfnOTelEnrichmentPropsMixin_OTelEnrichmentMetricSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnOmniDashboardMixinProps",
		reflect.TypeOf((*CfnOmniDashboardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnOmniDashboardPropsMixin",
		reflect.TypeOf((*CfnOmniDashboardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOmniDashboardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnResourceMetricsConfigurationMixinProps",
		reflect.TypeOf((*CfnResourceMetricsConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnResourceMetricsConfigurationPropsMixin",
		reflect.TypeOf((*CfnResourceMetricsConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceMetricsConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnResourceMetricsConfigurationPropsMixin.ResourceMetricSelectionProperty",
		reflect.TypeOf((*CfnResourceMetricsConfigurationPropsMixin_ResourceMetricSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnSpaceMixinProps",
		reflect.TypeOf((*CfnSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnSpacePropsMixin",
		reflect.TypeOf((*CfnSpacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSpacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnSpacePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSpacePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnViewMixinProps",
		reflect.TypeOf((*CfnViewMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnViewPropsMixin",
		reflect.TypeOf((*CfnViewPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnViewPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
