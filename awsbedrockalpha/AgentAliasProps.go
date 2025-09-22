package awsbedrockalpha


// Properties for creating a CDK-Managed Agent Alias.
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
type AgentAliasProps struct {
	// The agent associated to this alias.
	// Experimental.
	Agent IAgent `field:"required" json:"agent" yaml:"agent"`
	// The name for the agent alias.
	//
	// This will be used as the physical name of the agent alias.
	// Default: - "latest".
	//
	// Experimental.
	AgentAliasName *string `field:"optional" json:"agentAliasName" yaml:"agentAliasName"`
	// The version of the agent to associate with the agent alias.
	// Default: - Creates a new version of the agent.
	//
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Description for the agent alias.
	// Default: undefined - No description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

