package interfacesawssecurityagent


// A reference to a AgentSpace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentSpaceReference := &AgentSpaceReference{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   }
//
type AgentSpaceReference struct {
	// The AgentSpaceId of the AgentSpace resource.
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
}

