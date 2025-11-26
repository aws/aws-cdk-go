package previewawslexmixins


// Details about the the configuration of the built-in `Amazon.QnAIntent` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   qnAIntentConfigurationProperty := &QnAIntentConfigurationProperty{
//   	BedrockModelConfiguration: &BedrockModelSpecificationProperty{
//   		BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   			BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   			BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   		},
//   		BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   		BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   		ModelArn: jsii.String("modelArn"),
//   	},
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		BedrockKnowledgeStoreConfiguration: &BedrockKnowledgeStoreConfigurationProperty{
//   			BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
//   			BkbExactResponseFields: &BKBExactResponseFieldsProperty{
//   				AnswerField: jsii.String("answerField"),
//   			},
//   			ExactResponse: jsii.Boolean(false),
//   		},
//   		KendraConfiguration: &QnAKendraConfigurationProperty{
//   			ExactResponse: jsii.Boolean(false),
//   			KendraIndex: jsii.String("kendraIndex"),
//   			QueryFilterString: jsii.String("queryFilterString"),
//   			QueryFilterStringEnabled: jsii.Boolean(false),
//   		},
//   		OpensearchConfiguration: &OpensearchConfigurationProperty{
//   			DomainEndpoint: jsii.String("domainEndpoint"),
//   			ExactResponse: jsii.Boolean(false),
//   			ExactResponseFields: &ExactResponseFieldsProperty{
//   				AnswerField: jsii.String("answerField"),
//   				QuestionField: jsii.String("questionField"),
//   			},
//   			IncludeFields: []*string{
//   				jsii.String("includeFields"),
//   			},
//   			IndexName: jsii.String("indexName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnaintentconfiguration.html
//
type CfnBotPropsMixin_QnAIntentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnaintentconfiguration.html#cfn-lex-bot-qnaintentconfiguration-bedrockmodelconfiguration
	//
	BedrockModelConfiguration interface{} `field:"optional" json:"bedrockModelConfiguration" yaml:"bedrockModelConfiguration"`
	// Contains details about the configuration of the data source used for the `AMAZON.QnAIntent` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnaintentconfiguration.html#cfn-lex-bot-qnaintentconfiguration-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
}

