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
//   customLineItemFlatChargeDetailsProperty := &customLineItemFlatChargeDetailsProperty{
//   	chargeValue: jsii.Number(123),
//   }
//
type CfnCustomLineItem_CustomLineItemFlatChargeDetailsProperty struct {
	// The custom line item's fixed charge value in USD.
	ChargeValue *float64 `field:"required" json:"chargeValue" yaml:"chargeValue"`
}

