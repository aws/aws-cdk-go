package awswisdom


// Settings for parsing document contents.
//
// By default, the service converts the contents of each document into text before splitting it into chunks. To improve processing of PDF files with tables and images, you can configure the data source to convert the pages of text into images and use a model to describe the contents of each page.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html
//
type CfnKnowledgeBase_ParsingConfigurationProperty struct {
	// The parsing strategy for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html#cfn-wisdom-knowledgebase-parsingconfiguration-parsingstrategy
	//
	ParsingStrategy *string `field:"required" json:"parsingStrategy" yaml:"parsingStrategy"`
	// Settings for a foundation model used to parse documents for a data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html#cfn-wisdom-knowledgebase-parsingconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

