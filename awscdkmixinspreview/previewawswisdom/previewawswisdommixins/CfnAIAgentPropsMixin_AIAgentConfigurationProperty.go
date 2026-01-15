package previewawswisdommixins


// A typed union that specifies the configuration based on the type of AI Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var annotations interface{}
//   var inputSchema interface{}
//   var outputSchema interface{}
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
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
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
//   	CaseSummarizationAiAgentConfiguration: &CaseSummarizationAIAgentConfigurationProperty{
//   		CaseSummarizationAiGuardrailId: jsii.String("caseSummarizationAiGuardrailId"),
//   		CaseSummarizationAiPromptId: jsii.String("caseSummarizationAiPromptId"),
//   		Locale: jsii.String("locale"),
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
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
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
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
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
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
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
//   	NoteTakingAiAgentConfiguration: &NoteTakingAIAgentConfigurationProperty{
//   		Locale: jsii.String("locale"),
//   		NoteTakingAiGuardrailId: jsii.String("noteTakingAiGuardrailId"),
//   		NoteTakingAiPromptId: jsii.String("noteTakingAiPromptId"),
//   	},
//   	OrchestrationAiAgentConfiguration: &OrchestrationAIAgentConfigurationProperty{
//   		ConnectInstanceArn: jsii.String("connectInstanceArn"),
//   		Locale: jsii.String("locale"),
//   		OrchestrationAiGuardrailId: jsii.String("orchestrationAiGuardrailId"),
//   		OrchestrationAiPromptId: jsii.String("orchestrationAiPromptId"),
//   		ToolConfigurations: []interface{}{
//   			&ToolConfigurationProperty{
//   				Annotations: annotations,
//   				Description: jsii.String("description"),
//   				InputSchema: inputSchema,
//   				Instruction: &ToolInstructionProperty{
//   					Examples: []*string{
//   						jsii.String("examples"),
//   					},
//   					Instruction: jsii.String("instruction"),
//   				},
//   				OutputFilters: []interface{}{
//   					&ToolOutputFilterProperty{
//   						JsonPath: jsii.String("jsonPath"),
//   						OutputConfiguration: &ToolOutputConfigurationProperty{
//   							OutputVariableNameOverride: jsii.String("outputVariableNameOverride"),
//   							SessionDataNamespace: jsii.String("sessionDataNamespace"),
//   						},
//   					},
//   				},
//   				OutputSchema: outputSchema,
//   				OverrideInputValues: []interface{}{
//   					&ToolOverrideInputValueProperty{
//   						JsonPath: jsii.String("jsonPath"),
//   						Value: &ToolOverrideInputValueConfigurationProperty{
//   							Constant: &ToolOverrideConstantInputValueProperty{
//   								Type: jsii.String("type"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Title: jsii.String("title"),
//   				ToolId: jsii.String("toolId"),
//   				ToolName: jsii.String("toolName"),
//   				ToolType: jsii.String("toolType"),
//   				UserInteractionConfiguration: &UserInteractionConfigurationProperty{
//   					IsUserConfirmationRequired: jsii.Boolean(false),
//   				},
//   			},
//   		},
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
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							OrConditions: []interface{}{
//   								&OrConditionProperty{
//   									AndConditions: []interface{}{
//   										&TagConditionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									TagCondition: &TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TagCondition: &TagConditionProperty{
//   								Key: jsii.String("key"),
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
type CfnAIAgentPropsMixin_AIAgentConfigurationProperty struct {
	// The configuration for AI Agents of type `ANSWER_RECOMMENDATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-answerrecommendationaiagentconfiguration
	//
	AnswerRecommendationAiAgentConfiguration interface{} `field:"optional" json:"answerRecommendationAiAgentConfiguration" yaml:"answerRecommendationAiAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-casesummarizationaiagentconfiguration
	//
	CaseSummarizationAiAgentConfiguration interface{} `field:"optional" json:"caseSummarizationAiAgentConfiguration" yaml:"caseSummarizationAiAgentConfiguration"`
	// Configuration for the EMAIL_GENERATIVE_ANSWER AI agent that provides comprehensive knowledge-based answers for customer queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-emailgenerativeansweraiagentconfiguration
	//
	EmailGenerativeAnswerAiAgentConfiguration interface{} `field:"optional" json:"emailGenerativeAnswerAiAgentConfiguration" yaml:"emailGenerativeAnswerAiAgentConfiguration"`
	// Configuration for the EMAIL_OVERVIEW AI agent that generates structured overview of email conversations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-emailoverviewaiagentconfiguration
	//
	EmailOverviewAiAgentConfiguration interface{} `field:"optional" json:"emailOverviewAiAgentConfiguration" yaml:"emailOverviewAiAgentConfiguration"`
	// Configuration for the EMAIL_RESPONSE AI agent that generates professional email responses using knowledge base content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-emailresponseaiagentconfiguration
	//
	EmailResponseAiAgentConfiguration interface{} `field:"optional" json:"emailResponseAiAgentConfiguration" yaml:"emailResponseAiAgentConfiguration"`
	// The configuration for AI Agents of type `MANUAL_SEARCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-manualsearchaiagentconfiguration
	//
	ManualSearchAiAgentConfiguration interface{} `field:"optional" json:"manualSearchAiAgentConfiguration" yaml:"manualSearchAiAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-notetakingaiagentconfiguration
	//
	NoteTakingAiAgentConfiguration interface{} `field:"optional" json:"noteTakingAiAgentConfiguration" yaml:"noteTakingAiAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-orchestrationaiagentconfiguration
	//
	OrchestrationAiAgentConfiguration interface{} `field:"optional" json:"orchestrationAiAgentConfiguration" yaml:"orchestrationAiAgentConfiguration"`
	// The self-service AI agent configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-selfserviceaiagentconfiguration
	//
	SelfServiceAiAgentConfiguration interface{} `field:"optional" json:"selfServiceAiAgentConfiguration" yaml:"selfServiceAiAgentConfiguration"`
}

