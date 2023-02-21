package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCustomMetric`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomMetricProps := &cfnCustomMetricProps{
//   	metricType: jsii.String("metricType"),
//
//   	// the properties below are optional
//   	displayName: jsii.String("displayName"),
//   	metricName: jsii.String("metricName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomMetricProps struct {
	// The type of the custom metric. Types include `string-list` , `ip-address-list` , `number-list` , and `number` .
	//
	// > The type `number` only takes a single metric value as an input, but when you submit the metrics value in the DeviceMetrics report, you must pass it as an array with a single value.
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// The friendly name in the console for the custom metric.
	//
	// This name doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. You can update the friendly name after you define it.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the custom metric.
	//
	// This will be used in the metric report submitted from the device/thing. The name can't begin with `aws:` . You can’t change the name after you define it.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Metadata that can be used to manage the custom metric.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

