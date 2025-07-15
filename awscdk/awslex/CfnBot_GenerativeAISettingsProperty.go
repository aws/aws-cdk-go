package awslex


// Contains specifications about the generative AI capabilities from Amazon Bedrock that you can turn on for your bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generativeAISettingsProperty := &GenerativeAISettingsProperty{
//   	BuildtimeSettings: &BuildtimeSettingsProperty{
//   		DescriptiveBotBuilderSpecification: &DescriptiveBotBuilderSpecificationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   				ModelArn: jsii.String("modelArn"),
//
//   				// the properties below are optional
//   				BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   					BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   					BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   				},
//   				BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   				BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   			},
//   		},
//   		SampleUtteranceGenerationSpecification: &SampleUtteranceGenerationSpecificationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   				ModelArn: jsii.String("modelArn"),
//
//   				// the properties below are optional
//   				BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   					BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   					BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   				},
//   				BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   				BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   			},
//   		},
//   	},
//   	RuntimeSettings: &RuntimeSettingsProperty{
//   		NluImprovementSpecification: &NluImprovementSpecificationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		SlotResolutionImprovementSpecification: &SlotResolutionImprovementSpecificationProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   				ModelArn: jsii.String("modelArn"),
//
//   				// the properties below are optional
//   				BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   					BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   					BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   				},
//   				BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   				BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-generativeaisettings.html
//
type CfnBot_GenerativeAISettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-generativeaisettings.html#cfn-lex-bot-generativeaisettings-buildtimesettings
	//
	BuildtimeSettings interface{} `field:"optional" json:"buildtimeSettings" yaml:"buildtimeSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-generativeaisettings.html#cfn-lex-bot-generativeaisettings-runtimesettings
	//
	RuntimeSettings interface{} `field:"optional" json:"runtimeSettings" yaml:"runtimeSettings"`
}

