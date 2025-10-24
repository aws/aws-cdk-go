package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBillingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBillingGroupProps := &CfnBillingGroupProps{
//   	AccountGrouping: &AccountGroupingProperty{
//   		LinkedAccountIds: []*string{
//   			jsii.String("linkedAccountIds"),
//   		},
//
//   		// the properties below are optional
//   		AutoAssociate: jsii.Boolean(false),
//   	},
//   	ComputationPreference: &ComputationPreferenceProperty{
//   		PricingPlanArn: jsii.String("pricingPlanArn"),
//   	},
//   	Name: jsii.String("name"),
//   	PrimaryAccountId: jsii.String("primaryAccountId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html
//
type CfnBillingGroupProps struct {
	// The set of accounts that will be under the billing group.
	//
	// The set of accounts resemble the linked accounts in a consolidated billing family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html#cfn-billingconductor-billinggroup-accountgrouping
	//
	AccountGrouping interface{} `field:"required" json:"accountGrouping" yaml:"accountGrouping"`
	// The preferences and settings that will be used to compute the AWS charges for a billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html#cfn-billingconductor-billinggroup-computationpreference
	//
	ComputationPreference interface{} `field:"required" json:"computationPreference" yaml:"computationPreference"`
	// The billing group's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html#cfn-billingconductor-billinggroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account ID that serves as the main account in a billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html#cfn-billingconductor-billinggroup-primaryaccountid
	//
	PrimaryAccountId *string `field:"required" json:"primaryAccountId" yaml:"primaryAccountId"`
	// The description of the billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html#cfn-billingconductor-billinggroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map that contains tag keys and tag values that are attached to a billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html#cfn-billingconductor-billinggroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

