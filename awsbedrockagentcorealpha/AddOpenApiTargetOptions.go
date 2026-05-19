package awsbedrockagentcorealpha


// Options for adding an OpenAPI target to a gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiSchema ApiSchema
//   var credentialProviderConfig ICredentialProviderConfig
//
//   addOpenApiTargetOptions := &AddOpenApiTargetOptions{
//   	ApiSchema: apiSchema,
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
type AddOpenApiTargetOptions struct {
	// The OpenAPI schema defining the API.
	// Experimental.
	ApiSchema ApiSchema `field:"required" json:"apiSchema" yaml:"apiSchema"`
	// Credential providers for outbound authentication (OpenAPI targets use API Key or OAuth, not IAM).
	// Default: - none (no credential configuration on the target; supply providers for secured backends).
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Whether to validate the OpenAPI schema or not Note: Validation is only performed for inline and asset-based schema and  during CDK synthesis.
	//
	// S3 schemas cannot be validated at synthesis time.
	// Default: true.
	//
	// Experimental.
	ValidateOpenApiSchema *bool `field:"optional" json:"validateOpenApiSchema" yaml:"validateOpenApiSchema"`
}

