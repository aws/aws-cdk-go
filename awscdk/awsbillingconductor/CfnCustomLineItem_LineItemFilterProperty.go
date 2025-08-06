package awsbillingconductor


// A representation of the line item filter for your custom line item.
//
// You can use line item filters to include or exclude specific resource values from the billing group's total cost. For example, if you create a custom line item and you want to filter out a value, such as Savings Plan discounts, you can update `LineItemFilter` to exclude it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineItemFilterProperty := &LineItemFilterProperty{
//   	Attribute: jsii.String("attribute"),
//   	MatchOption: jsii.String("matchOption"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html
//
type CfnCustomLineItem_LineItemFilterProperty struct {
	// The attribute of the line item filter.
	//
	// This specifies what attribute that you can filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html#cfn-billingconductor-customlineitem-lineitemfilter-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The match criteria of the line item filter.
	//
	// This parameter specifies whether not to include the resource value from the billing group total cost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html#cfn-billingconductor-customlineitem-lineitemfilter-matchoption
	//
	MatchOption *string `field:"required" json:"matchOption" yaml:"matchOption"`
	// The values of the line item filter.
	//
	// This specifies the values to filter on. Currently, you can only exclude Savings Plan discounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html#cfn-billingconductor-customlineitem-lineitemfilter-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

