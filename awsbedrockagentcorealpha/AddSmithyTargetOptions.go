package awsbedrockagentcorealpha


// Options for adding a Smithy target to a gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   smithySchema := agentcore.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("models"), jsii.String("smithy-model.json")))
//   smithySchema.Bind(this)
//
//   smithyTarget := gateway.AddSmithyTarget(jsii.String("MySmithyTarget"), &AddSmithyTargetOptions{
//   	GatewayTargetName: jsii.String("my-smithy-target"),
//   	Description: jsii.String("Smithy model target"),
//   	SmithyModel: smithySchema,
//   })
//
// Experimental.
type AddSmithyTargetOptions struct {
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Experimental.
	GatewayTargetName *string `field:"required" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// The Smithy model defining the API.
	// Experimental.
	SmithyModel ApiSchema `field:"required" json:"smithyModel" yaml:"smithyModel"`
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
}

