package previewawsbedrockmixins


// Settings for parsing document contents.
//
// If you exclude this field, the default parser converts the contents of each document into text before splitting it into chunks. Specify the parsing strategy to use in the `parsingStrategy` field and include the relevant configuration, or omit it to use the Amazon Bedrock default parser. For more information, see [Parsing options for your data source](https://docs.aws.amazon.com/bedrock/latest/userguide/kb-advanced-parsing.html) .
//
// > If you specify `BEDROCK_DATA_AUTOMATION` or `BEDROCK_FOUNDATION_MODEL` and it fails to parse a file, the Amazon Bedrock default parser will be used instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parsingConfigurationProperty := &ParsingConfigurationProperty{
//   	BedrockDataAutomationConfiguration: &BedrockDataAutomationConfigurationProperty{
//   		ParsingModality: jsii.String("parsingModality"),
//   	},
//   	BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   		ModelArn: jsii.String("modelArn"),
//   		ParsingModality: jsii.String("parsingModality"),
//   		ParsingPrompt: &ParsingPromptProperty{
//   			ParsingPromptText: jsii.String("parsingPromptText"),
//   		},
//   	},
//   	ParsingStrategy: jsii.String("parsingStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html
//
type CfnDataSourcePropsMixin_ParsingConfigurationProperty struct {
	// If you specify `BEDROCK_DATA_AUTOMATION` as the parsing strategy for ingesting your data source, use this object to modify configurations for using the Amazon Bedrock Data Automation parser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html#cfn-bedrock-datasource-parsingconfiguration-bedrockdataautomationconfiguration
	//
	BedrockDataAutomationConfiguration interface{} `field:"optional" json:"bedrockDataAutomationConfiguration" yaml:"bedrockDataAutomationConfiguration"`
	// If you specify `BEDROCK_FOUNDATION_MODEL` as the parsing strategy for ingesting your data source, use this object to modify configurations for using a foundation model to parse documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html#cfn-bedrock-datasource-parsingconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
	// The parsing strategy for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html#cfn-bedrock-datasource-parsingconfiguration-parsingstrategy
	//
	ParsingStrategy *string `field:"optional" json:"parsingStrategy" yaml:"parsingStrategy"`
}

