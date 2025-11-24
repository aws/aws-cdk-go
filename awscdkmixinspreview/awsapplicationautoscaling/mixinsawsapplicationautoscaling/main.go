package mixinsawsapplicationautoscaling

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetMixinProps",
		reflect.TypeOf((*CfnScalableTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin",
		reflect.TypeOf((*CfnScalableTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScalableTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin.ScalableTargetActionProperty",
		reflect.TypeOf((*CfnScalableTargetPropsMixin_ScalableTargetActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin.ScheduledActionProperty",
		reflect.TypeOf((*CfnScalableTargetPropsMixin_ScheduledActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin.SuspendedStateProperty",
		reflect.TypeOf((*CfnScalableTargetPropsMixin_SuspendedStateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyMixinProps",
		reflect.TypeOf((*CfnScalingPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScalingPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.CustomizedMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_CustomizedMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredefinedMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredefinedMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingCustomizedCapacityMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedCapacityMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingCustomizedLoadMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedLoadMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingCustomizedScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingMetricDataQueryProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingMetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingMetricDimensionProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingMetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingMetricStatProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingMetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingPredefinedLoadMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedLoadMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingPredefinedMetricPairProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedMetricPairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.PredictiveScalingPredefinedScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.StepAdjustmentProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_StepAdjustmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.StepScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_StepScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.TargetTrackingMetricDataQueryProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingMetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.TargetTrackingMetricDimensionProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingMetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.TargetTrackingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.TargetTrackingMetricStatProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingMetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalingPolicyPropsMixin.TargetTrackingScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
}
