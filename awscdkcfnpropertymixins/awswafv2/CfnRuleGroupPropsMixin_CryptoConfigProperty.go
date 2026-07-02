package awswafv2


// Configures cryptocurrency payment settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cryptoConfigProperty := &CryptoConfigProperty{
//   	PaymentNetworks: []interface{}{
//   		&PaymentNetworkProperty{
//   			Chain: jsii.String("chain"),
//   			Prices: []interface{}{
//   				&PriceProperty{
//   					Amount: jsii.String("amount"),
//   					Currency: jsii.String("currency"),
//   				},
//   			},
//   			WalletAddress: jsii.String("walletAddress"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-cryptoconfig.html
//
type CfnRuleGroupPropsMixin_CryptoConfigProperty struct {
	// List of payment network configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-cryptoconfig.html#cfn-wafv2-rulegroup-cryptoconfig-paymentnetworks
	//
	PaymentNetworks interface{} `field:"optional" json:"paymentNetworks" yaml:"paymentNetworks"`
}

