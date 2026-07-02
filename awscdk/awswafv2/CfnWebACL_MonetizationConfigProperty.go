package awswafv2


// Configures monetization for the web ACL or rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monetizationConfigProperty := &MonetizationConfigProperty{
//   	CryptoConfig: &CryptoConfigProperty{
//   		PaymentNetworks: []interface{}{
//   			&PaymentNetworkProperty{
//   				Chain: jsii.String("chain"),
//   				Prices: []interface{}{
//   					&PriceProperty{
//   						Amount: jsii.String("amount"),
//   						Currency: jsii.String("currency"),
//   					},
//   				},
//   				WalletAddress: jsii.String("walletAddress"),
//   			},
//   		},
//   	},
//   	CurrencyMode: jsii.String("currencyMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-monetizationconfig.html
//
type CfnWebACL_MonetizationConfigProperty struct {
	// Configures cryptocurrency payment settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-monetizationconfig.html#cfn-wafv2-webacl-monetizationconfig-cryptoconfig
	//
	CryptoConfig interface{} `field:"optional" json:"cryptoConfig" yaml:"cryptoConfig"`
	// The currency mode for monetization.
	//
	// Use REAL for production payments and TEST for testing with testnet currencies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-monetizationconfig.html#cfn-wafv2-webacl-monetizationconfig-currencymode
	//
	CurrencyMode *string `field:"optional" json:"currencyMode" yaml:"currencyMode"`
}

