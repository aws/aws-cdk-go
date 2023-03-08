package awsnetworkfirewall


// The value to use in an Amazon CloudWatch custom metric dimension.
//
// This is used in the `PublishMetrics` custom action. A CloudWatch custom metric dimension is a name/value pair that's part of the identity of a metric.
//
// AWS Network Firewall sets the dimension name to `CustomAction` and you provide the dimension value.
//
// For more information about CloudWatch custom metric dimensions, see [Publishing Custom Metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html#usingDimensions) in the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/WhatIsCloudWatch.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionProperty := &dimensionProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnRuleGroup_DimensionProperty struct {
	// The value to use in the custom metric dimension.
	Value *string `field:"required" json:"value" yaml:"value"`
}

