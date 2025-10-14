package awsbedrockagentcore


// A reference to a Runtime resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeReference := &RuntimeReference{
//   	AgentRuntimeId: jsii.String("agentRuntimeId"),
//   }
//
type RuntimeReference struct {
	// The AgentRuntimeId of the Runtime resource.
	AgentRuntimeId *string `field:"required" json:"agentRuntimeId" yaml:"agentRuntimeId"`
}

