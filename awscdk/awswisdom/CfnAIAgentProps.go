package awswisdom


// Properties for defining a `CfnAIAgent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   			QueryReformulationAiPromptId: jsii.String("queryReformulationAiPromptId"),
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

