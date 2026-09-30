package interfacesawswellarchitected


// A reference to a AgentContext resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentContextReference := &AgentContextReference{
//   	AgentContextArn: jsii.String("agentContextArn"),
//   }
//
type AgentContextReference struct {
	// The Arn of the AgentContext resource.
	AgentContextArn *string `field:"required" json:"agentContextArn" yaml:"agentContextArn"`
}

