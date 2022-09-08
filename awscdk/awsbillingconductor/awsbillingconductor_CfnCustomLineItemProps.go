package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCustomLineItem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomLineItemProps := &cfnCustomLineItemProps{
//   	billingGroupArn: jsii.String("billingGroupArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	billingPeriodRange: &billingPeriodRangeProperty{
//   		exclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   		inclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   	},
//   	customLineItemChargeDetails: &customLineItemChargeDetailsProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		flat: &customLineItemFlatChargeDetailsProperty{
//   			chargeValue: jsii.Number(123),
//   		},
//   		percentage: &customLineItemPercentageChargeDetailsProperty{
//   			percentageValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			childAssociatedResources: []*string{
//   				jsii.String("childAssociatedResources"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomLineItemProps struct {
	// The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.
	BillingGroupArn *string `field:"required" json:"billingGroupArn" yaml:"billingGroupArn"`
	// The custom line item's name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A time range for which the custom line item is effective.
	BillingPeriodRange interface{} `field:"optional" json:"billingPeriodRange" yaml:"billingPeriodRange"`
	// The charge details of a custom line item.
	//
	// It should contain only one of `Flat` or `Percentage` .
	CustomLineItemChargeDetails interface{} `field:"optional" json:"customLineItemChargeDetails" yaml:"customLineItemChargeDetails"`
	// The custom line item's description.
	//
	// This is shown on the Bills page in association with the charge value.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::BillingConductor::CustomLineItem.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

