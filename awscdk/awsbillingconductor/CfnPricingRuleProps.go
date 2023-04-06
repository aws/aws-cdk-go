package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPricingRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPricingRuleProps := &CfnPricingRuleProps{
//   	Name: jsii.String("name"),
//   	Scope: jsii.String("scope"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	BillingEntity: jsii.String("billingEntity"),
//   	Description: jsii.String("description"),
//   	ModifierPercentage: jsii.Number(123),
//   	Operation: jsii.String("operation"),
//   	Service: jsii.String("service"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tiering: &TieringProperty{
//   		FreeTier: &FreeTierProperty{
//   			Activated: jsii.Boolean(false),
//   		},
//   	},
//   	UsageType: jsii.String("usageType"),
//   }
//
type CfnPricingRuleProps struct {
	// The name of a pricing rule.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scope of pricing rule that indicates if it is globally applicable, or if it is service-specific.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The type of pricing rule.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The seller of services provided by AWS , their affiliates, or third-party providers selling services via AWS Marketplace .
	BillingEntity *string `field:"optional" json:"billingEntity" yaml:"billingEntity"`
	// The pricing rule description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A percentage modifier applied on the public pricing rates.
	ModifierPercentage *float64 `field:"optional" json:"modifierPercentage" yaml:"modifierPercentage"`
	// `AWS::BillingConductor::PricingRule.Operation`.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// If the `Scope` attribute is `SERVICE` , this attribute indicates which service the `PricingRule` is applicable for.
	Service *string `field:"optional" json:"service" yaml:"service"`
	// A map that contains tag keys and tag values that are attached to a pricing rule.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The set of tiering configurations for the pricing rule.
	Tiering interface{} `field:"optional" json:"tiering" yaml:"tiering"`
	// Usage Type is the unit that each service uses to measure the usage of a specific type of resource.
	UsageType *string `field:"optional" json:"usageType" yaml:"usageType"`
}

