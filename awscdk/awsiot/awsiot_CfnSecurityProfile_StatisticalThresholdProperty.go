package awsiot


// A statistical ranking (percentile) that indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticalThresholdProperty := &statisticalThresholdProperty{
//   	statistic: jsii.String("statistic"),
//   }
//
type CfnSecurityProfile_StatisticalThresholdProperty struct {
	// The percentile that resolves to a threshold value by which compliance with a behavior is determined.
	//
	// Metrics are collected over the specified period ( `durationSeconds` ) from all reporting devices in your account and statistical ranks are calculated. Then, the measurements from a device are collected over the same period. If the accumulated measurements from the device fall above or below ( `comparisonOperator` ) the value associated with the percentile specified, then the device is considered to be in compliance with the behavior, otherwise a violation occurs.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

