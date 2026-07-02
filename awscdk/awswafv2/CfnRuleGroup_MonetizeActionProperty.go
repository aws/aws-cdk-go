package awswafv2


// Monetize action for rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monetizeActionProperty := &MonetizeActionProperty{
//   	PriceMultiplier: jsii.String("priceMultiplier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-monetizeaction.html
//
type CfnRuleGroup_MonetizeActionProperty struct {
	// The price multiplier for the monetize action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-monetizeaction.html#cfn-wafv2-rulegroup-monetizeaction-pricemultiplier
	//
	PriceMultiplier *string `field:"optional" json:"priceMultiplier" yaml:"priceMultiplier"`
}

