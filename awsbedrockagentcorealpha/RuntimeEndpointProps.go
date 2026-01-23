package awsbedrockagentcorealpha


// Properties for creating a Bedrock Agent Core Runtime Endpoint resource.
//
// Example:
//   // Reference an existing runtime by its ID
//   existingRuntimeId := "abc123-runtime-id" // The ID of an existing runtime
//
//   // Create a standalone endpoint
//   endpoint := agentcore.NewRuntimeEndpoint(this, jsii.String("MyEndpoint"), &RuntimeEndpointProps{
//   	EndpointName: jsii.String("production"),
//   	AgentRuntimeId: existingRuntimeId,
//   	AgentRuntimeVersion: jsii.String("1"),
//   	 // Specify which version to use
//   	Description: jsii.String("Production endpoint for existing runtime"),
//   })
//
// Experimental.
type RuntimeEndpointProps struct {
	// The ID of the agent runtime to associate with this endpoint This is the unique identifier of the runtime resource Pattern: ^[a-zA-Z][a-zA-Z0-9_]{0,99}-[a-zA-Z0-9]{10}$.
	// Experimental.
	AgentRuntimeId *string `field:"required" json:"agentRuntimeId" yaml:"agentRuntimeId"`
	// The version of the agent runtime to use for this endpoint If not specified, the endpoint will point to version "1" of the runtime.
	//
	// Pattern: ^([1-9][0-9]{0,4})$.
	// Default: "1".
	//
	// Experimental.
	AgentRuntimeVersion *string `field:"optional" json:"agentRuntimeVersion" yaml:"agentRuntimeVersion"`
	// Optional description for the agent runtime endpoint Length Minimum: 1 ,  Maximum: 256.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the agent runtime endpoint Valid characters are a-z, A-Z, 0-9, _ (underscore) Must start with a letter and can be up to 48 characters long Pattern: ^[a-zA-Z][a-zA-Z0-9_]{0,47}$.
	// Default: - auto generate.
	//
	// Experimental.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Tags for the agent runtime endpoint A list of key:value pairs of tags to apply to this RuntimeEndpoint resource Pattern: ^[a-zA-Z0-9\s._:/=+@-]*$.
	// Default: {} - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

