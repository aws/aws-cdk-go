package awsbedrockagentcorealpha


// Attributes for importing an existing API key credential provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   apiKeyCredentialProviderAttributes := &ApiKeyCredentialProviderAttributes{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//
//   	// the properties below are optional
//   	ApiKeySecretArn: jsii.String("apiKeySecretArn"),
//   	CreatedTime: jsii.String("createdTime"),
//   	LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type ApiKeyCredentialProviderAttributes struct {
	// ARN of the credential provider.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// ARN of the Secrets Manager secret for the API key, if known.
	// Default: - not set; required for {@link ApiKeyCredentialProvider.bindForGatewayApiKeyTarget } on imported providers
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ApiKeySecretArn *string `field:"optional" json:"apiKeySecretArn" yaml:"apiKeySecretArn"`
	// Resource creation time.
	// Default: - not set.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Resource last-updated time.
	// Default: - not set.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
}

