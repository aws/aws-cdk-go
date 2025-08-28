package awsiot


// Describes an action that captures a CloudWatch metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchMetricActionProperty := &CloudwatchMetricActionProperty{
//   	MetricName: jsii.String("metricName"),
//   	MetricNamespace: jsii.String("metricNamespace"),
//   	MetricUnit: jsii.String("metricUnit"),
//   	MetricValue: jsii.String("metricValue"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	MetricTimestamp: jsii.String("metricTimestamp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html
//
type CfnTopicRule_CloudwatchMetricActionProperty struct {
	// The CloudWatch metric name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The CloudWatch metric namespace name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricnamespace
	//
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The [metric unit](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricunit
	//
	MetricUnit *string `field:"required" json:"metricUnit" yaml:"metricUnit"`
	// The CloudWatch metric value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricvalue
	//
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// The IAM role that allows access to the CloudWatch metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// An optional [Unix timestamp](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metrictimestamp
	//
	MetricTimestamp *string `field:"optional" json:"metricTimestamp" yaml:"metricTimestamp"`
}

