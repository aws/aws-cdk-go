package previewawscloudwatchmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmMixinProps",
		reflect.TypeOf((*CfnAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin",
		reflect.TypeOf((*CfnAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin.MetricDataQueryProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin.MetricStatProperty",
		reflect.TypeOf((*CfnAlarmPropsMixin_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.DimensionProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.MetricCharacteristicsProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricCharacteristicsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.MetricDataQueryProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.MetricMathAnomalyDetectorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricMathAnomalyDetectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.MetricStatProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.RangeProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin.SingleMetricAnomalyDetectorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_SingleMetricAnomalyDetectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnCompositeAlarmMixinProps",
		reflect.TypeOf((*CfnCompositeAlarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnCompositeAlarmPropsMixin",
		reflect.TypeOf((*CfnCompositeAlarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCompositeAlarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnDashboardMixinProps",
		reflect.TypeOf((*CfnDashboardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnDashboardPropsMixin",
		reflect.TypeOf((*CfnDashboardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnInsightRuleMixinProps",
		reflect.TypeOf((*CfnInsightRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnInsightRulePropsMixin",
		reflect.TypeOf((*CfnInsightRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInsightRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamMixinProps",
		reflect.TypeOf((*CfnMetricStreamMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin",
		reflect.TypeOf((*CfnMetricStreamPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMetricStreamPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin.MetricStreamFilterProperty",
		reflect.TypeOf((*CfnMetricStreamPropsMixin_MetricStreamFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin.MetricStreamStatisticsConfigurationProperty",
		reflect.TypeOf((*CfnMetricStreamPropsMixin_MetricStreamStatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin.MetricStreamStatisticsMetricProperty",
		reflect.TypeOf((*CfnMetricStreamPropsMixin_MetricStreamStatisticsMetricProperty)(nil)).Elem(),
	)
}
