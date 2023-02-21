package awsrum


// Specifies the extended metrics that you want the CloudWatch RUM app monitor to send to a destination.
//
// Valid destinations include CloudWatch and Evidently.
//
// By default, RUM app monitors send some metrics to CloudWatch . These default metrics are listed in [CloudWatch metrics that you can collect with CloudWatch RUM](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-metrics.html) .
//
// If you also send extended metrics, you can send metrics to Evidently as well as CloudWatch , and you can also optionally send the metrics with additional dimensions. The valid dimension names for the additional dimensions are `BrowserName` , `CountryCode` , `DeviceType` , `FileType` , `OSName` , and `PageId` . For more information, see [Extended metrics that you can send to CloudWatch and CloudWatch Evidently](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html) .
//
// The maximum number of metric definitions that one destination can contain is 2000.
//
// Extended metrics sent are charged as CloudWatch custom metrics. Each combination of additional dimension name and dimension value counts as a custom metric.
//
// If some metric definitions that you specify are not valid, then the operation will not modify any metric definitions even if other metric definitions specified are valid.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinitionProperty := &MetricDefinitionProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DimensionKeys: map[string]*string{
//   		"dimensionKeysKey": jsii.String("dimensionKeys"),
//   	},
//   	EventPattern: jsii.String("eventPattern"),
//   	UnitLabel: jsii.String("unitLabel"),
//   	ValueKey: jsii.String("valueKey"),
//   }
//
type CfnAppMonitor_MetricDefinitionProperty struct {
	// The name of the metric that is defined in this structure.
	Name *string `field:"required" json:"name" yaml:"name"`
	// This field is a map of field paths to dimension names.
	//
	// It defines the dimensions to associate with this metric in CloudWatch The value of this field is used only if the metric destination is `CloudWatch` . If the metric destination is `Evidently` , the value of `DimensionKeys` is ignored.
	DimensionKeys interface{} `field:"optional" json:"dimensionKeys" yaml:"dimensionKeys"`
	// The pattern that defines the metric.
	//
	// RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.
	//
	// If the metrics destination is `CloudWatch` and the event also matches a value in `DimensionKeys` , then the metric is published with the specified dimensions.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Use this field only if you are sending this metric to CloudWatch .
	//
	// It defines the CloudWatch metric unit that this metric is measured in.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
	// The field within the event object that the metric value is sourced from.
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

