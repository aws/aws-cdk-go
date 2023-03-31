package awsrum


// Specifies one custom metric or extended metric that you want the CloudWatch RUM app monitor to send to a destination.
//
// Valid destinations include CloudWatch and Evidently.
//
// By default, RUM app monitors send some metrics to CloudWatch . These default metrics are listed in [CloudWatch metrics that you can collect.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-metrics.html)
//
// In addition to these default metrics, you can choose to send extended metrics or custom metrics or both.
//
// - Extended metrics enable you to send metrics with additional dimensions not included in the default metrics. You can also send extended metrics to Evidently as well as CloudWatch . The valid dimension names for the additional dimensions for extended metrics are `BrowserName` , `CountryCode` , `DeviceType` , `FileType` , `OSName` , and `PageId` . For more information, see [Extended metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html) .
// - Custom metrics are metrics that you define. You can send custom metrics to CloudWatch or to CloudWatch Evidently or to both. With custom metrics, you can use any metric name and namespace, and to derive the metrics you can use any custom events, built-in events, custom attributes, or default attributes.
//
// You can't send custom metrics to the `AWS/RUM` namespace. You must send custom metrics to a custom namespace that you define. The namespace that you use can't start with `AWS/` . CloudWatch RUM prepends `RUM/CustomMetrics/` to the custom namespace that you define, so the final namespace for your metrics in CloudWatch is `RUM/CustomMetrics/ *your-custom-namespace*` .
//
// For information about syntax rules for specifying custom metrics and extended metrics, see [MetridDefinitionRequest](https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_MetricDefinitionRequest.html) in the *CloudWatch RUM API Reference* .
//
// The maximum number of metric definitions that one destination can contain is 2000.
//
// Extended metrics sent to CloudWatch and RUM custom metrics are charged as CloudWatch custom metrics. Each combination of additional dimension name and dimension value counts as a custom metric.
//
// If some metric definitions that you specify are not valid, then the operation will not modify any metric definitions even if other metric definitions specified are valid.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinitionProperty := &metricDefinitionProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dimensionKeys: map[string]*string{
//   		"dimensionKeysKey": jsii.String("dimensionKeys"),
//   	},
//   	eventPattern: jsii.String("eventPattern"),
//   	namespace: jsii.String("namespace"),
//   	unitLabel: jsii.String("unitLabel"),
//   	valueKey: jsii.String("valueKey"),
//   }
//
type CfnAppMonitor_MetricDefinitionProperty struct {
	// The name of the metric that is defined in this structure.
	Name *string `field:"required" json:"name" yaml:"name"`
	// This field is a map of field paths to dimension names.
	//
	// It defines the dimensions to associate with this metric in CloudWatch . The value of this field is used only if the metric destination is `CloudWatch` . If the metric destination is `Evidently` , the value of `DimensionKeys` is ignored.
	DimensionKeys interface{} `field:"optional" json:"dimensionKeys" yaml:"dimensionKeys"`
	// The pattern that defines the metric.
	//
	// RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.
	//
	// If the metrics destination is `CloudWatch` and the event also matches a value in `DimensionKeys` , then the metric is published with the specified dimensions.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// If you are creating a custom metric instead of an extended metrics, use this parameter to define the metric namespace for that custom metric.
	//
	// Do not specify this parameter if you are creating an extended metric.
	//
	// You can't use any string that starts with `AWS/` for your namespace.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Use this field only if you are sending this metric to CloudWatch .
	//
	// It defines the CloudWatch metric unit that this metric is measured in.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
	// The field within the event object that the metric value is sourced from.
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

