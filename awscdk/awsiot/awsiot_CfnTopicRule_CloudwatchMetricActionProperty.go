package awsiot


// Describes an action that captures a CloudWatch metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchMetricActionProperty := &cloudwatchMetricActionProperty{
//   	metricName: jsii.String("metricName"),
//   	metricNamespace: jsii.String("metricNamespace"),
//   	metricUnit: jsii.String("metricUnit"),
//   	metricValue: jsii.String("metricValue"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	metricTimestamp: jsii.String("metricTimestamp"),
//   }
//
type CfnTopicRule_CloudwatchMetricActionProperty struct {
	// The CloudWatch metric name.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The CloudWatch metric namespace name.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The [metric unit](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.
	MetricUnit *string `field:"required" json:"metricUnit" yaml:"metricUnit"`
	// The CloudWatch metric value.
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// The IAM role that allows access to the CloudWatch metric.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// An optional [Unix timestamp](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp) .
	MetricTimestamp *string `field:"optional" json:"metricTimestamp" yaml:"metricTimestamp"`
}

