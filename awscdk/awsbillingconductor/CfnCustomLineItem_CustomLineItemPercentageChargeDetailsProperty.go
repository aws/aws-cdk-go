package awsbillingconductor


// A representation of the charge details associated with a percentage custom line item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLineItemPercentageChargeDetailsProperty := &CustomLineItemPercentageChargeDetailsProperty{
//   	PercentageValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	ChildAssociatedResources: []*string{
//   		jsii.String("childAssociatedResources"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html
//
type CfnCustomLineItem_CustomLineItemPercentageChargeDetailsProperty struct {
	// The custom line item's percentage value.
	//
	// This will be multiplied against the combined value of its associated resources to determine its charge value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html#cfn-billingconductor-customlineitem-customlineitempercentagechargedetails-percentagevalue
	//
	PercentageValue *float64 `field:"required" json:"percentageValue" yaml:"percentageValue"`
	// A list of resource ARNs to associate to the percentage custom line item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html#cfn-billingconductor-customlineitem-customlineitempercentagechargedetails-childassociatedresources
	//
	ChildAssociatedResources *[]*string `field:"optional" json:"childAssociatedResources" yaml:"childAssociatedResources"`
}

