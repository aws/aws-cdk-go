package awsbedrockagentcore


// Provider configuration output containing secret ARNs (no raw secrets).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paymentProviderConfigurationOutputProperty := &PaymentProviderConfigurationOutputProperty{
//   	CoinbaseCdpConfiguration: &CoinbaseCdpConfigurationOutputProperty{
//   		ApiKeyId: jsii.String("apiKeyId"),
//   		ApiKeySecretArn: &SecretInfoProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//
//   		// the properties below are optional
//   		WalletSecretArn: &SecretInfoProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	StripePrivyConfiguration: &StripePrivyConfigurationOutputProperty{
//   		AppId: jsii.String("appId"),
//   		AppSecretArn: &SecretInfoProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		AuthorizationId: jsii.String("authorizationId"),
//   		AuthorizationPrivateKeyArn: &SecretInfoProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationoutput.html
//
type CfnPaymentCredentialProvider_PaymentProviderConfigurationOutputProperty struct {
	// Coinbase CDP configuration output with secret ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationoutput-coinbasecdpconfiguration
	//
	CoinbaseCdpConfiguration interface{} `field:"optional" json:"coinbaseCdpConfiguration" yaml:"coinbaseCdpConfiguration"`
	// Stripe Privy configuration output with secret ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-paymentproviderconfigurationoutput-stripeprivyconfiguration
	//
	StripePrivyConfiguration interface{} `field:"optional" json:"stripePrivyConfiguration" yaml:"stripePrivyConfiguration"`
}

