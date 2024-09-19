package awsbedrock


// Settings for a foundation model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) used to parse documents for a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockFoundationModelConfigurationProperty := &BedrockFoundationModelConfigurationProperty{
//   	ModelArn: jsii.String("modelArn"),
//
//   	// the properties below are optional
//   	ParsingPrompt: &ParsingPromptProperty{
//   		ParsingPromptText: jsii.String("parsingPromptText"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html
//
type CfnDataSource_BedrockFoundationModelConfigurationProperty struct {
	// The ARN of the foundation model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelconfiguration-modelarn
	//
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// Instructions for interpreting the contents of a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelconfiguration-parsingprompt
	//
	ParsingPrompt interface{} `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

