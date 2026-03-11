package awsautoscaling

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupMixinProps",
		reflect.TypeOf((*CfnAutoScalingGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutoScalingGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.AcceleratorCountRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_AcceleratorCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.AcceleratorTotalMemoryMiBRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_AcceleratorTotalMemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.AvailabilityZoneDistributionProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_AvailabilityZoneDistributionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.AvailabilityZoneImpairmentPolicyProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_AvailabilityZoneImpairmentPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.BaselineEbsBandwidthMbpsRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_BaselineEbsBandwidthMbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.BaselinePerformanceFactorsRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_BaselinePerformanceFactorsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.CapacityReservationSpecificationProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_CapacityReservationSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.CapacityReservationTargetProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_CapacityReservationTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.CpuPerformanceFactorRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_CpuPerformanceFactorRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.InstanceLifecyclePolicyProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_InstanceLifecyclePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.InstanceMaintenancePolicyProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_InstanceMaintenancePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.InstanceRequirementsProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_InstanceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.InstancesDistributionProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_InstancesDistributionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.LaunchTemplateOverridesProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_LaunchTemplateOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.LaunchTemplateProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_LaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.LifecycleHookSpecificationProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_LifecycleHookSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.MemoryGiBPerVCpuRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_MemoryGiBPerVCpuRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.MemoryMiBRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_MemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.MetricsCollectionProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_MetricsCollectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.MixedInstancesPolicyProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_MixedInstancesPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.NetworkBandwidthGbpsRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_NetworkBandwidthGbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.NetworkInterfaceCountRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_NetworkInterfaceCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.PerformanceFactorReferenceRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_PerformanceFactorReferenceRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.RetentionTriggersProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_RetentionTriggersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.TagPropertyProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_TagPropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.TotalLocalStorageGBRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_TotalLocalStorageGBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.TrafficSourceIdentifierProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_TrafficSourceIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnAutoScalingGroupPropsMixin.VCpuCountRequestProperty",
		reflect.TypeOf((*CfnAutoScalingGroupPropsMixin_VCpuCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLaunchConfigurationMixinProps",
		reflect.TypeOf((*CfnLaunchConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLaunchConfigurationPropsMixin",
		reflect.TypeOf((*CfnLaunchConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLaunchConfigurationPropsMixin.BlockDeviceMappingProperty",
		reflect.TypeOf((*CfnLaunchConfigurationPropsMixin_BlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLaunchConfigurationPropsMixin.BlockDeviceProperty",
		reflect.TypeOf((*CfnLaunchConfigurationPropsMixin_BlockDeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLaunchConfigurationPropsMixin.MetadataOptionsProperty",
		reflect.TypeOf((*CfnLaunchConfigurationPropsMixin_MetadataOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLifecycleHookMixinProps",
		reflect.TypeOf((*CfnLifecycleHookMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnLifecycleHookPropsMixin",
		reflect.TypeOf((*CfnLifecycleHookPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLifecycleHookPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyMixinProps",
		reflect.TypeOf((*CfnScalingPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScalingPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.CustomizedMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_CustomizedMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.MetricDataQueryProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.MetricStatProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredefinedMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredefinedMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingCustomizedCapacityMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedCapacityMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingCustomizedLoadMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedLoadMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingCustomizedScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingCustomizedScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingMetricSpecificationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingPredefinedLoadMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedLoadMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingPredefinedMetricPairProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedMetricPairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.PredictiveScalingPredefinedScalingMetricProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedScalingMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.StepAdjustmentProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_StepAdjustmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.TargetTrackingConfigurationProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.TargetTrackingMetricDataQueryProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingMetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScalingPolicyPropsMixin.TargetTrackingMetricStatProperty",
		reflect.TypeOf((*CfnScalingPolicyPropsMixin_TargetTrackingMetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScheduledActionMixinProps",
		reflect.TypeOf((*CfnScheduledActionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnScheduledActionPropsMixin",
		reflect.TypeOf((*CfnScheduledActionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScheduledActionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnWarmPoolMixinProps",
		reflect.TypeOf((*CfnWarmPoolMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnWarmPoolPropsMixin",
		reflect.TypeOf((*CfnWarmPoolPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWarmPoolPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_autoscaling.CfnWarmPoolPropsMixin.InstanceReusePolicyProperty",
		reflect.TypeOf((*CfnWarmPoolPropsMixin_InstanceReusePolicyProperty)(nil)).Elem(),
	)
}
