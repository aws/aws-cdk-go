package awsbedrock


// Settings for parsing document contents.
//
// By default, the service converts the contents of each document into text before splitting it into chunks. To improve processing of PDF files with tables and images, you can configure the data source to convert the pages of text into images and use a model to describe the contents of each page.
//
// To use a model to parse PDF documents, set the parsing strategy to `BEDROCK_FOUNDATION_MODEL` and specify the model to use by ARN. You can also override the default parsing prompt with instructions for how to interpret images and tables in your documents. The following models are supported.
//
// - Anthropic Claude 3 Sonnet - `anthropic.claude-3-sonnet-20240229-v1:0`
// - Anthropic Claude 3 Haiku - `anthropic.claude-3-haiku-20240307-v1:0`
//
// You can get the ARN of a model with the [ListFoundationModels](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListFoundationModels.html) action. Standard model usage charges apply for the foundation model parsing strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parsingConfigurationProperty := &ParsingConfigurationProperty{
//   	ParsingStrategy: jsii.String("parsingStrategy"),
//
//   	// the properties below are optional
//   	BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   		ModelArn: jsii.String("modelArn"),
//
//   		// the properties below are optional
//   		ParsingPrompt: &ParsingPromptProperty{
//   			ParsingPromptText: jsii.String("parsingPromptText"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html
//
type CfnDataSource_ParsingConfigurationProperty struct {
	// The parsing strategy for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html#cfn-bedrock-datasource-parsingconfiguration-parsingstrategy
	//
	ParsingStrategy *string `field:"required" json:"parsingStrategy" yaml:"parsingStrategy"`
	// Settings for a foundation model used to parse documents for a data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-parsingconfiguration.html#cfn-bedrock-datasource-parsingconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

