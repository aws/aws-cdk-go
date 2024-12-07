package awswisdom


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html#cfn-wisdom-knowledgebase-parsingconfiguration-parsingstrategy
	//
	ParsingStrategy *string `field:"required" json:"parsingStrategy" yaml:"parsingStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-parsingconfiguration.html#cfn-wisdom-knowledgebase-parsingconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

