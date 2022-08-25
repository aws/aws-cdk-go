package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPricingPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPricingPlanProps := &cfnPricingPlanProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	pricingRuleArns: []*string{
//   		jsii.String("pricingRuleArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPricingPlanProps struct {
	// The name of a pricing plan.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The pricing plan description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The `PricingRuleArns` that are associated with the Pricing Plan.
	PricingRuleArns *[]*string `field:"optional" json:"pricingRuleArns" yaml:"pricingRuleArns"`
	// `AWS::BillingConductor::PricingPlan.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

