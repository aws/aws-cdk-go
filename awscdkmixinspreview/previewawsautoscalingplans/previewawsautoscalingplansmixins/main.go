package previewawsautoscalingplansmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanMixinProps",
		reflect.TypeOf((*CfnScalingPlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin",
		reflect.TypeOf((*CfnScalingPlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScalingPlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.ApplicationSourceProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_ApplicationSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.CustomizedLoadMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_CustomizedLoadMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.CustomizedScalingMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_CustomizedScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.PredefinedLoadMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_PredefinedLoadMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.PredefinedScalingMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_PredefinedScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.ScalingInstructionProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_ScalingInstructionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.TagFilterProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin.TargetTrackingConfigurationProperty",
		reflect.TypeOf((*CfnScalingPlanPropsMixin_TargetTrackingConfigurationProperty)(nil)).Elem(),
	)
}
