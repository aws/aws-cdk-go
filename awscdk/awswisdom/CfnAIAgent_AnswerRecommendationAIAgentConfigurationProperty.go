package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   answerRecommendationAIAgentConfigurationProperty := &AnswerRecommendationAIAgentConfigurationProperty{
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
//   	QueryReformulationAiPromptId: jsii.String("queryReformulationAiPromptId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html
//
type CfnAIAgent_AnswerRecommendationAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-answergenerationaipromptid
	//
	AnswerGenerationAiPromptId *string `field:"optional" json:"answerGenerationAiPromptId" yaml:"answerGenerationAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-associationconfigurations
	//
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-intentlabelinggenerationaipromptid
	//
	IntentLabelingGenerationAiPromptId *string `field:"optional" json:"intentLabelingGenerationAiPromptId" yaml:"intentLabelingGenerationAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html#cfn-wisdom-aiagent-answerrecommendationaiagentconfiguration-queryreformulationaipromptid
	//
	QueryReformulationAiPromptId *string `field:"optional" json:"queryReformulationAiPromptId" yaml:"queryReformulationAiPromptId"`
}

