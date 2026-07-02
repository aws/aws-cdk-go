package awswafv2


// A price configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   priceProperty := &PriceProperty{
//   	Amount: jsii.String("amount"),
//   	Currency: jsii.String("currency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-price.html
//
type CfnWebACL_PriceProperty struct {
	// The price amount.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-price.html#cfn-wafv2-webacl-price-amount
	//
	Amount *string `field:"required" json:"amount" yaml:"amount"`
	// The cryptocurrency to use for payment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-price.html#cfn-wafv2-webacl-price-currency
	//
	Currency *string `field:"required" json:"currency" yaml:"currency"`
}

