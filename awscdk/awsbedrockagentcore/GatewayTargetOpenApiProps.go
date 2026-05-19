package awsbedrockagentcore


// Properties for creating an OpenAPI-based Gateway Target.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // Outbound auth: ApiKeyCredentialProvider + bindForGatewayApiKeyTarget, or ARNs from console/API
//   apiKeyIdentityArn := "arn:aws:bedrock-agentcore:us-east-1:123456789012:token-vault/abc123/apikeycredentialprovider/my-apikey"
//   apiKeySecretArn := "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-apikey-secret-abc123"
//
//   opneapiSchema := agentcore.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("mySchema.yml")))
//   opneapiSchema.Bind(this)
//
//   // Create a gateway target with OpenAPI Schema
//   target := agentcore.GatewayTarget_ForOpenApi(this, jsii.String("MyTarget"), &GatewayTargetOpenApiProps{
//   	GatewayTargetName: jsii.String("my-api-target"),
//   	Description: jsii.String("Target for external API integration"),
//   	Gateway: gateway,
//   	 // Note: you need to pass the gateway reference
//   	ApiSchema: opneapiSchema,
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromApiKeyIdentityArn(&ApiKeyCredentialProviderOptions{
//   			ProviderArn: apiKeyIdentityArn,
//   			SecretArn: apiKeySecretArn,
//   		}),
//   	},
//   })
//
type GatewayTargetOpenApiProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// The OpenAPI schema defining the API.
	ApiSchema ApiSchema `field:"required" json:"apiSchema" yaml:"apiSchema"`
	// The gateway this target belongs to.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// Credential providers for authentication OpenAPI targets support API key and OAuth authentication (not IAM).
	// Default: - If not provided, defaults to IAM role which will fail at runtime.
	//
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Whether to validate the OpenAPI schema (only applies to inline schemas) Note: Validation is only performed for inline schemas during CDK synthesis.
	//
	// S3 and asset-based schemas cannot be validated at synthesis time.
	// Default: true.
	//
	ValidateOpenApiSchema *bool `field:"optional" json:"validateOpenApiSchema" yaml:"validateOpenApiSchema"`
}

