package awsdlm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyDeprecateRuleProperty := &CrossRegionCopyDeprecateRuleProperty{
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyDeprecateRuleProperty struct {
	// `CfnLifecyclePolicy.CrossRegionCopyDeprecateRuleProperty.Interval`.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// `CfnLifecyclePolicy.CrossRegionCopyDeprecateRuleProperty.IntervalUnit`.
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

