package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// ****************************************************************************                        Agent Collaborator Class ***************************************************************************.
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
type AgentCollaborator interface {
	// The agent alias that this collaborator represents.
	//
	// This is the agent that will be called upon for collaboration.
	// Experimental.
	AgentAlias() IAgentAlias
	// Instructions on how this agent should collaborate with the main agent.
	// Experimental.
	CollaborationInstruction() *string
	// A friendly name for the collaborator.
	// Experimental.
	CollaboratorName() *string
	// Whether to relay conversation history to this collaborator.
	// Default: - undefined (uses service default).
	//
	// Experimental.
	RelayConversationHistory() *bool
	// Grants the given identity permissions to collaborate with the agent.
	//
	// Returns: The Grant object.
	// Experimental.
	Grant(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for AgentCollaborator
type jsiiProxy_AgentCollaborator struct {
	_ byte // padding
}

func (j *jsiiProxy_AgentCollaborator) AgentAlias() IAgentAlias {
	var returns IAgentAlias
	_jsii_.Get(
		j,
		"agentAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaborator) CollaborationInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaborator) CollaboratorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaborator) RelayConversationHistory() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"relayConversationHistory",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentCollaborator(props *AgentCollaboratorProps) AgentCollaborator {
	_init_.Initialize()

	if err := validateNewAgentCollaboratorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentCollaborator{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaborator",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentCollaborator_Override(a AgentCollaborator, props *AgentCollaboratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaborator",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AgentCollaborator) Grant(grantee awsiam.IGrantable) awsiam.Grant {
	if err := a.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grant",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

