package mixinsawsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPricingRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPricingRuleMixinProps := &CfnPricingRuleMixinProps{
//   	BillingEntity: jsii.String("billingEntity"),
//   	Description: jsii.String("description"),
//   	ModifierPercentage: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Operation: jsii.String("operation"),
//   	Scope: jsii.String("scope"),
//   	Service: jsii.String("service"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tiering: &TieringProperty{
//   		FreeTier: &FreeTierProperty{
//   			Activated: jsii.Boolean(false),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	UsageType: jsii.String("usageType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html
//
type CfnPricingRuleMixinProps struct {
	// The seller of services provided by AWS , their affiliates, or third-party providers selling services via AWS Marketplace .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-billingentity
	//
	BillingEntity *string `field:"optional" json:"billingEntity" yaml:"billingEntity"`
	// The pricing rule description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A percentage modifier applied on the public pricing rates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-modifierpercentage
	//
	ModifierPercentage *float64 `field:"optional" json:"modifierPercentage" yaml:"modifierPercentage"`
	// The name of a pricing rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Operation is the specific AWS action covered by this line item.
	//
	// This describes the specific usage of the line item.
	//
	// If the `Scope` attribute is set to `SKU` , this attribute indicates which operation the `PricingRule` is modifying. For example, a value of `RunInstances:0202` indicates the operation of running an Amazon EC2 instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-operation
	//
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// The scope of pricing rule that indicates if it's globally applicable or service-specific.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// If the `Scope` attribute is `SERVICE` , this attribute indicates which service the `PricingRule` is applicable for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
	// A map that contains tag keys and tag values that are attached to a pricing rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The set of tiering configurations for the pricing rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-tiering
	//
	Tiering interface{} `field:"optional" json:"tiering" yaml:"tiering"`
	// The type of pricing rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Usage Type is the unit that each service uses to measure the usage of a specific type of resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-pricingrule.html#cfn-billingconductor-pricingrule-usagetype
	//
	UsageType *string `field:"optional" json:"usageType" yaml:"usageType"`
}

