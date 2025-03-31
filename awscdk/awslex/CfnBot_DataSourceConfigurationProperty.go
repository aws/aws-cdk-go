package awslex


// Contains details about the configuration of the data source used for the AMAZON.QnAIntent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	BedrockKnowledgeStoreConfiguration: &BedrockKnowledgeStoreConfigurationProperty{
//   		BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
//   		BkbExactResponseFields: &BKBExactResponseFieldsProperty{
//   			AnswerField: jsii.String("answerField"),
//   		},
//   		ExactResponse: jsii.Boolean(false),
//   	},
//   	KendraConfiguration: &QnAKendraConfigurationProperty{
//   		ExactResponse: jsii.Boolean(false),
//   		KendraIndex: jsii.String("kendraIndex"),
//   		QueryFilterStringEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		QueryFilterString: jsii.String("queryFilterString"),
//   	},
//   	OpensearchConfiguration: &OpensearchConfigurationProperty{
//   		DomainEndpoint: jsii.String("domainEndpoint"),
//   		ExactResponse: jsii.Boolean(false),
//   		ExactResponseFields: &ExactResponseFieldsProperty{
//   			AnswerField: jsii.String("answerField"),
//   			QuestionField: jsii.String("questionField"),
//   		},
//   		IncludeFields: []*string{
//   			jsii.String("includeFields"),
//   		},
//   		IndexName: jsii.String("indexName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html
//
type CfnBot_DataSourceConfigurationProperty struct {
	// Contains details about the configuration of a Amazon Bedrock knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html#cfn-lex-bot-datasourceconfiguration-bedrockknowledgestoreconfiguration
	//
	BedrockKnowledgeStoreConfiguration interface{} `field:"optional" json:"bedrockKnowledgeStoreConfiguration" yaml:"bedrockKnowledgeStoreConfiguration"`
	// Contains details about the configuration of the Amazon Kendra index used for the AMAZON.QnAIntent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html#cfn-lex-bot-datasourceconfiguration-kendraconfiguration
	//
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// Contains details about the configuration of the Amazon OpenSearch Service database used for the AMAZON.QnAIntent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html#cfn-lex-bot-datasourceconfiguration-opensearchconfiguration
	//
	OpensearchConfiguration interface{} `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
}

