package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnBillingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBillingGroupProps := &cfnBillingGroupProps{
//   	accountGrouping: &accountGroupingProperty{
//   		linkedAccountIds: []*string{
//   			jsii.String("linkedAccountIds"),
//   		},
//   	},
//   	computationPreference: &computationPreferenceProperty{
//   		pricingPlanArn: jsii.String("pricingPlanArn"),
//   	},
//   	name: jsii.String("name"),
//   	primaryAccountId: jsii.String("primaryAccountId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBillingGroupProps struct {
	// The set of accounts that will be under the billing group.
	//
	// The set of accounts resemble the linked accounts in a consolidated family.
	AccountGrouping interface{} `field:"required" json:"accountGrouping" yaml:"accountGrouping"`
	// The preferences and settings that will be used to compute the AWS charges for a billing group.
	ComputationPreference interface{} `field:"required" json:"computationPreference" yaml:"computationPreference"`
	// The billing group's name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account ID that serves as the main account in a billing group.
	PrimaryAccountId *string `field:"required" json:"primaryAccountId" yaml:"primaryAccountId"`
	// The billing group description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::BillingConductor::BillingGroup.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

