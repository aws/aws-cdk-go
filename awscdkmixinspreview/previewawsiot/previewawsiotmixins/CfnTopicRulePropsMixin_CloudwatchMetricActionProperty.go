package previewawsiotmixins


// Describes an action that captures a CloudWatch metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudwatchMetricActionProperty := &CloudwatchMetricActionProperty{
//   	MetricName: jsii.String("metricName"),
//   	MetricNamespace: jsii.String("metricNamespace"),
//   	MetricTimestamp: jsii.String("metricTimestamp"),
//   	MetricUnit: jsii.String("metricUnit"),
//   	MetricValue: jsii.String("metricValue"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html
//
type CfnTopicRulePropsMixin_CloudwatchMetricActionProperty struct {
	// The CloudWatch metric name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The CloudWatch metric namespace name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricnamespace
	//
	MetricNamespace *string `field:"optional" json:"metricNamespace" yaml:"metricNamespace"`
	// An optional [Unix timestamp](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metrictimestamp
	//
	MetricTimestamp *string `field:"optional" json:"metricTimestamp" yaml:"metricTimestamp"`
	// The [metric unit](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricunit
	//
	MetricUnit *string `field:"optional" json:"metricUnit" yaml:"metricUnit"`
	// The CloudWatch metric value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-metricvalue
	//
	MetricValue *string `field:"optional" json:"metricValue" yaml:"metricValue"`
	// The IAM role that allows access to the CloudWatch metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchmetricaction.html#cfn-iot-topicrule-cloudwatchmetricaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

