package previewawsinspectorv2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationMixinProps",
		reflect.TypeOf((*CfnCisScanConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCisScanConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin.CisTargetsProperty",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin_CisTargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin.DailyScheduleProperty",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin_DailyScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin.MonthlyScheduleProperty",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin_MonthlyScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin.TimeProperty",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin_TimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin.WeeklyScheduleProperty",
		reflect.TypeOf((*CfnCisScanConfigurationPropsMixin_WeeklyScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationMixinProps",
		reflect.TypeOf((*CfnCodeSecurityIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin",
		reflect.TypeOf((*CfnCodeSecurityIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeSecurityIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin.CreateDetailsProperty",
		reflect.TypeOf((*CfnCodeSecurityIntegrationPropsMixin_CreateDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin.CreateGitLabSelfManagedIntegrationDetailProperty",
		reflect.TypeOf((*CfnCodeSecurityIntegrationPropsMixin_CreateGitLabSelfManagedIntegrationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin.UpdateDetailsProperty",
		reflect.TypeOf((*CfnCodeSecurityIntegrationPropsMixin_UpdateDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin.UpdateGitHubIntegrationDetailProperty",
		reflect.TypeOf((*CfnCodeSecurityIntegrationPropsMixin_UpdateGitHubIntegrationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin.UpdateGitLabSelfManagedIntegrationDetailProperty",
		reflect.TypeOf((*CfnCodeSecurityIntegrationPropsMixin_UpdateGitLabSelfManagedIntegrationDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationMixinProps",
		reflect.TypeOf((*CfnCodeSecurityScanConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin",
		reflect.TypeOf((*CfnCodeSecurityScanConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeSecurityScanConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin.CodeSecurityScanConfigurationProperty",
		reflect.TypeOf((*CfnCodeSecurityScanConfigurationPropsMixin_CodeSecurityScanConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin.ContinuousIntegrationScanConfigurationProperty",
		reflect.TypeOf((*CfnCodeSecurityScanConfigurationPropsMixin_ContinuousIntegrationScanConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin.PeriodicScanConfigurationProperty",
		reflect.TypeOf((*CfnCodeSecurityScanConfigurationPropsMixin_PeriodicScanConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityScanConfigurationPropsMixin.ScopeSettingsProperty",
		reflect.TypeOf((*CfnCodeSecurityScanConfigurationPropsMixin_ScopeSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterMixinProps",
		reflect.TypeOf((*CfnFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin",
		reflect.TypeOf((*CfnFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.DateFilterProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_DateFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.FilterCriteriaProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_FilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.MapFilterProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_MapFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.NumberFilterProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_NumberFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.PackageFilterProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_PackageFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.PortRangeFilterProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_PortRangeFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnFilterPropsMixin.StringFilterProperty",
		reflect.TypeOf((*CfnFilterPropsMixin_StringFilterProperty)(nil)).Elem(),
	)
}
