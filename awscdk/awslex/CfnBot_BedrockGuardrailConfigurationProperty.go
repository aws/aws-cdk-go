package awslex


// The guardrail configuration in the Bedrock model specification details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockGuardrailConfigurationProperty := &BedrockGuardrailConfigurationProperty{
//   	BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   	BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockguardrailconfiguration.html
//
type CfnBot_BedrockGuardrailConfigurationProperty struct {
	// The unique guardrail id for the Bedrock guardrail configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockguardrailconfiguration.html#cfn-lex-bot-bedrockguardrailconfiguration-bedrockguardrailidentifier
	//
	BedrockGuardrailIdentifier *string `field:"optional" json:"bedrockGuardrailIdentifier" yaml:"bedrockGuardrailIdentifier"`
	// The guardrail version for the Bedrock guardrail configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockguardrailconfiguration.html#cfn-lex-bot-bedrockguardrailconfiguration-bedrockguardrailversion
	//
	BedrockGuardrailVersion *string `field:"optional" json:"bedrockGuardrailVersion" yaml:"bedrockGuardrailVersion"`
}

