package awsbedrockagentcorealpha


// Attributes for importing an existing OAuth2 credential provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   oAuth2CredentialProviderAttributes := &OAuth2CredentialProviderAttributes{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//
//   	// the properties below are optional
//   	CallbackUrl: jsii.String("callbackUrl"),
//   	ClientSecretArn: jsii.String("clientSecretArn"),
//   	CreatedTime: jsii.String("createdTime"),
//   	LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type OAuth2CredentialProviderAttributes struct {
	// ARN of the credential provider.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// Vendor string for this provider.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// Callback URL from the deployed provider, if known.
	// Default: - not set.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CallbackUrl *string `field:"optional" json:"callbackUrl" yaml:"callbackUrl"`
	// ARN of the Secrets Manager secret for OAuth2 client credentials, if known.
	// Default: - not set; required for {@link OAuth2CredentialProvider.bindForGatewayOAuthTarget } on imported providers
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ClientSecretArn *string `field:"optional" json:"clientSecretArn" yaml:"clientSecretArn"`
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

