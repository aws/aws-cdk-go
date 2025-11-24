package mixinsawslex


// Contains details about the configuration of the knowledge store used for the `AMAZON.QnAIntent` . You must have already created the knowledge store and indexed the documents within it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   		QueryFilterString: jsii.String("queryFilterString"),
//   		QueryFilterStringEnabled: jsii.Boolean(false),
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
type CfnBotPropsMixin_DataSourceConfigurationProperty struct {
	// Contains details about the configuration of the Amazon Bedrock knowledge base used for the `AMAZON.QnAIntent` . To set up a knowledge base, follow the steps at [Building a knowledge base](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html#cfn-lex-bot-datasourceconfiguration-bedrockknowledgestoreconfiguration
	//
	BedrockKnowledgeStoreConfiguration interface{} `field:"optional" json:"bedrockKnowledgeStoreConfiguration" yaml:"bedrockKnowledgeStoreConfiguration"`
	// Contains details about the configuration of the Amazon Kendra index used for the `AMAZON.QnAIntent` . To create a Amazon Kendra index, follow the steps at [Creating an index](https://docs.aws.amazon.com/kendra/latest/dg/create-index.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html#cfn-lex-bot-datasourceconfiguration-kendraconfiguration
	//
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// Contains details about the configuration of the Amazon OpenSearch Service database used for the `AMAZON.QnAIntent` . To create a domain, follow the steps at [Creating and managing Amazon OpenSearch Service domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-datasourceconfiguration.html#cfn-lex-bot-datasourceconfiguration-opensearchconfiguration
	//
	OpensearchConfiguration interface{} `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
}

