package awsdlm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deprecateRuleProperty := &DeprecateRuleProperty{
//   	Count: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_DeprecateRuleProperty struct {
	// `CfnLifecyclePolicy.DeprecateRuleProperty.Count`.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// `CfnLifecyclePolicy.DeprecateRuleProperty.Interval`.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// `CfnLifecyclePolicy.DeprecateRuleProperty.IntervalUnit`.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

