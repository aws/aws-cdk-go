package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPaymentCredentialProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPaymentCredentialProviderProps := &CfnPaymentCredentialProviderProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ProviderConfigurationInput: &PaymentProviderConfigurationInputProperty{
//   		CoinbaseCdpConfiguration: &CoinbaseCdpConfigurationInputProperty{
//   			ApiKeyId: jsii.String("apiKeyId"),
//
//   			// the properties below are optional
//   			ApiKeySecret: jsii.String("apiKeySecret"),
//   			ApiKeySecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ApiKeySecretSource: jsii.String("apiKeySecretSource"),
//   			WalletSecret: jsii.String("walletSecret"),
//   			WalletSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			WalletSecretSource: jsii.String("walletSecretSource"),
//   		},
//   		StripePrivyConfiguration: &StripePrivyConfigurationInputProperty{
//   			AppId: jsii.String("appId"),
//   			AuthorizationId: jsii.String("authorizationId"),
//
//   			// the properties below are optional
//   			AppSecret: jsii.String("appSecret"),
//   			AppSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			AppSecretSource: jsii.String("appSecretSource"),
//   			AuthorizationPrivateKey: jsii.String("authorizationPrivateKey"),
//   			AuthorizationPrivateKeyConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			AuthorizationPrivateKeySource: jsii.String("authorizationPrivateKeySource"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentcredentialprovider.html
//
type CfnPaymentCredentialProviderProps struct {
	// Supported vendor types for payment providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentcredentialprovider.html#cfn-bedrockagentcore-paymentcredentialprovider-credentialprovidervendor
	//
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// Unique name for the payment credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentcredentialprovider.html#cfn-bedrockagentcore-paymentcredentialprovider-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Provider configuration input containing secrets for creation/update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentcredentialprovider.html#cfn-bedrockagentcore-paymentcredentialprovider-providerconfigurationinput
	//
	ProviderConfigurationInput interface{} `field:"optional" json:"providerConfigurationInput" yaml:"providerConfigurationInput"`
	// Tags for the payment credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-paymentcredentialprovider.html#cfn-bedrockagentcore-paymentcredentialprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

