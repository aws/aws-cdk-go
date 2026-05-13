package awsbedrockagentcorealpha


// API key credential provider ARNs for gateway outbound auth (Token Vault identity).
//
// Pass this to {@link GatewayCredentialProvider.fromApiKeyIdentityArn } or to {@link ApiKeyCredentialProviderConfiguration}.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // ARNs from the console/API, or from ApiKeyCredentialProvider + bindForGatewayApiKeyTarget
//   apiKeyProviderArn := "arn:aws:bedrock-agentcore:us-east-1:123456789012:token-vault/abc123/apikeycredentialprovider/my-apikey"
//   apiKeySecretArn := "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-apikey-secret-abc123"
//
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3mySchema := agentcore.ApiSchema_FromS3File(bucket, jsii.String("schemas/myschema.yaml"))
//
//   // Add an OpenAPI target using ARNs directly
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
//   // This makes sure your s3 bucket is available before target
//   target.Node.AddDependency(bucket)
//
// Experimental.
type ApiKeyCredentialProviderProps struct {
	// The API key credential provider ARN.
	//
	// This is returned when creating the API key credential provider via Console or API.
	// Format: arn:aws:bedrock-agentcore:region:account:token-vault/id/apikeycredentialprovider/name.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// The ARN of the Secrets Manager secret containing the API key.
	//
	// This is returned when creating the API key credential provider via Console or API.
	// Format: arn:aws:secretsmanager:region:account:secret:name.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The location of the API key credential.
	//
	// This field specifies where in the request the API key should be placed.
	// Default: - HEADER.
	//
	// Experimental.
	CredentialLocation ApiKeyCredentialLocation `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
}

