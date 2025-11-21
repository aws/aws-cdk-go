package awsbedrockagentcorealpha


// Properties for creating a Gateway Target resource.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   myLambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	    exports.handler = async (event) => ({ statusCode: 200 });
//   	  `)),
//   })
//
//   myToolSchema := agentcore.ToolSchema_FromInline([]ToolDefinition{
//   	&ToolDefinition{
//   		Name: jsii.String("my_tool"),
//   		Description: jsii.String("My custom tool"),
//   		InputSchema: &SchemaDefinition{
//   			Type: agentcore.SchemaDefinitionType_OBJECT,
//   			Properties: map[string]interface{}{
//   			},
//   		},
//   	},
//   })
//
//   // Create a custom Lambda configuration
//   customConfig := agentcore.LambdaTargetConfiguration_Create(myLambdaFunction, myToolSchema)
//
//   // Use the GatewayTarget constructor directly
//   target := agentcore.NewGatewayTarget(this, jsii.String("AdvancedTarget"), &GatewayTargetProps{
//   	Gateway: gateway,
//   	GatewayTargetName: jsii.String("advanced-target"),
//   	TargetConfiguration: customConfig,
//   	 // Manually created configuration
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromIamRole(),
//   	},
//   })
//
// Experimental.
type GatewayTargetProps struct {
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
	// The target configuration (Lambda, OpenAPI, or Smithy) Use one of the configuration helper classes: - LambdaTargetConfiguration.create() - OpenApiTargetConfiguration.create() - SmithyTargetConfiguration.create().
	// Experimental.
	TargetConfiguration ITargetConfiguration `field:"required" json:"targetConfiguration" yaml:"targetConfiguration"`
	// Credential providers for authentication.
	// Default: - [GatewayCredentialProvider.fromIamRole()]
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
}

