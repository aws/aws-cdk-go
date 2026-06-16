package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApiKeyCredentialProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyCredentialProviderProps := &CfnApiKeyCredentialProviderProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ApiKey: jsii.String("apiKey"),
//   	ApiKeySecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	ApiKeySecretSource: jsii.String("apiKeySecretSource"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html
//
type CfnApiKeyCredentialProviderProps struct {
	// The name of the API key credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API key to use for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-apikeysecretconfig
	//
	ApiKeySecretConfig interface{} `field:"optional" json:"apiKeySecretConfig" yaml:"apiKeySecretConfig"`
	// The source of the API key secret.
	//
	// Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-apikeysecretsource
	//
	ApiKeySecretSource *string `field:"optional" json:"apiKeySecretSource" yaml:"apiKeySecretSource"`
	// Tags to assign to the API key credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

