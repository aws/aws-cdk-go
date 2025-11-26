package previewawsbillingconductormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCustomLineItemPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomLineItemMixinProps := &CfnCustomLineItemMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	BillingGroupArn: jsii.String("billingGroupArn"),
//   	BillingPeriodRange: &BillingPeriodRangeProperty{
//   		ExclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   		InclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   	},
//   	ComputationRule: jsii.String("computationRule"),
//   	CustomLineItemChargeDetails: &CustomLineItemChargeDetailsProperty{
//   		Flat: &CustomLineItemFlatChargeDetailsProperty{
//   			ChargeValue: jsii.Number(123),
//   		},
//   		LineItemFilters: []interface{}{
//   			&LineItemFilterProperty{
//   				Attribute: jsii.String("attribute"),
//   				MatchOption: jsii.String("matchOption"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Percentage: &CustomLineItemPercentageChargeDetailsProperty{
//   			ChildAssociatedResources: []*string{
//   				jsii.String("childAssociatedResources"),
//   			},
//   			PercentageValue: jsii.Number(123),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PresentationDetails: &PresentationDetailsProperty{
//   		Service: jsii.String("service"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html
//
type CfnCustomLineItemMixinProps struct {
	// The AWS account in which this custom line item will be applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-billinggrouparn
	//
	BillingGroupArn *string `field:"optional" json:"billingGroupArn" yaml:"billingGroupArn"`
	// A time range for which the custom line item is effective.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-billingperiodrange
	//
	BillingPeriodRange interface{} `field:"optional" json:"billingPeriodRange" yaml:"billingPeriodRange"`
	// The display settings of the Custom Line Item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-computationrule
	//
	ComputationRule *string `field:"optional" json:"computationRule" yaml:"computationRule"`
	// The charge details of a custom line item.
	//
	// It should contain only one of `Flat` or `Percentage` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-customlineitemchargedetails
	//
	CustomLineItemChargeDetails interface{} `field:"optional" json:"customLineItemChargeDetails" yaml:"customLineItemChargeDetails"`
	// The custom line item's description.
	//
	// This is shown on the Bills page in association with the charge value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The custom line item's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-presentationdetails
	//
	PresentationDetails interface{} `field:"optional" json:"presentationDetails" yaml:"presentationDetails"`
	// A map that contains tag keys and tag values that are attached to a custom line item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

