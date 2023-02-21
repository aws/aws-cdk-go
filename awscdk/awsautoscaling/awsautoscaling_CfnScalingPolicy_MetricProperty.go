package awsautoscaling


// Represents a specific metric.
//
// `Metric` is a property of the [AWS::AutoScaling::ScalingPolicy MetricStat](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricstat.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricProperty := &metricProperty{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	dimensions: []interface{}{
//   		&metricDimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnScalingPolicy_MetricProperty struct {
	// The name of the metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	//
	// For more information, see the table in [AWS services that publish CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the *Amazon CloudWatch User Guide* .
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The dimensions for the metric.
	//
	// For the list of available dimensions, see the AWS documentation available from the table in [AWS services that publish CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the *Amazon CloudWatch User Guide* .
	//
	// Conditional: If you published your metric with dimensions, you must specify the same dimensions in your scaling policy.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}

