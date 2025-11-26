package previewawswisdommixins


// The configuration of the Bedrock foundation model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockFoundationModelConfigurationProperty := &BedrockFoundationModelConfigurationProperty{
//   	ModelArn: jsii.String("modelArn"),
//   	ParsingPrompt: &ParsingPromptProperty{
//   		ParsingPromptText: jsii.String("parsingPromptText"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-bedrockfoundationmodelconfiguration.html
//
type CfnKnowledgeBasePropsMixin_BedrockFoundationModelConfigurationProperty struct {
	// The model ARN of the Bedrock foundation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-bedrockfoundationmodelconfiguration.html#cfn-wisdom-knowledgebase-bedrockfoundationmodelconfiguration-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// The parsing prompt of the Bedrock foundation model configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-bedrockfoundationmodelconfiguration.html#cfn-wisdom-knowledgebase-bedrockfoundationmodelconfiguration-parsingprompt
	//
	ParsingPrompt interface{} `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

