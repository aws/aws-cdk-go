package mixinsawsapplicationsignals

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnDiscoveryMixinProps",
		reflect.TypeOf((*CfnDiscoveryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnDiscoveryPropsMixin",
		reflect.TypeOf((*CfnDiscoveryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDiscoveryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnGroupingConfigurationMixinProps",
		reflect.TypeOf((*CfnGroupingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnGroupingConfigurationPropsMixin",
		reflect.TypeOf((*CfnGroupingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGroupingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnGroupingConfigurationPropsMixin.GroupingAttributeDefinitionProperty",
		reflect.TypeOf((*CfnGroupingConfigurationPropsMixin_GroupingAttributeDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectiveMixinProps",
		reflect.TypeOf((*CfnServiceLevelObjectiveMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServiceLevelObjectivePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.BurnRateConfigurationProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_BurnRateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.CalendarIntervalProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_CalendarIntervalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.DependencyConfigProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_DependencyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.ExclusionWindowProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_ExclusionWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.GoalProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_GoalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.IntervalProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_IntervalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.MetricDataQueryProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.MetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.MetricStatProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.MonitoredRequestCountMetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_MonitoredRequestCountMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.RecurrenceRuleProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_RecurrenceRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.RequestBasedSliMetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_RequestBasedSliMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.RequestBasedSliProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_RequestBasedSliProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.RollingIntervalProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_RollingIntervalProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.SliMetricProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_SliMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.SliProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_SliProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationsignals.mixins.CfnServiceLevelObjectivePropsMixin.WindowProperty",
		reflect.TypeOf((*CfnServiceLevelObjectivePropsMixin_WindowProperty)(nil)).Elem(),
	)
}
