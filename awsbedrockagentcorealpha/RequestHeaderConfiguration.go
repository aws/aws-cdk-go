package awsbedrockagentcorealpha


// Configuration for HTTP request headers that will be passed through to the runtime.
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
//   	RequestHeaderConfiguration: &RequestHeaderConfiguration{
//   		AllowlistedHeaders: []*string{
//   			jsii.String("X-Amzn-Bedrock-AgentCore-Runtime-Custom-H1"),
//   		},
//   	},
//   })
//
// Experimental.
type RequestHeaderConfiguration struct {
	// A list of HTTP request headers that are allowed to be passed through to the runtime.
	// Default: - No request headers allowed.
	//
	// Experimental.
	AllowlistedHeaders *[]*string `field:"optional" json:"allowlistedHeaders" yaml:"allowlistedHeaders"`
}

