package awswafv2


// Configuration for a single payment network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   paymentNetworkProperty := &PaymentNetworkProperty{
//   	Chain: jsii.String("chain"),
//   	Prices: []interface{}{
//   		&PriceProperty{
//   			Amount: jsii.String("amount"),
//   			Currency: jsii.String("currency"),
//   		},
//   	},
//   	WalletAddress: jsii.String("walletAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-paymentnetwork.html
//
type CfnWebACLPropsMixin_PaymentNetworkProperty struct {
	// The blockchain chain to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-paymentnetwork.html#cfn-wafv2-webacl-paymentnetwork-chain
	//
	Chain *string `field:"optional" json:"chain" yaml:"chain"`
	// List of price configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-paymentnetwork.html#cfn-wafv2-webacl-paymentnetwork-prices
	//
	Prices interface{} `field:"optional" json:"prices" yaml:"prices"`
	// The wallet address for receiving payments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-paymentnetwork.html#cfn-wafv2-webacl-paymentnetwork-walletaddress
	//
	WalletAddress *string `field:"optional" json:"walletAddress" yaml:"walletAddress"`
}

