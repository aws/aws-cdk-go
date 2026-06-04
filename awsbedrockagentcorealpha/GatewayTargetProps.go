package awsbedrockagentcorealpha


// Properties for creating a Gateway Target resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var credentialProviderConfig ICredentialProviderConfig
//   var gateway Gateway
//   var targetConfiguration ITargetConfiguration
//
//   gatewayTargetProps := &GatewayTargetProps{
//   	Gateway: gateway,
//   	TargetConfiguration: targetConfiguration,
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
type GatewayTargetProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// The gateway this target belongs to.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The target configuration (Lambda, OpenAPI, or Smithy) Use one of the configuration helper classes: - LambdaTargetConfiguration.create() - OpenApiTargetConfiguration.create() - SmithyTargetConfiguration.create().
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	TargetConfiguration ITargetConfiguration `field:"required" json:"targetConfiguration" yaml:"targetConfiguration"`
	// Credential providers for authentication.
	// Default: - [GatewayCredentialProvider.fromIamRole()]
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
}

