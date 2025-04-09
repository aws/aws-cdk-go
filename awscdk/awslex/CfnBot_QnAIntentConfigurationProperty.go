package awslex


// Details about the the configuration of the built-in `Amazon.QnAIntent` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qnAIntentConfigurationProperty := &QnAIntentConfigurationProperty{
//   	BedrockModelConfiguration: &BedrockModelSpecificationProperty{
//   		ModelArn: jsii.String("modelArn"),
//
//   		// the properties below are optional
//   		BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   			BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   			BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   		},
//   		BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   		BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
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
//   			QueryFilterStringEnabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			QueryFilterString: jsii.String("queryFilterString"),
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
type CfnBot_QnAIntentConfigurationProperty struct {
	// Contains information about the Amazon Bedrock model used to interpret the prompt used in descriptive bot building.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnaintentconfiguration.html#cfn-lex-bot-qnaintentconfiguration-bedrockmodelconfiguration
	//
	BedrockModelConfiguration interface{} `field:"required" json:"bedrockModelConfiguration" yaml:"bedrockModelConfiguration"`
	// Contains details about the configuration of the data source used for the `AMAZON.QnAIntent` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-qnaintentconfiguration.html#cfn-lex-bot-qnaintentconfiguration-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"required" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
}

