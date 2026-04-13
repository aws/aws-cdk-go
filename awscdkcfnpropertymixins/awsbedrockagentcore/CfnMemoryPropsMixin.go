package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Memory allows AI agents to maintain both immediate and long-term knowledge, enabling context-aware and personalized interactions.
//
// For more information about using Memory in Amazon Bedrock AgentCore, see [Host agent or tools with Amazon Bedrock AgentCore Memory](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/memory-getting-started.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMemoryPropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnMemoryPropsMixin(&CfnMemoryMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EventExpiryDuration: jsii.Number(123),
//   	MemoryExecutionRoleArn: jsii.String("memoryExecutionRoleArn"),
//   	MemoryStrategies: []interface{}{
//   		&MemoryStrategyProperty{
//   			CustomMemoryStrategy: &CustomMemoryStrategyProperty{
//   				Configuration: &CustomConfigurationInputProperty{
//   					EpisodicOverride: &EpisodicOverrideProperty{
//   						Consolidation: &EpisodicOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Extraction: &EpisodicOverrideExtractionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Reflection: &EpisodicOverrideReflectionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   							Namespaces: []*string{
//   								jsii.String("namespaces"),
//   							},
//   							NamespaceTemplates: []*string{
//   								jsii.String("namespaceTemplates"),
//   							},
//   						},
//   					},
//   					SelfManagedConfiguration: &SelfManagedConfigurationProperty{
//   						HistoricalContextWindowSize: jsii.Number(123),
//   						InvocationConfiguration: &InvocationConfigurationInputProperty{
//   							PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   							TopicArn: jsii.String("topicArn"),
//   						},
//   						TriggerConditions: []interface{}{
//   							&TriggerConditionInputProperty{
//   								MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   									MessageCount: jsii.Number(123),
//   								},
//   								TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   									IdleSessionTimeout: jsii.Number(123),
//   								},
//   								TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   									TokenCount: jsii.Number(123),
//   								},
//   							},
//   						},
//   					},
//   					SemanticOverride: &SemanticOverrideProperty{
//   						Consolidation: &SemanticOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Extraction: &SemanticOverrideExtractionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   					},
//   					SummaryOverride: &SummaryOverrideProperty{
//   						Consolidation: &SummaryOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   					},
//   					UserPreferenceOverride: &UserPreferenceOverrideProperty{
//   						Consolidation: &UserPreferenceOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Extraction: &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   					},
//   				},
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			EpisodicMemoryStrategy: &EpisodicMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				ReflectionConfiguration: &EpisodicReflectionConfigurationInputProperty{
//   					Namespaces: []*string{
//   						jsii.String("namespaces"),
//   					},
//   					NamespaceTemplates: []*string{
//   						jsii.String("namespaceTemplates"),
//   					},
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			SemanticMemoryStrategy: &SemanticMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			SummaryMemoryStrategy: &SummaryMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			UserPreferenceMemoryStrategy: &UserPreferenceMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	StreamDeliveryResources: &StreamDeliveryResourcesProperty{
//   		Resources: []interface{}{
//   			&StreamDeliveryResourceProperty{
//   				Kinesis: &KinesisResourceProperty{
//   					ContentConfigurations: []interface{}{
//   						&ContentConfigurationProperty{
//   							Level: jsii.String("level"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					DataStreamArn: jsii.String("dataStreamArn"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html
//
type CfnMemoryPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMemoryMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMemoryPropsMixin
type jsiiProxy_CfnMemoryPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMemoryPropsMixin) Props() *CfnMemoryMixinProps {
	var returns *CfnMemoryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMemoryPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::Memory`.
func NewCfnMemoryPropsMixin(props *CfnMemoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMemoryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMemoryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMemoryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::Memory`.
func NewCfnMemoryPropsMixin_Override(c CfnMemoryPropsMixin, props *CfnMemoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMemoryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMemoryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMemoryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnMemoryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMemoryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMemoryPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

