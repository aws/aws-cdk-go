package awsbedrockagentcorealpha


// API Key additional configuration.
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
type ApiKeyAdditionalConfiguration struct {
	// The name of the credential parameter for the API key.
	//
	// This parameter name is used when sending the API key to the target endpoint.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 64.
	// Default: - 'Authorization' for HEADER, 'api_key' for QUERY_PARAMETER.
	//
	// Experimental.
	CredentialParameterName *string `field:"optional" json:"credentialParameterName" yaml:"credentialParameterName"`
	// The prefix for the API key credential.
	//
	// This prefix is added to the API key when sending it to the target endpoint.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 64.
	// Default: - 'Bearer ' for HEADER, no prefix for QUERY_PARAMETER.
	//
	// Experimental.
	CredentialPrefix *string `field:"optional" json:"credentialPrefix" yaml:"credentialPrefix"`
}

