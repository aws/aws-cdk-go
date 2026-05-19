package awsbedrockagentcore


// Attributes for importing an existing API key credential provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type ApiKeyCredentialProviderAttributes struct {
	// ARN of the credential provider.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// ARN of the Secrets Manager secret for the API key, if known.
	// Default: - not set; required for {@link ApiKeyCredentialProvider.bindForGatewayApiKeyTarget } on imported providers
	//
	ApiKeySecretArn *string `field:"optional" json:"apiKeySecretArn" yaml:"apiKeySecretArn"`
	// Resource creation time.
	// Default: - not set.
	//
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Resource last-updated time.
	// Default: - not set.
	//
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
}

