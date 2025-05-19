package awsbillingconductor


// The billing period range in which the custom line item request will be applied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingPeriodRangeProperty := &BillingPeriodRangeProperty{
//   	ExclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   	InclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-billingperiodrange.html
//
type CfnCustomLineItem_BillingPeriodRangeProperty struct {
	// The exclusive end billing period that defines a billing period range where a custom line is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-billingperiodrange.html#cfn-billingconductor-customlineitem-billingperiodrange-exclusiveendbillingperiod
	//
	ExclusiveEndBillingPeriod *string `field:"optional" json:"exclusiveEndBillingPeriod" yaml:"exclusiveEndBillingPeriod"`
	// The inclusive start billing period that defines a billing period range where a custom line is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-billingperiodrange.html#cfn-billingconductor-customlineitem-billingperiodrange-inclusivestartbillingperiod
	//
	InclusiveStartBillingPeriod *string `field:"optional" json:"inclusiveStartBillingPeriod" yaml:"inclusiveStartBillingPeriod"`
}

