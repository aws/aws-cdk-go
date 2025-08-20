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
//   customLineItemChargeDetailsProperty := &CustomLineItemChargeDetailsProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Flat: &CustomLineItemFlatChargeDetailsProperty{
//   		ChargeValue: jsii.Number(123),
//   	},
//   	LineItemFilters: []interface{}{
//   		&LineItemFilterProperty{
//   			Attribute: jsii.String("attribute"),
//   			MatchOption: jsii.String("matchOption"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Percentage: &CustomLineItemPercentageChargeDetailsProperty{
//   		PercentageValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		ChildAssociatedResources: []*string{
//   			jsii.String("childAssociatedResources"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html
//
type CfnCustomLineItem_CustomLineItemChargeDetailsProperty struct {
	// The type of the custom line item that indicates whether the charge is a fee or credit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A `CustomLineItemFlatChargeDetails` that describes the charge details of a flat custom line item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-flat
	//
	Flat interface{} `field:"optional" json:"flat" yaml:"flat"`
	// A representation of the line item filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-lineitemfilters
	//
	LineItemFilters interface{} `field:"optional" json:"lineItemFilters" yaml:"lineItemFilters"`
	// A `CustomLineItemPercentageChargeDetails` that describes the charge details of a percentage custom line item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html#cfn-billingconductor-customlineitem-customlineitemchargedetails-percentage
	//
	Percentage interface{} `field:"optional" json:"percentage" yaml:"percentage"`
}

