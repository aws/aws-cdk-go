package interfacesawsbedrock


// A reference to a Agent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentReference := &AgentReference{
//   	AgentArn: jsii.String("agentArn"),
//   	AgentId: jsii.String("agentId"),
//   }
//
type AgentReference struct {
	// The ARN of the Agent resource.
	AgentArn *string `field:"required" json:"agentArn" yaml:"agentArn"`
	// The AgentId of the Agent resource.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
}

