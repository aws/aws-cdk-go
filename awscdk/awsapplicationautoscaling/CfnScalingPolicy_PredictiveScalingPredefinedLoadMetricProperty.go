package awsapplicationautoscaling


// Describes a load metric for a predictive scaling policy.
//
// When returned in the output of `DescribePolicies` , it indicates that a predictive scaling policy uses individually specified load and scaling metrics instead of a metric pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingPredefinedLoadMetricProperty := &PredictiveScalingPredefinedLoadMetricProperty{
//   	PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   	// the properties below are optional
//   	ResourceLabel: jsii.String("resourceLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedloadmetric.html
//
type CfnScalingPolicy_PredictiveScalingPredefinedLoadMetricProperty struct {
	// The metric type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedloadmetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedloadmetric-predefinedmetrictype
	//
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedloadmetric.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpredefinedloadmetric-resourcelabel
	//
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

