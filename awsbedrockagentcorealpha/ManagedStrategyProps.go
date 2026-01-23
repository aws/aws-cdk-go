package awsbedrockagentcorealpha


// Configuration parameters for a memory strategy that can override existing built-in default prompts/models.
//
// Example:
//   // Create memory with custom strategies
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my_memory"),
//   	Description: jsii.String("Memory with custom strategies"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		agentcore.MemoryStrategy_UsingUserPreference(&ManagedStrategyProps{
//   			Name: jsii.String("CustomerPreferences"),
//   			Namespaces: []*string{
//   				jsii.String("support/customer/{actorId}/preferences"),
//   			},
//   		}),
//   		agentcore.MemoryStrategy_UsingSemantic(&ManagedStrategyProps{
//   			Name: jsii.String("CustomerSupportSemantic"),
//   			Namespaces: []*string{
//   				jsii.String("support/customer/{actorId}/semantic"),
//   			},
//   		}),
//   		agentcore.MemoryStrategy_UsingEpisodic(&ManagedStrategyProps{
//   			Name: jsii.String("customerJourneyEpisodic"),
//   			Namespaces: []*string{
//   				jsii.String("/journey/customer/{actorId}/episodes"),
//   			},
//   			ReflectionConfiguration: &EpisodicReflectionConfiguration{
//   				Namespaces: []*string{
//   					jsii.String("/journey/customer/{actorId}/reflections"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ManagedStrategyProps struct {
	// The name for the strategy.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the strategy.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The namespaces for the strategy Represents a namespace for organizing memory data Use a hierarchical format separated by forward slashes (/).
	//
	// Use a hierarchical format separated by forward slashes (/) to organize namespaces logically.
	// You can include these defined variables:
	//
	// - {sessionId} - the user identifier to be created in the CreateEvent API
	// - {memoryStrategyId} - an identifier for an extraction strategy
	// - {sessionId} - an identifier for each session
	//
	// Example namespace path:
	// /strategies/{memoryStrategyId}/actions/{actionId}/sessions/{sessionId}
	//
	// After memory creation, this namespace might look like:
	// /actor/actor-3afc5aa8fef9/strategy/summarization-fy5c5fwc7/session/session-qj7tpd1kvr8.
	// Experimental.
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
	// The configuration for the custom consolidation.
	//
	// This configuration provides customization to how the model identifies
	// and extracts relevant information for memory storage.
	// Default: - No custom extraction.
	//
	// Experimental.
	CustomConsolidation *OverrideConfig `field:"optional" json:"customConsolidation" yaml:"customConsolidation"`
	// The configuration for the custom extraction.
	//
	// This configuration provides customization to how the model identifies
	// and extracts relevant information for memory storage.
	// Default: - No custom extraction.
	//
	// Experimental.
	CustomExtraction *OverrideConfig `field:"optional" json:"customExtraction" yaml:"customExtraction"`
	// Configuration for episodic memory reflection (only applicable for EPISODIC strategy type) Defines additional namespaces for reflection-based episodic recall.
	// Default: - No reflection configuration.
	//
	// Experimental.
	ReflectionConfiguration *EpisodicReflectionConfiguration `field:"optional" json:"reflectionConfiguration" yaml:"reflectionConfiguration"`
}

