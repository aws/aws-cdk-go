package awsbedrockagentcore


// Optional gateway settings when binding an {@link IApiKeyCredentialProvider} to a target.
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
type FromApiKeyIdentityOptions struct {
	// Where to place the API key on outbound requests.
	// Default: header `Authorization` with `Bearer ` prefix.
	//
	CredentialLocation ApiKeyCredentialLocation `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
}

