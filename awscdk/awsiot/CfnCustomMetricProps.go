package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomMetric`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomMetricProps := &CfnCustomMetricProps{
//   	MetricType: jsii.String("metricType"),
//
//   	// the properties below are optional
//   	DisplayName: jsii.String("displayName"),
//   	MetricName: jsii.String("metricName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-custommetric.html
//
type CfnCustomMetricProps struct {
	// The type of the custom metric. Types include `string-list` , `ip-address-list` , `number-list` , and `number` .
	//
	// > The type `number` only takes a single metric value as an input, but when you submit the metrics value in the DeviceMetrics report, you must pass it as an array with a single value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-custommetric.html#cfn-iot-custommetric-metrictype
	//
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// The friendly name in the console for the custom metric.
	//
	// This name doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. You can update the friendly name after you define it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-custommetric.html#cfn-iot-custommetric-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the custom metric.
	//
	// This will be used in the metric report submitted from the device/thing. The name can't begin with `aws:` . You can’t change the name after you define it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-custommetric.html#cfn-iot-custommetric-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Metadata that can be used to manage the custom metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-custommetric.html#cfn-iot-custommetric-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

