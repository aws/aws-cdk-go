package awsbillingconductor


// The billing period range in which the custom line item request will be applied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingPeriodRangeProperty := &billingPeriodRangeProperty{
//   	exclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   	inclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   }
//
type CfnCustomLineItem_BillingPeriodRangeProperty struct {
	// The inclusive end billing period that defines a billing period range where a custom line is applied.
	ExclusiveEndBillingPeriod *string `field:"optional" json:"exclusiveEndBillingPeriod" yaml:"exclusiveEndBillingPeriod"`
	// The inclusive start billing period that defines a billing period range where a custom line is applied.
	InclusiveStartBillingPeriod *string `field:"optional" json:"inclusiveStartBillingPeriod" yaml:"inclusiveStartBillingPeriod"`
}

