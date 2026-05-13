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
// Experimental.
type OAuth2CredentialProviderAttributes struct {
	// ARN of the credential provider.
	// Experimental.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// Vendor string for this provider.
	// Experimental.
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// Callback URL from the deployed provider, if known.
	// Default: - not set.
	//
	// Experimental.
	CallbackUrl *string `field:"optional" json:"callbackUrl" yaml:"callbackUrl"`
	// ARN of the Secrets Manager secret for OAuth2 client credentials, if known.
	// Default: - not set; required for {@link OAuth2CredentialProvider.bindForGatewayOAuthTarget } on imported providers
	//
	// Experimental.
	ClientSecretArn *string `field:"optional" json:"clientSecretArn" yaml:"clientSecretArn"`
	// Resource creation time.
	// Default: - not set.
	//
	// Experimental.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Resource last-updated time.
	// Default: - not set.
	//
	// Experimental.
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
}

