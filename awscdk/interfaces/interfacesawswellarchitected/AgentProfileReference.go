package interfacesawswellarchitected


// A reference to a AgentProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentProfileReference := &AgentProfileReference{
//   	AgentProfileArn: jsii.String("agentProfileArn"),
//   }
//
type AgentProfileReference struct {
	// The Arn of the AgentProfile resource.
	AgentProfileArn *string `field:"required" json:"agentProfileArn" yaml:"agentProfileArn"`
}

