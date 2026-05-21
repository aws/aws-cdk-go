package awsbedrockagentcore


// Coinbase CDP configuration with API credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coinbaseCdpConfigurationInputProperty := &CoinbaseCdpConfigurationInputProperty{
//   	ApiKeyId: jsii.String("apiKeyId"),
//   	ApiKeySecret: jsii.String("apiKeySecret"),
//
//   	// the properties below are optional
//   	WalletSecret: jsii.String("walletSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html
//
type CfnPaymentCredentialProvider_CoinbaseCdpConfigurationInputProperty struct {
	// The Coinbase CDP API key ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-apikeyid
	//
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
	// The Coinbase CDP API key secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-apikeysecret
	//
	ApiKeySecret *string `field:"required" json:"apiKeySecret" yaml:"apiKeySecret"`
	// The Coinbase CDP wallet secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-walletsecret
	//
	WalletSecret *string `field:"optional" json:"walletSecret" yaml:"walletSecret"`
}

