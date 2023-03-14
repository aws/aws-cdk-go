package awsautoscalingplans


// `MetricDimension` is a subproperty of [CustomizedScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-customizedscalingmetricspecification.html) that specifies a dimension for a customized metric to use with AWS Auto Scaling ( Auto Scaling Plans ). Dimensions are arbitrary name/value pairs that can be associated with a CloudWatch metric. Duplicate dimensions are not allowed.
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
type CfnScalingPlan_MetricDimensionProperty struct {
	// The name of the dimension.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the dimension.
	Value *string `field:"required" json:"value" yaml:"value"`
}

