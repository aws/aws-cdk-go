package awsbedrockalpha


// Configuration for agent collaboration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var agentCollaborator AgentCollaborator
//
//   agentCollaborationConfig := &AgentCollaborationConfig{
//   	Collaborators: []AgentCollaborator{
//   		agentCollaborator,
//   	},
//   	Type: bedrock_alpha.AgentCollaboratorType_SUPERVISOR,
//   }
//
// Experimental.
type AgentCollaborationConfig struct {
	// Collaborators that this agent will work with.
	// Experimental.
	Collaborators *[]AgentCollaborator `field:"required" json:"collaborators" yaml:"collaborators"`
	// The collaboration type for the agent.
	// Experimental.
	Type AgentCollaboratorType `field:"required" json:"type" yaml:"type"`
}

