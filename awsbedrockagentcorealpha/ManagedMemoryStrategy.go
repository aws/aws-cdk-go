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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var bedrockInvokable IBedrockInvokable
//
//   managedMemoryStrategy := bedrock_agentcore_alpha.NewManagedMemoryStrategy(bedrock_agentcore_alpha.MemoryStrategyType_SUMMARIZATION, &ManagedStrategyProps{
//   	Name: jsii.String("name"),
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//
//   	// the properties below are optional
//   	CustomConsolidation: &OverrideConfig{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		Model: bedrockInvokable,
//   	},
//   	CustomExtraction: &OverrideConfig{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		Model: bedrockInvokable,
//   	},
//   	Description: jsii.String("description"),
//   	ReflectionConfiguration: &EpisodicReflectionConfiguration{
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   	},
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type ManagedMemoryStrategy interface {
	IMemoryStrategy
	// The configuration for the custom consolidation.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ConsolidationOverride() *OverrideConfig
	// The description of the memory strategy.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description() *string
	// The configuration for the custom extraction.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ExtractionOverride() *OverrideConfig
	// The name of the memory strategy.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Name() *string
	// The namespaces for the strategy.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Namespaces() *[]*string
	// The configuration for episodic reflection.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ReflectionConfiguration() *EpisodicReflectionConfiguration
	// The type of memory strategy.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	StrategyType() MemoryStrategyType
	// Grants the necessary permissions to the role.
	//
	// [disable-awslint:no-grants].
	//
	// Returns: The Grant object for chaining.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Grant(grantee awsiam.IGrantable) awsiam.Grant
	// Renders the network configuration as a CloudFormation property.
	//
	// Returns: The CloudFormation property for the memory strategy.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
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
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
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
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
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

