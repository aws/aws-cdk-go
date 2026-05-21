package awsbedrockagentcore


// Coinbase CDP configuration output with secret ARNs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   coinbaseCdpConfigurationOutputProperty := &CoinbaseCdpConfigurationOutputProperty{
//   	ApiKeyId: jsii.String("apiKeyId"),
//   	ApiKeySecretArn: &SecretInfoProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	WalletSecretArn: &SecretInfoProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html
//
type CfnPaymentCredentialProviderPropsMixin_CoinbaseCdpConfigurationOutputProperty struct {
	// The Coinbase CDP API key ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-apikeyid
	//
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-apikeysecretarn
	//
	ApiKeySecretArn interface{} `field:"optional" json:"apiKeySecretArn" yaml:"apiKeySecretArn"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-walletsecretarn
	//
	WalletSecretArn interface{} `field:"optional" json:"walletSecretArn" yaml:"walletSecretArn"`
}

