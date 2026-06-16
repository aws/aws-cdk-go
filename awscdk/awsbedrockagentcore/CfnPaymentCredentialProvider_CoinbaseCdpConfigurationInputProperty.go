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
//
//   	// the properties below are optional
//   	ApiKeySecret: jsii.String("apiKeySecret"),
//   	ApiKeySecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	ApiKeySecretSource: jsii.String("apiKeySecretSource"),
//   	WalletSecret: jsii.String("walletSecret"),
//   	WalletSecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	WalletSecretSource: jsii.String("walletSecretSource"),
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
	ApiKeySecret *string `field:"optional" json:"apiKeySecret" yaml:"apiKeySecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-apikeysecretconfig
	//
	ApiKeySecretConfig interface{} `field:"optional" json:"apiKeySecretConfig" yaml:"apiKeySecretConfig"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-apikeysecretsource
	//
	ApiKeySecretSource *string `field:"optional" json:"apiKeySecretSource" yaml:"apiKeySecretSource"`
	// The Coinbase CDP wallet secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-walletsecret
	//
	WalletSecret *string `field:"optional" json:"walletSecret" yaml:"walletSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-walletsecretconfig
	//
	WalletSecretConfig interface{} `field:"optional" json:"walletSecretConfig" yaml:"walletSecretConfig"`
	// The source of the secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput.html#cfn-bedrockagentcore-paymentcredentialprovider-coinbasecdpconfigurationinput-walletsecretsource
	//
	WalletSecretSource *string `field:"optional" json:"walletSecretSource" yaml:"walletSecretSource"`
}

