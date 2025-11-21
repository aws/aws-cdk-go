package awsbedrockagentcorealpha


// Properties for creating a Smithy-based Gateway Target.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   smithySchema := agentcore.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("models"), jsii.String("smithy-model.json")))
//   smithySchema.Bind(this)
//
//   // Create a gateway target with Smithy Model and OAuth
//   target := agentcore.GatewayTarget_ForSmithy(this, jsii.String("MySmithyTarget"), &GatewayTargetSmithyProps{
//   	GatewayTargetName: jsii.String("my-smithy-target"),
//   	Description: jsii.String("Target for Smithy model integration"),
//   	Gateway: gateway,
//   	SmithyModel: smithySchema,
//   })
//
// Experimental.
type GatewayTargetSmithyProps struct {
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Experimental.
	GatewayTargetName *string `field:"required" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The gateway this target belongs to.
	// Experimental.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The Smithy model defining the API.
	// Experimental.
	SmithyModel ApiSchema `field:"required" json:"smithyModel" yaml:"smithyModel"`
	// Credential providers for authentication Smithy targets only support IAM role authentication.
	// Default: - [GatewayCredentialProvider.fromIamRole()]
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
}

