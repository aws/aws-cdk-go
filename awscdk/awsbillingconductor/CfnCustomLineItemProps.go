package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomLineItem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomLineItemProps := &CfnCustomLineItemProps{
//   	BillingGroupArn: jsii.String("billingGroupArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	BillingPeriodRange: &BillingPeriodRangeProperty{
//   		ExclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   		InclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   	},
//   	CustomLineItemChargeDetails: &CustomLineItemChargeDetailsProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
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
//   			PercentageValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			ChildAssociatedResources: []*string{
//   				jsii.String("childAssociatedResources"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html
//
type CfnCustomLineItemProps struct {
	// The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-billinggrouparn
	//
	BillingGroupArn *string `field:"required" json:"billingGroupArn" yaml:"billingGroupArn"`
	// The custom line item's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A time range for which the custom line item is effective.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-billingperiodrange
	//
	BillingPeriodRange interface{} `field:"optional" json:"billingPeriodRange" yaml:"billingPeriodRange"`
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
	// A map that contains tag keys and tag values that are attached to a custom line item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html#cfn-billingconductor-customlineitem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

