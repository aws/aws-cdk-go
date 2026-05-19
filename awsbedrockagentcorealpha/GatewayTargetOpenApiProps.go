package awsbedrockagentcorealpha


// Properties for creating an OpenAPI-based Gateway Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiSchema ApiSchema
//   var credentialProviderConfig ICredentialProviderConfig
//   var gateway Gateway
//
//   gatewayTargetOpenApiProps := &GatewayTargetOpenApiProps{
//   	ApiSchema: apiSchema,
//   	Gateway: gateway,
//
//   	// the properties below are optional
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		credentialProviderConfig,
//   	},
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   	ValidateOpenApiSchema: jsii.Boolean(false),
//   }
//
// Experimental.
type GatewayTargetOpenApiProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// The OpenAPI schema defining the API.
	// Experimental.
	ApiSchema ApiSchema `field:"required" json:"apiSchema" yaml:"apiSchema"`
	// The gateway this target belongs to.
	// Experimental.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// Credential providers for authentication OpenAPI targets support API key and OAuth authentication (not IAM).
	// Default: - If not provided, defaults to IAM role which will fail at runtime.
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Whether to validate the OpenAPI schema (only applies to inline schemas) Note: Validation is only performed for inline schemas during CDK synthesis.
	//
	// S3 and asset-based schemas cannot be validated at synthesis time.
	// Default: true.
	//
	// Experimental.
	ValidateOpenApiSchema *bool `field:"optional" json:"validateOpenApiSchema" yaml:"validateOpenApiSchema"`
}

