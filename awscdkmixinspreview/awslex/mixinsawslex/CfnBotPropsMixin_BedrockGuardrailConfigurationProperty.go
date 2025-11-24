package mixinsawslex


// The details on the Bedrock guardrail configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockGuardrailConfigurationProperty := &BedrockGuardrailConfigurationProperty{
//   	BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   	BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockguardrailconfiguration.html
//
type CfnBotPropsMixin_BedrockGuardrailConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockguardrailconfiguration.html#cfn-lex-bot-bedrockguardrailconfiguration-bedrockguardrailidentifier
	//
	BedrockGuardrailIdentifier *string `field:"optional" json:"bedrockGuardrailIdentifier" yaml:"bedrockGuardrailIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockguardrailconfiguration.html#cfn-lex-bot-bedrockguardrailconfiguration-bedrockguardrailversion
	//
	BedrockGuardrailVersion *string `field:"optional" json:"bedrockGuardrailVersion" yaml:"bedrockGuardrailVersion"`
}

