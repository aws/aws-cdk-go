package awswisdom


// A typed union that specifies the configuration based on the type of AI Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIAgentConfigurationProperty := &AIAgentConfigurationProperty{
//   	AnswerRecommendationAiAgentConfiguration: &AnswerRecommendationAIAgentConfigurationProperty{
//   		AnswerGenerationAiGuardrailId: jsii.String("answerGenerationAiGuardrailId"),
//   		AnswerGenerationAiPromptId: jsii.String("answerGenerationAiPromptId"),
//   		AssociationConfigurations: []interface{}{
//   			&AssociationConfigurationProperty{
//   				AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   					KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   						ContentTagFilter: &TagFilterProperty{
//   							AndConditions: []interface{}{
//   								&TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MaxResults: jsii.Number(123),
//   						OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   					},
//   				},
//   				AssociationId: jsii.String("associationId"),
//   				AssociationType: jsii.String("associationType"),
//   			},
//   		},
//   		IntentLabelingGenerationAiPromptId: jsii.String("intentLabelingGenerationAiPromptId"),
//   		Locale: jsii.String("locale"),
//   		QueryReformulationAiPromptId: jsii.String("queryReformulationAiPromptId"),
//   	},
//   	EmailGenerativeAnswerAiAgentConfiguration: &EmailGenerativeAnswerAIAgentConfigurationProperty{
//   		AssociationConfigurations: []interface{}{
//   			&AssociationConfigurationProperty{
//   				AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   					KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   						ContentTagFilter: &TagFilterProperty{
//   							AndConditions: []interface{}{
//   								&TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MaxResults: jsii.Number(123),
//   						OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   					},
//   				},
//   				AssociationId: jsii.String("associationId"),
//   				AssociationType: jsii.String("associationType"),
//   			},
//   		},
//   		EmailGenerativeAnswerAiPromptId: jsii.String("emailGenerativeAnswerAiPromptId"),
//   		EmailQueryReformulationAiPromptId: jsii.String("emailQueryReformulationAiPromptId"),
//   		Locale: jsii.String("locale"),
//   	},
//   	EmailOverviewAiAgentConfiguration: &EmailOverviewAIAgentConfigurationProperty{
//   		EmailOverviewAiPromptId: jsii.String("emailOverviewAiPromptId"),
//   		Locale: jsii.String("locale"),
//   	},
//   	EmailResponseAiAgentConfiguration: &EmailResponseAIAgentConfigurationProperty{
//   		AssociationConfigurations: []interface{}{
//   			&AssociationConfigurationProperty{
//   				AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   					KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   						ContentTagFilter: &TagFilterProperty{
//   							AndConditions: []interface{}{
//   								&TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MaxResults: jsii.Number(123),
//   						OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   					},
//   				},
//   				AssociationId: jsii.String("associationId"),
//   				AssociationType: jsii.String("associationType"),
//   			},
//   		},
//   		EmailQueryReformulationAiPromptId: jsii.String("emailQueryReformulationAiPromptId"),
//   		EmailResponseAiPromptId: jsii.String("emailResponseAiPromptId"),
//   		Locale: jsii.String("locale"),
//   	},
//   	ManualSearchAiAgentConfiguration: &ManualSearchAIAgentConfigurationProperty{
//   		AnswerGenerationAiGuardrailId: jsii.String("answerGenerationAiGuardrailId"),
//   		AnswerGenerationAiPromptId: jsii.String("answerGenerationAiPromptId"),
//   		AssociationConfigurations: []interface{}{
//   			&AssociationConfigurationProperty{
//   				AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   					KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   						ContentTagFilter: &TagFilterProperty{
//   							AndConditions: []interface{}{
//   								&TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MaxResults: jsii.Number(123),
//   						OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   					},
//   				},
//   				AssociationId: jsii.String("associationId"),
//   				AssociationType: jsii.String("associationType"),
//   			},
//   		},
//   		Locale: jsii.String("locale"),
//   	},
//   	SelfServiceAiAgentConfiguration: &SelfServiceAIAgentConfigurationProperty{
//   		AssociationConfigurations: []interface{}{
//   			&AssociationConfigurationProperty{
//   				AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   					KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   						ContentTagFilter: &TagFilterProperty{
//   							AndConditions: []interface{}{
//   								&TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MaxResults: jsii.Number(123),
//   						OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   					},
//   				},
//   				AssociationId: jsii.String("associationId"),
//   				AssociationType: jsii.String("associationType"),
//   			},
//   		},
//   		SelfServiceAiGuardrailId: jsii.String("selfServiceAiGuardrailId"),
//   		SelfServiceAnswerGenerationAiPromptId: jsii.String("selfServiceAnswerGenerationAiPromptId"),
//   		SelfServicePreProcessingAiPromptId: jsii.String("selfServicePreProcessingAiPromptId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html
//
type CfnAIAgent_AIAgentConfigurationProperty struct {
	// The configuration for AI Agents of type `ANSWER_RECOMMENDATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-answerrecommendationaiagentconfiguration
	//
	AnswerRecommendationAiAgentConfiguration interface{} `field:"optional" json:"answerRecommendationAiAgentConfiguration" yaml:"answerRecommendationAiAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-emailgenerativeansweraiagentconfiguration
	//
	EmailGenerativeAnswerAiAgentConfiguration interface{} `field:"optional" json:"emailGenerativeAnswerAiAgentConfiguration" yaml:"emailGenerativeAnswerAiAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-emailoverviewaiagentconfiguration
	//
	EmailOverviewAiAgentConfiguration interface{} `field:"optional" json:"emailOverviewAiAgentConfiguration" yaml:"emailOverviewAiAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-emailresponseaiagentconfiguration
	//
	EmailResponseAiAgentConfiguration interface{} `field:"optional" json:"emailResponseAiAgentConfiguration" yaml:"emailResponseAiAgentConfiguration"`
	// The configuration for AI Agents of type `MANUAL_SEARCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-manualsearchaiagentconfiguration
	//
	ManualSearchAiAgentConfiguration interface{} `field:"optional" json:"manualSearchAiAgentConfiguration" yaml:"manualSearchAiAgentConfiguration"`
	// The self-service AI agent configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-selfserviceaiagentconfiguration
	//
	SelfServiceAiAgentConfiguration interface{} `field:"optional" json:"selfServiceAiAgentConfiguration" yaml:"selfServiceAiAgentConfiguration"`
}

