package previewawslexmixins


// Contains specifications about the Amazon Lex build time generative AI capabilities from Amazon Bedrock that you can turn on for your bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   buildtimeSettingsProperty := &BuildtimeSettingsProperty{
//   	DescriptiveBotBuilderSpecification: &DescriptiveBotBuilderSpecificationProperty{
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
//   	SampleUtteranceGenerationSpecification: &SampleUtteranceGenerationSpecificationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-buildtimesettings.html
//
type CfnBotPropsMixin_BuildtimeSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-buildtimesettings.html#cfn-lex-bot-buildtimesettings-descriptivebotbuilderspecification
	//
	DescriptiveBotBuilderSpecification interface{} `field:"optional" json:"descriptiveBotBuilderSpecification" yaml:"descriptiveBotBuilderSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-buildtimesettings.html#cfn-lex-bot-buildtimesettings-sampleutterancegenerationspecification
	//
	SampleUtteranceGenerationSpecification interface{} `field:"optional" json:"sampleUtteranceGenerationSpecification" yaml:"sampleUtteranceGenerationSpecification"`
}

