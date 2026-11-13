package awswafv2


// Configures monetization for the web ACL or rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   monetizationConfigProperty := &MonetizationConfigProperty{
//   	ConnectedAccountConfigs: []interface{}{
//   		&ConnectedAccountConfigProperty{
//   			ConnectedAccountArn: jsii.String("connectedAccountArn"),
//   			PaymentNetworks: []interface{}{
//   				&ConnectedAccountChainConfigProperty{
//   					Chain: jsii.String("chain"),
//   					Prices: []interface{}{
//   						&ConnectedAccountPriceProperty{
//   							Amount: jsii.String("amount"),
//   							Currency: jsii.String("currency"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
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
type CfnWebACLPropsMixin_MonetizationConfigProperty struct {
	// List of connected account configurations.
	//
	// A connected account delegates deposit-address provisioning to the payments provider instead of specifying a wallet directly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-monetizationconfig.html#cfn-wafv2-webacl-monetizationconfig-connectedaccountconfigs
	//
	ConnectedAccountConfigs interface{} `field:"optional" json:"connectedAccountConfigs" yaml:"connectedAccountConfigs"`
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

