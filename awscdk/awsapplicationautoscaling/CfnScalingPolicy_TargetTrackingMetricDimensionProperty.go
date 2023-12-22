package awsapplicationautoscaling


// `TargetTrackingMetricDimension` specifies a name/value pair that is part of the identity of a CloudWatch metric for the `Dimensions` property of the [AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingMetric](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetric.html) property type. Duplicate dimensions are not allowed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingMetricDimensionProperty := &TargetTrackingMetricDimensionProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricdimension.html
//
type CfnScalingPolicy_TargetTrackingMetricDimensionProperty struct {
	// The name of the dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricdimension.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricdimension-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricdimension.html#cfn-applicationautoscaling-scalingpolicy-targettrackingmetricdimension-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

