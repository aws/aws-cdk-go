package awsbillingconductor


// The charge details of a custom line item.
//
// It should contain only one of `Flat` or `Percentage` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLineItemFlatChargeDetailsProperty := &CustomLineItemFlatChargeDetailsProperty{
//   	ChargeValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemflatchargedetails.html
//
type CfnCustomLineItem_CustomLineItemFlatChargeDetailsProperty struct {
	// The custom line item's fixed charge value in USD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemflatchargedetails.html#cfn-billingconductor-customlineitem-customlineitemflatchargedetails-chargevalue
	//
	ChargeValue *float64 `field:"required" json:"chargeValue" yaml:"chargeValue"`
}

