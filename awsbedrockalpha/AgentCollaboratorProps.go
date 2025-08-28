package awsbedrockalpha


// ****************************************************************************                   PROPS - Agent Collaborator Class ***************************************************************************.
//
// Example:
//   // Create a specialized agent
//   customerSupportAgent := bedrock.NewAgent(this, jsii.String("CustomerSupportAgent"), &AgentProps{
//   	Instruction: jsii.String("You specialize in answering customer support questions."),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   })
//
//   // Create an agent alias
//   customerSupportAlias := bedrock.NewAgentAlias(this, jsii.String("CustomerSupportAlias"), &AgentAliasProps{
//   	Agent: customerSupportAgent,
//   	AgentAliasName: jsii.String("production"),
//   })
//
//   // Create a main agent that collaborates with the specialized agent
//   mainAgent := bedrock.NewAgent(this, jsii.String("MainAgent"), &AgentProps{
//   	Instruction: jsii.String("You route specialized questions to other agents."),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	AgentCollaboration: map[string]interface{}{
//   		"type": bedrock.AgentCollaboratorType_SUPERVISOR,
//   		"collaborators": []AgentCollaborator{
//   			bedrock.NewAgentCollaborator(&AgentCollaboratorProps{
//   				"agentAlias": customerSupportAlias,
//   				"collaborationInstruction": jsii.String("Route customer support questions to this agent."),
//   				"collaboratorName": jsii.String("CustomerSupport"),
//   				"relayConversationHistory": jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type AgentCollaboratorProps struct {
	// Descriptor for the collaborating agent.
	//
	// This cannot be the TSTALIASID (`agent.testAlias`).
	// Experimental.
	AgentAlias IAgentAlias `field:"required" json:"agentAlias" yaml:"agentAlias"`
	// Instructions on how this agent should collaborate with the main agent.
	// Experimental.
	CollaborationInstruction *string `field:"required" json:"collaborationInstruction" yaml:"collaborationInstruction"`
	// A friendly name for the collaborator.
	// Experimental.
	CollaboratorName *string `field:"required" json:"collaboratorName" yaml:"collaboratorName"`
	// Whether to relay conversation history to this collaborator.
	// Default: - undefined (uses service default).
	//
	// Experimental.
	RelayConversationHistory *bool `field:"optional" json:"relayConversationHistory" yaml:"relayConversationHistory"`
}

