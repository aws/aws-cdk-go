package awswisdom


// The configuration for AI Agents of type `ANSWER_RECOMMENDATION` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   answerRecommendationAIAgentConfigurationProperty := &AnswerRecommendationAIAgentConfigurationProperty{
//   	AnswerGenerationAiGuardrailId: jsii.String("answerGenerationAiGuardrailId"),
//   	AnswerGenerationAiPromptId: jsii.String("answerGenerationAiPromptId"),
//   	AssociationConfigurations: []interface{}{
//   		&AssociationConfigurationProperty{
//   			AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   				KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   					ContentTagFilter: &TagFilterProperty{
//   						AndConditions: []interface{}{
//   							&TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						OrConditions: []interface{}{
//   							&OrConditionProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						TagCondition: &TagConditionProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					MaxResults: jsii.Number(123),
//   					OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   				},
//   			},
//   			AssociationId: jsii.String("associationId"),
//   			AssociationType: jsii.String("associationType"),
//   		},
//   	},
//   	IntentLabelingGenerationAiPromptId: jsii.String("intentLabelingGenerationAiPromptId"),
//   	Locale: jsii.String("locale"),
//   	QueryReformulationAiPromptId: jsii.String("queryReformulationAiPromptId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html
//
type CfnAIAgent_AnswerRecommendationAIAgentConfigurationProperty struct {
	// The ID of the answer generation AI guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-answergenerationaiguardrailid
	//
	AnswerGenerationAiGuardrailId *string `field:"optional" json:"answerGenerationAiGuardrailId" yaml:"answerGenerationAiGuardrailId"`
	// The AI Prompt identifier for the Answer Generation prompt used by the `ANSWER_RECOMMENDATION` AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-answergenerationaipromptid
	//
	AnswerGenerationAiPromptId *string `field:"optional" json:"answerGenerationAiPromptId" yaml:"answerGenerationAiPromptId"`
	// The association configurations for overriding behavior on this AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-associationconfigurations
	//
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// The AI Prompt identifier for the Intent Labeling prompt used by the `ANSWER_RECOMMENDATION` AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-intentlabelinggenerationaipromptid
	//
	IntentLabelingGenerationAiPromptId *string `field:"optional" json:"intentLabelingGenerationAiPromptId" yaml:"intentLabelingGenerationAiPromptId"`
	// The locale to which specifies the language and region settings that determine the response language for [QueryAssistant](https://docs.aws.amazon.com/connect/latest/APIReference/API_amazon-q-connect_QueryAssistant.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// The AI Prompt identifier for the Query Reformulation prompt used by the `ANSWER_RECOMMENDATION` AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-queryreformulationaipromptid
	//
	QueryReformulationAiPromptId *string `field:"optional" json:"queryReformulationAiPromptId" yaml:"queryReformulationAiPromptId"`
}

