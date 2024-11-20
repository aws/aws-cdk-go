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
//   		QueryReformulationAiPromptId: jsii.String("queryReformulationAiPromptId"),
//   	},
//   	ManualSearchAiAgentConfiguration: &ManualSearchAIAgentConfigurationProperty{
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
	// The configuration for AI Agents of type `MANUAL_SEARCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-aiagentconfiguration.html#cfn-wisdom-aiagent-aiagentconfiguration-manualsearchaiagentconfiguration
	//
	ManualSearchAiAgentConfiguration interface{} `field:"optional" json:"manualSearchAiAgentConfiguration" yaml:"manualSearchAiAgentConfiguration"`
}

