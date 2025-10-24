package awsbedrockagentcorealpha


// Options for adding an endpoint to the runtime.
//
// Example:
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//
//   agentRuntimeArtifactNew := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v2.0.0"))
//
//   runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifactNew,
//   })
//
//   stagingEndpoint := runtime.AddEndpoint(jsii.String("staging"), &AddEndpointOptions{
//   	Version: jsii.String("2"),
//   	Description: jsii.String("Staging environment for testing new version"),
//   })
//
// Experimental.
type AddEndpointOptions struct {
	// Description for the endpoint, Must be between 1 and 1200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Override the runtime version for this endpoint.
	// Default: 1.
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

