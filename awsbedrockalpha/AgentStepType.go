package awsbedrockalpha


// The step in the agent sequence that this prompt configuration applies to.
//
// Example:
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	Instruction: jsii.String("You are a helpful assistant."),
//   	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_FromSteps([]PromptStepConfigBase{
//   		&PromptRoutingClassifierConfigCustomParser{
//   			StepType: bedrock.AgentStepType_ROUTING_CLASSIFIER,
//   			StepEnabled: jsii.Boolean(true),
//   			CustomPromptTemplate: jsii.String("Your routing template here"),
//   			FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_V2(),
//   		}.(PromptRoutingClassifierConfigCustomParser),
//   	}),
//   })
//
// Experimental.
type AgentStepType string

const (
	// Pre-processing step that prepares the user input for orchestration.
	// Experimental.
	AgentStepType_PRE_PROCESSING AgentStepType = "PRE_PROCESSING"
	// Main orchestration step that determines the agent's actions.
	// Experimental.
	AgentStepType_ORCHESTRATION AgentStepType = "ORCHESTRATION"
	// Post-processing step that refines the agent's response.
	// Experimental.
	AgentStepType_POST_PROCESSING AgentStepType = "POST_PROCESSING"
	// Step that classifies and routes requests to appropriate collaborators.
	// Experimental.
	AgentStepType_ROUTING_CLASSIFIER AgentStepType = "ROUTING_CLASSIFIER"
	// Step that summarizes conversation history for memory retention.
	// Experimental.
	AgentStepType_MEMORY_SUMMARIZATION AgentStepType = "MEMORY_SUMMARIZATION"
	// Step that generates responses using knowledge base content.
	// Experimental.
	AgentStepType_KNOWLEDGE_BASE_RESPONSE_GENERATION AgentStepType = "KNOWLEDGE_BASE_RESPONSE_GENERATION"
)

