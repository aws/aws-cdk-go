package awsbedrockagentcore


// Coinbase CDP configuration output with secret ARNs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coinbaseCdpConfigurationOutputProperty := &CoinbaseCdpConfigurationOutputProperty{
//   	ApiKeyId: jsii.String("apiKeyId"),
//   	ApiKeySecretArn: &SecretInfoProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//
//   	// the properties below are optional
//   	ApiKeySecretJsonKey: jsii.String("apiKeySecretJsonKey"),
//   	ApiKeySecretSource: jsii.String("apiKeySecretSource"),
//   	WalletSecretArn: &SecretInfoProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	WalletSecretJsonKey: jsii.String("walletSecretJsonKey"),
//   	WalletSecretSource: jsii.String("walletSecretSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html
//
type CfnPaymentCredentialProvider_CoinbaseCdpConfigurationOutputProperty struct {
	// The Coinbase CDP API key ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-apikeyid
	//
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-apikeysecretarn
	//
	ApiKeySecretArn interface{} `field:"required" json:"apiKeySecretArn" yaml:"apiKeySecretArn"`
	// The JSON key within the secret that contains the API key secret value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-apikeysecretjsonkey
	//
	ApiKeySecretJsonKey *string `field:"optional" json:"apiKeySecretJsonKey" yaml:"apiKeySecretJsonKey"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-apikeysecretsource
	//
	ApiKeySecretSource *string `field:"optional" json:"apiKeySecretSource" yaml:"apiKeySecretSource"`
	// Contains information about a secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-walletsecretarn
	//
	WalletSecretArn interface{} `field:"optional" json:"walletSecretArn" yaml:"walletSecretArn"`
	// The JSON key within the secret that contains the wallet secret value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-walletsecretjsonkey
	//
	WalletSecretJsonKey *string `field:"optional" json:"walletSecretJsonKey" yaml:"walletSecretJsonKey"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationoutput-walletsecretsource
	//
	WalletSecretSource *string `field:"optional" json:"walletSecretSource" yaml:"walletSecretSource"`
}

