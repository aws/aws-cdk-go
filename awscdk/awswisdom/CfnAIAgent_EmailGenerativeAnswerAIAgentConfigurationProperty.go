package awswisdom


// Configuration settings for the EMAIL_GENERATIVE_ANSWER AI agent including prompts, locale, and knowledge base associations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailGenerativeAnswerAIAgentConfigurationProperty := &EmailGenerativeAnswerAIAgentConfigurationProperty{
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
//   	EmailGenerativeAnswerAiPromptId: jsii.String("emailGenerativeAnswerAiPromptId"),
//   	EmailQueryReformulationAiPromptId: jsii.String("emailQueryReformulationAiPromptId"),
//   	Locale: jsii.String("locale"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html
//
type CfnAIAgent_EmailGenerativeAnswerAIAgentConfigurationProperty struct {
	// Configuration settings for knowledge base associations used by the email generative answer agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-associationconfigurations
	//
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// The ID of the System AI prompt used for generating comprehensive knowledge-based answers from email queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-emailgenerativeansweraipromptid
	//
	EmailGenerativeAnswerAiPromptId *string `field:"optional" json:"emailGenerativeAnswerAiPromptId" yaml:"emailGenerativeAnswerAiPromptId"`
	// The ID of the System AI prompt used for reformulating email queries to optimize knowledge base search results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-emailqueryreformulationaipromptid
	//
	EmailQueryReformulationAiPromptId *string `field:"optional" json:"emailQueryReformulationAiPromptId" yaml:"emailQueryReformulationAiPromptId"`
	// The locale setting for language-specific email processing and response generation (for example, en_US, es_ES).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

