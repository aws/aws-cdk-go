package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Managed memory strategy that handles both built-in and override configurations.
//
// This strategy can be used for quick setup with built-in defaults or customized
// with specific models and prompt templates.
//
// Example:
//   // Create a custom semantic memory strategy
//   customSemanticStrategy := agentcore.MemoryStrategy_UsingSemantic(&ManagedStrategyProps{
//   	Name: jsii.String("customSemanticStrategy"),
//   	Description: jsii.String("Custom semantic memory strategy"),
//   	Namespaces: []*string{
//   		jsii.String("/custom/strategies/{memoryStrategyId}/actors/{actorId}"),
//   	},
//   	CustomConsolidation: &OverrideConfig{
//   		Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   		AppendToPrompt: jsii.String("Custom consolidation prompt for semantic memory"),
//   	},
//   	CustomExtraction: &OverrideConfig{
//   		Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   		AppendToPrompt: jsii.String("Custom extraction prompt for semantic memory"),
//   	},
//   })
//
//   // Create memory with custom strategy
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my-custom-memory"),
//   	Description: jsii.String("Memory with custom strategy"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		customSemanticStrategy,
//   	},
//   })
//
// Experimental.
type ManagedMemoryStrategy interface {
	IMemoryStrategy
	// The configuration for the custom consolidation.
	// Experimental.
	ConsolidationOverride() *OverrideConfig
	// The description of the memory strategy.
	// Experimental.
	Description() *string
	// The configuration for the custom extraction.
	// Experimental.
	ExtractionOverride() *OverrideConfig
	// The name of the memory strategy.
	// Experimental.
	Name() *string
	// The namespaces for the strategy.
	// Experimental.
	Namespaces() *[]*string
	// The configuration for episodic reflection.
	// Experimental.
	ReflectionConfiguration() *EpisodicReflectionConfiguration
	// The type of memory strategy.
	// Experimental.
	StrategyType() MemoryStrategyType
	// Grants the necessary permissions to the role.
	//
	// [disable-awslint:no-grants].
	//
	// Returns: The Grant object for chaining.
	// Experimental.
	Grant(grantee awsiam.IGrantable) awsiam.Grant
	// Renders the network configuration as a CloudFormation property.
	//
	// Returns: The CloudFormation property for the memory strategy.
	// Experimental.
	Render() *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty
}

// The jsii proxy struct for ManagedMemoryStrategy
type jsiiProxy_ManagedMemoryStrategy struct {
	jsiiProxy_IMemoryStrategy
}

func (j *jsiiProxy_ManagedMemoryStrategy) ConsolidationOverride() *OverrideConfig {
	var returns *OverrideConfig
	_jsii_.Get(
		j,
		"consolidationOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedMemoryStrategy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedMemoryStrategy) ExtractionOverride() *OverrideConfig {
	var returns *OverrideConfig
	_jsii_.Get(
		j,
		"extractionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedMemoryStrategy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedMemoryStrategy) Namespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"namespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedMemoryStrategy) ReflectionConfiguration() *EpisodicReflectionConfiguration {
	var returns *EpisodicReflectionConfiguration
	_jsii_.Get(
		j,
		"reflectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedMemoryStrategy) StrategyType() MemoryStrategyType {
	var returns MemoryStrategyType
	_jsii_.Get(
		j,
		"strategyType",
		&returns,
	)
	return returns
}


// Constructor to create a new managed memory strategy.
// Experimental.
func NewManagedMemoryStrategy(strategyType MemoryStrategyType, props *ManagedStrategyProps) ManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateNewManagedMemoryStrategyParameters(strategyType, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedMemoryStrategy{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ManagedMemoryStrategy",
		[]interface{}{strategyType, props},
		&j,
	)

	return &j
}

// Constructor to create a new managed memory strategy.
// Experimental.
func NewManagedMemoryStrategy_Override(m ManagedMemoryStrategy, strategyType MemoryStrategyType, props *ManagedStrategyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ManagedMemoryStrategy",
		[]interface{}{strategyType, props},
		m,
	)
}

func (m *jsiiProxy_ManagedMemoryStrategy) Grant(grantee awsiam.IGrantable) awsiam.Grant {
	if err := m.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		m,
		"grant",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedMemoryStrategy) Render() *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty {
	var returns *awsbedrockagentcore.CfnMemory_MemoryStrategyProperty

	_jsii_.Invoke(
		m,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

