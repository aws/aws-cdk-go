package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A hook the AgentCore service calls on the agent's container.
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
type LifecycleHook interface {
}

// The jsii proxy struct for LifecycleHook
type jsiiProxy_LifecycleHook struct {
	_ byte // padding
}

// Send an HTTP POST to the agent's container.
//
// Any 2xx response ends the hook; any other response fails it.
func LifecycleHook_HttpPost(options *HttpPostLifecycleHookOptions) LifecycleHook {
	_init_.Initialize()

	if err := validateLifecycleHook_HttpPostParameters(options); err != nil {
		panic(err)
	}
	var returns LifecycleHook

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.LifecycleHook",
		"httpPost",
		[]interface{}{options},
		&returns,
	)

	return returns
}

