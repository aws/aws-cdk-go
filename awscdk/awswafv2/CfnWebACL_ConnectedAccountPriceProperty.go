package awswafv2


// A connected-account price configuration.
//
// The payment provider rejects amounts below $0.01 and does not support sub-cent amounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectedAccountPriceProperty := &ConnectedAccountPriceProperty{
//   	Amount: jsii.String("amount"),
//   	Currency: jsii.String("currency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-connectedaccountprice.html
//
type CfnWebACL_ConnectedAccountPriceProperty struct {
	// The price amount for a connected account.
	//
	// Whole cents only, from 0.01 up to the payment provider's maximum of 999999.99.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-connectedaccountprice.html#cfn-wafv2-webacl-connectedaccountprice-amount
	//
	Amount *string `field:"required" json:"amount" yaml:"amount"`
	// The cryptocurrency to use for payment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-connectedaccountprice.html#cfn-wafv2-webacl-connectedaccountprice-currency
	//
	Currency *string `field:"required" json:"currency" yaml:"currency"`
}

