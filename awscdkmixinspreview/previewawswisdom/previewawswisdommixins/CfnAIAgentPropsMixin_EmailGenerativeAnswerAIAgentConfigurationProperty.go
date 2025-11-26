package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						OrConditions: []interface{}{
//   							&OrConditionProperty{
//   								AndConditions: []interface{}{
//   									&TagConditionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								TagCondition: &TagConditionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						TagCondition: &TagConditionProperty{
//   							Key: jsii.String("key"),
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
type CfnAIAgentPropsMixin_EmailGenerativeAnswerAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-associationconfigurations
	//
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-emailgenerativeansweraipromptid
	//
	EmailGenerativeAnswerAiPromptId *string `field:"optional" json:"emailGenerativeAnswerAiPromptId" yaml:"emailGenerativeAnswerAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-emailqueryreformulationaipromptid
	//
	EmailQueryReformulationAiPromptId *string `field:"optional" json:"emailQueryReformulationAiPromptId" yaml:"emailQueryReformulationAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html#cfn-wisdom-aiagent-emailgenerativeansweraiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

