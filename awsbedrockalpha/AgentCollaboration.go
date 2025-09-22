package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Class to manage agent collaboration configuration.
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
type AgentCollaboration interface {
	// Collaborators that this agent will work with.
	// Experimental.
	Collaborators() *[]AgentCollaborator
	// The collaboration type for the agent.
	// Experimental.
	Type() AgentCollaboratorType
}

// The jsii proxy struct for AgentCollaboration
type jsiiProxy_AgentCollaboration struct {
	_ byte // padding
}

func (j *jsiiProxy_AgentCollaboration) Collaborators() *[]AgentCollaborator {
	var returns *[]AgentCollaborator
	_jsii_.Get(
		j,
		"collaborators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaboration) Type() AgentCollaboratorType {
	var returns AgentCollaboratorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentCollaboration(config *AgentCollaborationConfig) AgentCollaboration {
	_init_.Initialize()

	if err := validateNewAgentCollaborationParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentCollaboration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaboration",
		[]interface{}{config},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentCollaboration_Override(a AgentCollaboration, config *AgentCollaborationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaboration",
		[]interface{}{config},
		a,
	)
}

