package previewawslexmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockAgentIntentConfigurationProperty := &BedrockAgentIntentConfigurationProperty{
//   	BedrockAgentConfiguration: &BedrockAgentConfigurationProperty{
//   		BedrockAgentAliasId: jsii.String("bedrockAgentAliasId"),
//   		BedrockAgentId: jsii.String("bedrockAgentId"),
//   	},
//   	BedrockAgentIntentKnowledgeBaseConfiguration: &BedrockAgentIntentKnowledgeBaseConfigurationProperty{
//   		BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
//   		BedrockModelConfiguration: &BedrockModelSpecificationProperty{
//   			BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   				BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   				BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   			},
//   			BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   			BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   			ModelArn: jsii.String("modelArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentintentconfiguration.html
//
type CfnBotPropsMixin_BedrockAgentIntentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentintentconfiguration.html#cfn-lex-bot-bedrockagentintentconfiguration-bedrockagentconfiguration
	//
	BedrockAgentConfiguration interface{} `field:"optional" json:"bedrockAgentConfiguration" yaml:"bedrockAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentintentconfiguration.html#cfn-lex-bot-bedrockagentintentconfiguration-bedrockagentintentknowledgebaseconfiguration
	//
	BedrockAgentIntentKnowledgeBaseConfiguration interface{} `field:"optional" json:"bedrockAgentIntentKnowledgeBaseConfiguration" yaml:"bedrockAgentIntentKnowledgeBaseConfiguration"`
}

