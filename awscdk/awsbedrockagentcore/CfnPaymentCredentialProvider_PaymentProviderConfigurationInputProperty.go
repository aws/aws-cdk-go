package awsbedrockagentcore


// Provider configuration input containing secrets for creation/update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paymentProviderConfigurationInputProperty := &PaymentProviderConfigurationInputProperty{
//   	CoinbaseCdpConfiguration: &CoinbaseCdpConfigurationInputProperty{
//   		ApiKeyId: jsii.String("apiKeyId"),
//   		ApiKeySecret: jsii.String("apiKeySecret"),
//
//   		// the properties below are optional
//   		WalletSecret: jsii.String("walletSecret"),
//   	},
//   	StripePrivyConfiguration: &StripePrivyConfigurationInputProperty{
//   		AppId: jsii.String("appId"),
//   		AppSecret: jsii.String("appSecret"),
//   		AuthorizationId: jsii.String("authorizationId"),
//   		AuthorizationPrivateKey: jsii.String("authorizationPrivateKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput.html
//
type CfnPaymentCredentialProvider_PaymentProviderConfigurationInputProperty struct {
	// Coinbase CDP configuration with API credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput-coinbasecdpconfiguration
	//
	CoinbaseCdpConfiguration interface{} `field:"optional" json:"coinbaseCdpConfiguration" yaml:"coinbaseCdpConfiguration"`
	// Stripe Privy configuration with credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput-stripeprivyconfiguration
	//
	StripePrivyConfiguration interface{} `field:"optional" json:"stripePrivyConfiguration" yaml:"stripePrivyConfiguration"`
}

