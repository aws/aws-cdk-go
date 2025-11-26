package previewawslexmixins


// Contains specifications about the Amazon Lex runtime generative AI capabilities from Amazon Bedrock that you can turn on for your bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   runtimeSettingsProperty := &RuntimeSettingsProperty{
//   	NluImprovementSpecification: &NluImprovementSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	SlotResolutionImprovementSpecification: &SlotResolutionImprovementSpecificationProperty{
//   		BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   			BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   				BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   				BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   			},
//   			BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   			BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   			ModelArn: jsii.String("modelArn"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-runtimesettings.html
//
type CfnBotPropsMixin_RuntimeSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-runtimesettings.html#cfn-lex-bot-runtimesettings-nluimprovementspecification
	//
	NluImprovementSpecification interface{} `field:"optional" json:"nluImprovementSpecification" yaml:"nluImprovementSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-runtimesettings.html#cfn-lex-bot-runtimesettings-slotresolutionimprovementspecification
	//
	SlotResolutionImprovementSpecification interface{} `field:"optional" json:"slotResolutionImprovementSpecification" yaml:"slotResolutionImprovementSpecification"`
}

