package previewawsbedrockmixins


// Settings for a foundation model used to parse documents for a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockFoundationModelConfigurationProperty := &BedrockFoundationModelConfigurationProperty{
//   	ModelArn: jsii.String("modelArn"),
//   	ParsingModality: jsii.String("parsingModality"),
//   	ParsingPrompt: &ParsingPromptProperty{
//   		ParsingPromptText: jsii.String("parsingPromptText"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html
//
type CfnDataSourcePropsMixin_BedrockFoundationModelConfigurationProperty struct {
	// The ARN of the foundation model to use for parsing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelconfiguration-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// Specifies whether to enable parsing of multimodal data, including both text and/or images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelconfiguration-parsingmodality
	//
	ParsingModality *string `field:"optional" json:"parsingModality" yaml:"parsingModality"`
	// Instructions for interpreting the contents of a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelconfiguration-parsingprompt
	//
	ParsingPrompt interface{} `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

