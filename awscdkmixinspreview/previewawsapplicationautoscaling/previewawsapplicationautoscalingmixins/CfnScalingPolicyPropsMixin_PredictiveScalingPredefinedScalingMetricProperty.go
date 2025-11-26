package previewawsapplicationautoscalingmixins


// Describes a scaling metric for a predictive scaling policy.
//
// When returned in the output of `DescribePolicies` , it indicates that a predictive scaling policy uses individually specified load and scaling metrics instead of a metric pair.
//
// The following predefined metrics are available for predictive scaling:
//
// - `ECSServiceAverageCPUUtilization`
// - `ECSServiceAverageMemoryUtilization`
// - `ECSServiceCPUUtilization`
// - `ECSServiceMemoryUtilization`
// - `ECSServiceTotalCPUUtilization`
// - `ECSServiceTotalMemoryUtilization`
// - `ALBRequestCount`
// - `ALBRequestCountPerTarget`
// - `TotalALBRequestCount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predictiveScalingPredefinedScalingMetricProperty := &PredictiveScalingPredefinedScalingMetricProperty{
//   	PredefinedMetricType: jsii.String("predefinedMetricType"),
//   	ResourceLabel: jsii.String("resourceLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric.html
//
type CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedScalingMetricProperty struct {
	// The metric type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric-predefinedmetrictype
	//
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a specific target group from which to determine the average request count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric-resourcelabel
	//
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

