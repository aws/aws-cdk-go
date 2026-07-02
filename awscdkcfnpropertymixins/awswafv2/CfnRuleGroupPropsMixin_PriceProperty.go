package awswafv2


// A price configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   priceProperty := &PriceProperty{
//   	Amount: jsii.String("amount"),
//   	Currency: jsii.String("currency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-price.html
//
type CfnRuleGroupPropsMixin_PriceProperty struct {
	// The price amount.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-price.html#cfn-wafv2-rulegroup-price-amount
	//
	Amount *string `field:"optional" json:"amount" yaml:"amount"`
	// The cryptocurrency to use for payment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-price.html#cfn-wafv2-rulegroup-price-currency
	//
	Currency *string `field:"optional" json:"currency" yaml:"currency"`
}

