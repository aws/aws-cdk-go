package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockAgentIntentKnowledgeBaseConfigurationProperty := &BedrockAgentIntentKnowledgeBaseConfigurationProperty{
//   	BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentintentknowledgebaseconfiguration.html
//
type CfnBot_BedrockAgentIntentKnowledgeBaseConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentintentknowledgebaseconfiguration.html#cfn-lex-bot-bedrockagentintentknowledgebaseconfiguration-bedrockknowledgebasearn
	//
	BedrockKnowledgeBaseArn *string `field:"required" json:"bedrockKnowledgeBaseArn" yaml:"bedrockKnowledgeBaseArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentintentknowledgebaseconfiguration.html#cfn-lex-bot-bedrockagentintentknowledgebaseconfiguration-bedrockmodelconfiguration
	//
	BedrockModelConfiguration interface{} `field:"required" json:"bedrockModelConfiguration" yaml:"bedrockModelConfiguration"`
}

