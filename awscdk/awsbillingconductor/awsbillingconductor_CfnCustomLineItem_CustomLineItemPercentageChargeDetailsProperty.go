package awsbillingconductor


// A representation of the charge details associated with a percentage custom line item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLineItemPercentageChargeDetailsProperty := &customLineItemPercentageChargeDetailsProperty{
//   	percentageValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	childAssociatedResources: []*string{
//   		jsii.String("childAssociatedResources"),
//   	},
//   }
//
type CfnCustomLineItem_CustomLineItemPercentageChargeDetailsProperty struct {
	// The custom line item's percentage value.
	//
	// This will be multiplied against the combined value of its associated resources to determine its charge value.
	PercentageValue *float64 `field:"required" json:"percentageValue" yaml:"percentageValue"`
	// A list of resource ARNs to associate to the percentage custom line item.
	ChildAssociatedResources *[]*string `field:"optional" json:"childAssociatedResources" yaml:"childAssociatedResources"`
}

