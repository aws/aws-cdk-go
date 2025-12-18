package awswisdom


// Configuration settings for the EMAIL_RESPONSE AI agent including prompts, locale, and knowledge base associations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailResponseAIAgentConfigurationProperty := &EmailResponseAIAgentConfigurationProperty{
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
//   	EmailQueryReformulationAiPromptId: jsii.String("emailQueryReformulationAiPromptId"),
//   	EmailResponseAiPromptId: jsii.String("emailResponseAiPromptId"),
//   	Locale: jsii.String("locale"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.html
//
type CfnAIAgent_EmailResponseAIAgentConfigurationProperty struct {
	// Configuration settings for knowledge base associations used by the email response agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.html#cfn-wisdom-aiagent-emailresponseaiagentconfiguration-associationconfigurations
	//
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// The ID of the System AI prompt used for reformulating email queries to optimize knowledge base search for response generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.html#cfn-wisdom-aiagent-emailresponseaiagentconfiguration-emailqueryreformulationaipromptid
	//
	EmailQueryReformulationAiPromptId *string `field:"optional" json:"emailQueryReformulationAiPromptId" yaml:"emailQueryReformulationAiPromptId"`
	// The ID of the System AI prompt used for generating professional email responses based on knowledge base content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.html#cfn-wisdom-aiagent-emailresponseaiagentconfiguration-emailresponseaipromptid
	//
	EmailResponseAiPromptId *string `field:"optional" json:"emailResponseAiPromptId" yaml:"emailResponseAiPromptId"`
	// The locale setting for language-specific email response generation (for example, en_US, es_ES).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.html#cfn-wisdom-aiagent-emailresponseaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

