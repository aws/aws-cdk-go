package awsbedrockagentcore


// Options for adding an OpenAPI target to a gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // Create an API key credential provider in Token Vault
//   apiKeyProvider := agentcore.NewApiKeyCredentialProvider(this, jsii.String("MyApiKeyProvider"), &ApiKeyCredentialProviderProps{
//   	ApiKeyCredentialProviderName: jsii.String("my-apikey"),
//   })
//
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3mySchema := agentcore.ApiSchema_FromS3File(bucket, jsii.String("schemas/myschema.yaml"))
//
//   // Add an OpenAPI target using the L2 construct directly
//   target := gateway.AddOpenApiTarget(jsii.String("MyTarget"), &AddOpenApiTargetOptions{
//   	GatewayTargetName: jsii.String("my-api-target"),
//   	Description: jsii.String("Target for external API integration"),
//   	ApiSchema: s3mySchema,
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromApiKeyIdentity(apiKeyProvider, &FromApiKeyIdentityOptions{
//   			CredentialLocation: agentcore.ApiKeyCredentialLocation_Header(&ApiKeyAdditionalConfiguration{
//   				CredentialParameterName: jsii.String("X-API-Key"),
//   			}),
//   		}),
//   	},
//   })
//
//   // This makes sure your s3 bucket is available before target
//   target.Node.AddDependency(bucket)
//
type AddOpenApiTargetOptions struct {
	// The OpenAPI schema defining the API.
	ApiSchema ApiSchema `field:"required" json:"apiSchema" yaml:"apiSchema"`
	// Credential providers for outbound authentication (OpenAPI targets use API Key or OAuth, not IAM).
	// Default: - none (no credential configuration on the target; supply providers for secured backends).
	//
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Whether to validate the OpenAPI schema or not Note: Validation is only performed for inline and asset-based schema and  during CDK synthesis.
	//
	// S3 schemas cannot be validated at synthesis time.
	// Default: true.
	//
	ValidateOpenApiSchema *bool `field:"optional" json:"validateOpenApiSchema" yaml:"validateOpenApiSchema"`
}

