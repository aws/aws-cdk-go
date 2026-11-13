package awsbedrockagentcore


// Options for an HTTP POST lifecycle hook.
//
// Example:
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   agentcore.NewRuntime(this, jsii.String("test-runtime"), &RuntimeProps{
//   	RuntimeName: jsii.String("test_runtime"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   	PlatformVersion: agentcore.PlatformVersion_V2(),
//   	BuildHooks: &BuildLifecycleHooks{
//   		Warmup: agentcore.LifecycleHook_HttpPost(&HttpPostLifecycleHookOptions{
//   			Path: jsii.String("/warmup"),
//   			Port: jsii.Number(8080),
//   		}),
//   	},
//   })
//
type HttpPostLifecycleHookOptions struct {
	// The path in the agent's container that the service calls, for example `/warmup`.
	//
	// Must be 1-32 characters, start with `/`, and must not be a route the service already calls
	// on the container, such as the invocation path or the default health check.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The port in the agent's container that the service calls.
	//
	// Must be between 8000 and 12000.
	// Default: - the runtime's server port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

