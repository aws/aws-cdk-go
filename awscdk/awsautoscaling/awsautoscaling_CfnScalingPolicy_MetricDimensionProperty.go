package awsautoscaling


// `MetricDimension` specifies a name/value pair that is part of the identity of a CloudWatch metric for the `Dimensions` property of the [AWS::AutoScaling::ScalingPolicy CustomizedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html) property type. Duplicate dimensions are not allowed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnScalingPolicy_MetricDimensionProperty struct {
	// The name of the dimension.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the dimension.
	Value *string `field:"required" json:"value" yaml:"value"`
}

