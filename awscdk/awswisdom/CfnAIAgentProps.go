package awswisdom


// Properties for defining a `CfnAIAgent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var annotations interface{}
//   var inputSchema interface{}
//   var outputSchema interface{}
//
//   cfnAIAgentProps := &CfnAIAgentProps{
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
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
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
//   		CaseSummarizationAiAgentConfiguration: &CaseSummarizationAIAgentConfigurationProperty{
//   			CaseSummarizationAiGuardrailId: jsii.String("caseSummarizationAiGuardrailId"),
//   			CaseSummarizationAiPromptId: jsii.String("caseSummarizationAiPromptId"),
//   			Locale: jsii.String("locale"),
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
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
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
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
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
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
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
//   		NoteTakingAiAgentConfiguration: &NoteTakingAIAgentConfigurationProperty{
//   			Locale: jsii.String("locale"),
//   			NoteTakingAiGuardrailId: jsii.String("noteTakingAiGuardrailId"),
//   			NoteTakingAiPromptId: jsii.String("noteTakingAiPromptId"),
//   		},
//   		OrchestrationAiAgentConfiguration: &OrchestrationAIAgentConfigurationProperty{
//   			OrchestrationAiPromptId: jsii.String("orchestrationAiPromptId"),
//
//   			// the properties below are optional
//   			ConnectInstanceArn: jsii.String("connectInstanceArn"),
//   			Locale: jsii.String("locale"),
//   			OrchestrationAiGuardrailId: jsii.String("orchestrationAiGuardrailId"),
//   			ToolConfigurations: []interface{}{
//   				&ToolConfigurationProperty{
//   					ToolName: jsii.String("toolName"),
//   					ToolType: jsii.String("toolType"),
//
//   					// the properties below are optional
//   					Annotations: annotations,
//   					Description: jsii.String("description"),
//   					InputSchema: inputSchema,
//   					Instruction: &ToolInstructionProperty{
//   						Examples: []*string{
//   							jsii.String("examples"),
//   						},
//   						Instruction: jsii.String("instruction"),
//   					},
//   					OutputFilters: []interface{}{
//   						&ToolOutputFilterProperty{
//   							JsonPath: jsii.String("jsonPath"),
//
//   							// the properties below are optional
//   							OutputConfiguration: &ToolOutputConfigurationProperty{
//   								OutputVariableNameOverride: jsii.String("outputVariableNameOverride"),
//   								SessionDataNamespace: jsii.String("sessionDataNamespace"),
//   							},
//   						},
//   					},
//   					OutputSchema: outputSchema,
//   					OverrideInputValues: []interface{}{
//   						&ToolOverrideInputValueProperty{
//   							JsonPath: jsii.String("jsonPath"),
//   							Value: &ToolOverrideInputValueConfigurationProperty{
//   								Constant: &ToolOverrideConstantInputValueProperty{
//   									Type: jsii.String("type"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Title: jsii.String("title"),
//   					ToolId: jsii.String("toolId"),
//   					UserInteractionConfiguration: &UserInteractionConfigurationProperty{
//   						IsUserConfirmationRequired: jsii.Boolean(false),
//   					},
//   				},
//   			},
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
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								OrConditions: []interface{}{
//   									&OrConditionProperty{
//   										AndConditions: []interface{}{
//   											&TagConditionProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										TagCondition: &TagConditionProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
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
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html
//
type CfnAIAgentProps struct {
	// The identifier of the Amazon Q in Connect assistant.
	//
	// Can be either the ID or the ARN. URLs cannot contain the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html#cfn-wisdom-aiagent-assistantid
	//
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// Configuration for the AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html#cfn-wisdom-aiagent-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The type of the AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html#cfn-wisdom-aiagent-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html#cfn-wisdom-aiagent-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html#cfn-wisdom-aiagent-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagent.html#cfn-wisdom-aiagent-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

