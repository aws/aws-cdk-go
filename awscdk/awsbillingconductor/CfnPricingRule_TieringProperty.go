package awsbillingconductor


// The set of tiering configurations for the pricing rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tieringProperty := &TieringProperty{
//   	CustomTiers: []interface{}{
//   		&CustomTierProperty{
//   			BeginRangeInclusive: jsii.Number(123),
//   			RateValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			EndRangeExclusive: jsii.Number(123),
//   		},
//   	},
//   	FreeTier: &FreeTierProperty{
//   		Activated: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-tiering.html
//
type CfnPricingRule_TieringProperty struct {
	// The set of custom volume tiers for a SKU-scoped TIERING pricing rule.
	//
	// Tiers must start at 0, be contiguous, and the last tier must have no end range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-tiering.html#cfn-billingconductor-pricingrule-tiering-customtiers
	//
	CustomTiers interface{} `field:"optional" json:"customTiers" yaml:"customTiers"`
	// The possible AWS Free Tier configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-tiering.html#cfn-billingconductor-pricingrule-tiering-freetier
	//
	FreeTier interface{} `field:"optional" json:"freeTier" yaml:"freeTier"`
}

