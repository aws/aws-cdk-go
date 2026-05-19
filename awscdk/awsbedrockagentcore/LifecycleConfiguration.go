package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// LifecycleConfiguration lets you manage the lifecycle of runtime sessions and resources in AgentCore Runtime.
//
// This configuration helps optimize resource utilization by automatically cleaning up idle sessions and preventing
// long-running instances from consuming resources indefinitely.
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
//   	LifecycleConfiguration: &LifecycleConfiguration{
//   		IdleRuntimeSessionTimeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   		MaxLifetime: awscdk.Duration_Hours(jsii.Number(4)),
//   	},
//   })
//
type LifecycleConfiguration struct {
	// Timeout in seconds for idle runtime sessions.
	//
	// When a session remains idle for this duration,
	// it will be automatically terminated.
	// Default: undefined - service default setting is 900 seconds (15 minutes).
	//
	IdleRuntimeSessionTimeout awscdk.Duration `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Maximum lifetime for the instance in seconds.
	//
	// Once reached, instances will be automatically
	// terminated and replaced.
	// Default: undefined - service default setting is 28800 seconds (8 hours).
	//
	MaxLifetime awscdk.Duration `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

