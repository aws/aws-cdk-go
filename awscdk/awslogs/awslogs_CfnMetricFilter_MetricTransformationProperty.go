package awslogs


// `MetricTransformation` is a property of the `AWS::Logs::MetricFilter` resource that describes how to transform log streams into a CloudWatch metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricTransformationProperty := &metricTransformationProperty{
//   	metricName: jsii.String("metricName"),
//   	metricNamespace: jsii.String("metricNamespace"),
//   	metricValue: jsii.String("metricValue"),
//
//   	// the properties below are optional
//   	defaultValue: jsii.Number(123),
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	unit: jsii.String("unit"),
//   }
//
type CfnMetricFilter_MetricTransformationProperty struct {
	// The name of the CloudWatch metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A custom namespace to contain your metric in CloudWatch.
	//
	// Use namespaces to group together metrics that are similar. For more information, see [Namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Namespace) .
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value that is published to the CloudWatch metric.
	//
	// For example, if you're counting the occurrences of a particular term like `Error` , specify 1 for the metric value. If you're counting the number of bytes transferred, reference the value that is in the log event by using $ followed by the name of the field that you specified in the filter pattern, such as `$.size` .
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// (Optional) The value to emit when a filter pattern does not match a log event.
	//
	// This value can be null.
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// `CfnMetricFilter.MetricTransformationProperty.Dimensions`.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// `CfnMetricFilter.MetricTransformationProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

