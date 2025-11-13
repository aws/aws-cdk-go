package interfacesawsbillingconductor


// A reference to a PricingRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pricingRuleReference := &PricingRuleReference{
//   	PricingRuleArn: jsii.String("pricingRuleArn"),
//   }
//
type PricingRuleReference struct {
	// The Arn of the PricingRule resource.
	PricingRuleArn *string `field:"required" json:"pricingRuleArn" yaml:"pricingRuleArn"`
}

