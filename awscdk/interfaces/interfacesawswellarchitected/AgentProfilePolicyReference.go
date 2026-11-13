package interfacesawswellarchitected


// A reference to a AgentProfilePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentProfilePolicyReference := &AgentProfilePolicyReference{
//   	ProfileArn: jsii.String("profileArn"),
//   }
//
type AgentProfilePolicyReference struct {
	// The ProfileArn of the AgentProfilePolicy resource.
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
}

