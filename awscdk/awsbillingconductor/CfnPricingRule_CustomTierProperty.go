package awsbillingconductor


// A custom volume tier that defines the rate applied to usage within the given range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customTierProperty := &CustomTierProperty{
//   	BeginRangeInclusive: jsii.Number(123),
//   	RateValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	EndRangeExclusive: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-customtier.html
//
type CfnPricingRule_CustomTierProperty struct {
	// The inclusive beginning of the tier's usage range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-customtier.html#cfn-billingconductor-pricingrule-customtier-beginrangeinclusive
	//
	BeginRangeInclusive *float64 `field:"required" json:"beginRangeInclusive" yaml:"beginRangeInclusive"`
	// The custom rate applied to usage within the tier's range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-customtier.html#cfn-billingconductor-pricingrule-customtier-ratevalue
	//
	RateValue *float64 `field:"required" json:"rateValue" yaml:"rateValue"`
	// The exclusive end of the tier's usage range.
	//
	// Omit for the last tier (infinity).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-customtier.html#cfn-billingconductor-pricingrule-customtier-endrangeexclusive
	//
	EndRangeExclusive *float64 `field:"optional" json:"endRangeExclusive" yaml:"endRangeExclusive"`
}

