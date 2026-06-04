package awsbedrockagentcorealpha


// Options for adding a Smithy target to a gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiSchema ApiSchema
//   var credentialProviderConfig ICredentialProviderConfig
//
//   addSmithyTargetOptions := &AddSmithyTargetOptions{
//   	SmithyModel: apiSchema,
//
//   	// the properties below are optional
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		credentialProviderConfig,
//   	},
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type AddSmithyTargetOptions struct {
	// The Smithy model defining the API.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	SmithyModel ApiSchema `field:"required" json:"smithyModel" yaml:"smithyModel"`
	// Credential providers for authentication.
	// Default: - [GatewayCredentialProvider.iamRole()]
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
}

