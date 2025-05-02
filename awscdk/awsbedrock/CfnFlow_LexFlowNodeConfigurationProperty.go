package awsbedrock


// Contains configurations for a Lex node in the flow.
//
// You specify a Amazon Lex bot to invoke. This node takes an utterance as the input and returns as the output the intent identified by the Amazon Lex bot. For more information, see [Node types in Amazon Bedrock works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lexFlowNodeConfigurationProperty := &LexFlowNodeConfigurationProperty{
//   	BotAliasArn: jsii.String("botAliasArn"),
//   	LocaleId: jsii.String("localeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-lexflownodeconfiguration.html
//
type CfnFlow_LexFlowNodeConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Lex bot alias to invoke.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-lexflownodeconfiguration.html#cfn-bedrock-flow-lexflownodeconfiguration-botaliasarn
	//
	BotAliasArn *string `field:"required" json:"botAliasArn" yaml:"botAliasArn"`
	// The Region to invoke the Amazon Lex bot in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-lexflownodeconfiguration.html#cfn-bedrock-flow-lexflownodeconfiguration-localeid
	//
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

