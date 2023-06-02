package awsdlm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionArchiveTierProperty := &RetentionArchiveTierProperty{
//   	Count: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_RetentionArchiveTierProperty struct {
	// `CfnLifecyclePolicy.RetentionArchiveTierProperty.Count`.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// `CfnLifecyclePolicy.RetentionArchiveTierProperty.Interval`.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// `CfnLifecyclePolicy.RetentionArchiveTierProperty.IntervalUnit`.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

