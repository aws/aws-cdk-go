package awsbedrockagentcorealpha


// Options for adding an OpenAPI target to a gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // These ARNs are returned when creating the API key credential provider via Console or API
//   apiKeyProviderArn := "arn:aws:bedrock-agentcore:us-east-1:123456789012:token-vault/abc123/apikeycredentialprovider/my-apikey"
//   apiKeySecretArn := "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-apikey-secret-abc123"
//
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3mySchema := agentcore.ApiSchema_FromS3File(bucket, jsii.String("schemas/myschema.yaml"))
//
//   // Add an OpenAPI target directly to the gateway
//   target := gateway.AddOpenApiTarget(jsii.String("MyTarget"), &AddOpenApiTargetOptions{
//   	GatewayTargetName: jsii.String("my-api-target"),
//   	Description: jsii.String("Target for external API integration"),
//   	ApiSchema: s3mySchema,
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromApiKeyIdentityArn(&ApiKeyCredentialProviderProps{
//   			ProviderArn: apiKeyProviderArn,
//   			SecretArn: apiKeySecretArn,
//   			CredentialLocation: agentcore.ApiKeyCredentialLocation_Header(&ApiKeyAdditionalConfiguration{
//   				CredentialParameterName: jsii.String("X-API-Key"),
//   			}),
//   		}),
//   	},
//   })
//
//   // This make sure your s3 bucket is available before target
//   target.Node.AddDependency(bucket)
//
// Experimental.
type AddOpenApiTargetOptions struct {
	// The OpenAPI schema defining the API.
	// Experimental.
	ApiSchema ApiSchema `field:"required" json:"apiSchema" yaml:"apiSchema"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Experimental.
	GatewayTargetName *string `field:"required" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Credential providers for authentication.
	// Default: - [GatewayCredentialProvider.iamRole()]
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to validate the OpenAPI schema or not Note: Validation is only performed for inline and asset-based schema and  during CDK synthesis.
	//
	// S3 schemas cannot be validated at synthesis time.
	// Default: true.
	//
	// Experimental.
	ValidateOpenApiSchema *bool `field:"optional" json:"validateOpenApiSchema" yaml:"validateOpenApiSchema"`
}

