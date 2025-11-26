package previewawsbillingconductormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPricingPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPricingPlanMixinProps := &CfnPricingPlanMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PricingRuleArns: []*string{
//   		jsii.String("pricingRuleArns"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingplan.html
//
type CfnPricingPlanMixinProps struct {
	// The pricing plan description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingplan.html#cfn-billingconductor-pricingplan-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of a pricing plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingplan.html#cfn-billingconductor-pricingplan-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The `PricingRuleArns` that are associated with the Pricing Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingplan.html#cfn-billingconductor-pricingplan-pricingrulearns
	//
	PricingRuleArns *[]*string `field:"optional" json:"pricingRuleArns" yaml:"pricingRuleArns"`
	// A map that contains tag keys and tag values that are attached to a pricing plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingplan.html#cfn-billingconductor-pricingplan-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

