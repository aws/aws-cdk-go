package awsbillingconductor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tieringProperty := &tieringProperty{
//   	freeTier: &freeTierProperty{
//   		activated: jsii.Boolean(false),
//   	},
//   }
//
type CfnPricingRule_TieringProperty struct {
	// `CfnPricingRule.TieringProperty.FreeTier`.
	FreeTier interface{} `field:"optional" json:"freeTier" yaml:"freeTier"`
}

