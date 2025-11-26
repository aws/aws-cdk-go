package previewawsapplicationautoscalingmixins


// Represents a metric pair for a predictive scaling policy.
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
//   predictiveScalingPredefinedMetricPairProperty := &PredictiveScalingPredefinedMetricPairProperty{
//   	PredefinedMetricType: jsii.String("predefinedMetricType"),
//   	ResourceLabel: jsii.String("resourceLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.html
//
type CfnScalingPolicyPropsMixin_PredictiveScalingPredefinedMetricPairProperty struct {
	// Indicates which metrics to use.
	//
	// There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair-predefinedmetrictype
	//
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a specific target group from which to determine the total and average request count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair-resourcelabel
	//
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

