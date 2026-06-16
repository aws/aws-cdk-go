package awsbedrockagentcore


// Provider configuration input containing secrets for creation/update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   paymentProviderConfigurationInputProperty := &PaymentProviderConfigurationInputProperty{
//   	CoinbaseCdpConfiguration: &CoinbaseCdpConfigurationInputProperty{
//   		ApiKeyId: jsii.String("apiKeyId"),
//   		ApiKeySecret: jsii.String("apiKeySecret"),
//   		ApiKeySecretConfig: &SecretReferenceProperty{
//   			JsonKey: jsii.String("jsonKey"),
//   			SecretId: jsii.String("secretId"),
//   		},
//   		ApiKeySecretSource: jsii.String("apiKeySecretSource"),
//   		WalletSecret: jsii.String("walletSecret"),
//   		WalletSecretConfig: &SecretReferenceProperty{
//   			JsonKey: jsii.String("jsonKey"),
//   			SecretId: jsii.String("secretId"),
//   		},
//   		WalletSecretSource: jsii.String("walletSecretSource"),
//   	},
//   	StripePrivyConfiguration: &StripePrivyConfigurationInputProperty{
//   		AppId: jsii.String("appId"),
//   		AppSecret: jsii.String("appSecret"),
//   		AppSecretConfig: &SecretReferenceProperty{
//   			JsonKey: jsii.String("jsonKey"),
//   			SecretId: jsii.String("secretId"),
//   		},
//   		AppSecretSource: jsii.String("appSecretSource"),
//   		AuthorizationId: jsii.String("authorizationId"),
//   		AuthorizationPrivateKey: jsii.String("authorizationPrivateKey"),
//   		AuthorizationPrivateKeyConfig: &SecretReferenceProperty{
//   			JsonKey: jsii.String("jsonKey"),
//   			SecretId: jsii.String("secretId"),
//   		},
//   		AuthorizationPrivateKeySource: jsii.String("authorizationPrivateKeySource"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput.html
//
type CfnPaymentCredentialProviderPropsMixin_PaymentProviderConfigurationInputProperty struct {
	// Coinbase CDP configuration with API credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput-coinbasecdpconfiguration
	//
	CoinbaseCdpConfiguration interface{} `field:"optional" json:"coinbaseCdpConfiguration" yaml:"coinbaseCdpConfiguration"`
	// Stripe Privy configuration with credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationinput-stripeprivyconfiguration
	//
	StripePrivyConfiguration interface{} `field:"optional" json:"stripePrivyConfiguration" yaml:"stripePrivyConfiguration"`
}

