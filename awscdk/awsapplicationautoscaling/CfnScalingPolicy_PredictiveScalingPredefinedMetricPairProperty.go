package awsapplicationautoscaling


// Represents a metric pair for a predictive scaling policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingPredefinedMetricPairProperty := &PredictiveScalingPredefinedMetricPairProperty{
//   	PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   	// the properties below are optional
//   	ResourceLabel: jsii.String("resourceLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.html
//
type CfnScalingPolicy_PredictiveScalingPredefinedMetricPairProperty struct {
	// Indicates which metrics to use.
	//
	// There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair-predefinedmetrictype
	//
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a specific target group from which to determine the total and average request count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair-resourcelabel
	//
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

