package previewawswisdommixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswisdom/previewawswisdommixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q in Connect AI Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAIAgentPropsMixin := awscdkmixinspreview.Mixins.NewCfnAIAgentPropsMixin(&CfnAIAgentMixinProps{
//   	AssistantId: jsii.String("assistantId"),
//   	Configuration: &AIAgentConfigurationProperty{
//   		AnswerRecommendationAiAgentConfiguration: &AnswerRecommendationAIAgentConfigurationProperty{
//   			AnswerGenerationAiGuardrailId: jsii.String("answerGenerationAiGuardrailId"),
//   			AnswerGenerationAiPromptId: jsii.String("answerGenerationAiPromptId"),
//   			AssociationConfigurations: []interface{}{
//   				&AssociationConfigurationProperty{
//   					AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   						KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   							ContentTagFilter: &TagFilterProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MaxResults: jsii.Number(123),
//   							OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   						},
//   					},
//   					AssociationId: jsii.String("associationId"),
//   					AssociationType: jsii.String("associationType"),
//   				},
//   			},
//   			IntentLabelingGenerationAiPromptId: jsii.String("intentLabelingGenerationAiPromptId"),
//   			Locale: jsii.String("locale"),
//   			QueryReformulationAiPromptId: jsii.String("queryReformulationAiPromptId"),
//   		},
//   		EmailGenerativeAnswerAiAgentConfiguration: &EmailGenerativeAnswerAIAgentConfigurationProperty{
//   			AssociationConfigurations: []interface{}{
//   				&AssociationConfigurationProperty{
//   					AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   						KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   							ContentTagFilter: &TagFilterProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MaxResults: jsii.Number(123),
//   							OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   						},
//   					},
//   					AssociationId: jsii.String("associationId"),
//   					AssociationType: jsii.String("associationType"),
//   				},
//   			},
//   			EmailGenerativeAnswerAiPromptId: jsii.String("emailGenerativeAnswerAiPromptId"),
//   			EmailQueryReformulationAiPromptId: jsii.String("emailQueryReformulationAiPromptId"),
//   			Locale: jsii.String("locale"),
//   		},
//   		EmailOverviewAiAgentConfiguration: &EmailOverviewAIAgentConfigurationProperty{
//   			EmailOverviewAiPromptId: jsii.String("emailOverviewAiPromptId"),
//   			Locale: jsii.String("locale"),
//   		},
//   		EmailResponseAiAgentConfiguration: &EmailResponseAIAgentConfigurationProperty{
//   			AssociationConfigurations: []interface{}{
//   				&AssociationConfigurationProperty{
//   					AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   						KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   							ContentTagFilter: &TagFilterProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MaxResults: jsii.Number(123),
//   							OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   						},
//   					},
//   					AssociationId: jsii.String("associationId"),
//   					AssociationType: jsii.String("associationType"),
//   				},
//   			},
//   			EmailQueryReformulationAiPromptId: jsii.String("emailQueryReformulationAiPromptId"),
//   			EmailResponseAiPromptId: jsii.String("emailResponseAiPromptId"),
//   			Locale: jsii.String("locale"),
//   		},
//   		ManualSearchAiAgentConfiguration: &ManualSearchAIAgentConfigurationProperty{
//   			AnswerGenerationAiGuardrailId: jsii.String("answerGenerationAiGuardrailId"),
//   			AnswerGenerationAiPromptId: jsii.String("answerGenerationAiPromptId"),
//   			AssociationConfigurations: []interface{}{
//   				&AssociationConfigurationProperty{
//   					AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   						KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   							ContentTagFilter: &TagFilterProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MaxResults: jsii.Number(123),
//   							OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   						},
//   					},
//   					AssociationId: jsii.String("associationId"),
//   					AssociationType: jsii.String("associationType"),
//   				},
//   			},
//   			Locale: jsii.String("locale"),
//   		},
//   		SelfServiceAiAgentConfiguration: &SelfServiceAIAgentConfigurationProperty{
//   			AssociationConfigurations: []interface{}{
//   				&AssociationConfigurationProperty{
//   					AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   						KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   							ContentTagFilter: &TagFilterProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MaxResults: jsii.Number(123),
//   							OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   						},
//   					},
//   					AssociationId: jsii.String("associationId"),
//   					AssociationType: jsii.String("associationType"),
//   				},
//   			},
//   			SelfServiceAiGuardrailId: jsii.String("selfServiceAiGuardrailId"),
//   			SelfServiceAnswerGenerationAiPromptId: jsii.String("selfServiceAnswerGenerationAiPromptId"),
//   			SelfServicePreProcessingAiPromptId: jsii.String("selfServicePreProcessingAiPromptId"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html
//
type CfnAIAgentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAIAgentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAIAgentPropsMixin
type jsiiProxy_CfnAIAgentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAIAgentPropsMixin) Props() *CfnAIAgentMixinProps {
	var returns *CfnAIAgentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAIAgentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::AIAgent`.
func NewCfnAIAgentPropsMixin(props *CfnAIAgentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAIAgentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAIAgentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAIAgentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAIAgentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::AIAgent`.
func NewCfnAIAgentPropsMixin_Override(c CfnAIAgentPropsMixin, props *CfnAIAgentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAIAgentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAIAgentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAIAgentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAIAgentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAIAgentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAIAgentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAIAgentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAIAgentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

