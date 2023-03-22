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
//   customLineItemChargeDetailsProperty := &customLineItemChargeDetailsProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	flat: &customLineItemFlatChargeDetailsProperty{
//   		chargeValue: jsii.Number(123),
//   	},
//   	percentage: &customLineItemPercentageChargeDetailsProperty{
//   		percentageValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		childAssociatedResources: []*string{
//   			jsii.String("childAssociatedResources"),
//   		},
//   	},
//   }
//
type CfnCustomLineItem_CustomLineItemChargeDetailsProperty struct {
	// The type of the custom line item that indicates whether the charge is a fee or credit.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A `CustomLineItemFlatChargeDetails` that describes the charge details of a flat custom line item.
	Flat interface{} `field:"optional" json:"flat" yaml:"flat"`
	// A `CustomLineItemPercentageChargeDetails` that describes the charge details of a percentage custom line item.
	Percentage interface{} `field:"optional" json:"percentage" yaml:"percentage"`
}

