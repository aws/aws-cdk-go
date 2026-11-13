package awsbedrockagentcore


// Hooks the AgentCore service calls while it builds the runtime.
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
type BuildLifecycleHooks struct {
	// Called to warm up the agent before the service snapshots it.
	// Default: - no warmup hook.
	//
	Warmup LifecycleHook `field:"optional" json:"warmup" yaml:"warmup"`
}

