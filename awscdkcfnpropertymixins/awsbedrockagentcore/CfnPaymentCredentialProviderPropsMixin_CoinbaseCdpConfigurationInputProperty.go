package awsbedrockagentcore


// Coinbase CDP configuration with API credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   coinbaseCdpConfigurationInputProperty := &CoinbaseCdpConfigurationInputProperty{
//   	ApiKeyId: jsii.String("apiKeyId"),
//   	ApiKeySecret: jsii.String("apiKeySecret"),
//   	WalletSecret: jsii.String("walletSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html
//
type CfnPaymentCredentialProviderPropsMixin_CoinbaseCdpConfigurationInputProperty struct {
	// The Coinbase CDP API key ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-apikeyid
	//
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// The Coinbase CDP API key secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-apikeysecret
	//
	ApiKeySecret *string `field:"optional" json:"apiKeySecret" yaml:"apiKeySecret"`
	// The Coinbase CDP wallet secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-walletsecret
	//
	WalletSecret *string `field:"optional" json:"walletSecret" yaml:"walletSecret"`
}

