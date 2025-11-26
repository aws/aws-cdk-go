package previewawsbillingconductormixins


// The set of tiering configurations for the pricing rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tieringProperty := &TieringProperty{
//   	FreeTier: &FreeTierProperty{
//   		Activated: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-tiering.html
//
type CfnPricingRulePropsMixin_TieringProperty struct {
	// The possible AWS Free Tier configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-tiering.html#cfn-billingconductor-pricingrule-tiering-freetier
	//
	FreeTier interface{} `field:"optional" json:"freeTier" yaml:"freeTier"`
}

