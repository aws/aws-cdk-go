package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// LifecycleConfiguration lets you manage the lifecycle of runtime sessions and resources in AgentCore Runtime.
//
// This configuration helps optimize resource utilization by automatically cleaning up idle sessions and preventing
// long-running instances from consuming resources indefinitely.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleConfiguration := &LifecycleConfiguration{
//   	IdleRuntimeSessionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	MaxLifetime: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type LifecycleConfiguration struct {
	// Timeout in seconds for idle runtime sessions.
	//
	// When a session remains idle for this duration,
	// it will be automatically terminated.
	// Default: undefined - service default setting is 900 seconds (15 minutes).
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	IdleRuntimeSessionTimeout awscdk.Duration `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Maximum lifetime for the instance in seconds.
	//
	// Once reached, instances will be automatically
	// terminated and replaced.
	// Default: undefined - service default setting is 28800 seconds (8 hours).
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MaxLifetime awscdk.Duration `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

