package mixinsawslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockAgentConfigurationProperty := &BedrockAgentConfigurationProperty{
//   	BedrockAgentAliasId: jsii.String("bedrockAgentAliasId"),
//   	BedrockAgentId: jsii.String("bedrockAgentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentconfiguration.html
//
type CfnBotPropsMixin_BedrockAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentconfiguration.html#cfn-lex-bot-bedrockagentconfiguration-bedrockagentaliasid
	//
	BedrockAgentAliasId *string `field:"optional" json:"bedrockAgentAliasId" yaml:"bedrockAgentAliasId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockagentconfiguration.html#cfn-lex-bot-bedrockagentconfiguration-bedrockagentid
	//
	BedrockAgentId *string `field:"optional" json:"bedrockAgentId" yaml:"bedrockAgentId"`
}

