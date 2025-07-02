package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buildtimeSettingsProperty := &BuildtimeSettingsProperty{
//   	DescriptiveBotBuilderSpecification: &DescriptiveBotBuilderSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   			ModelArn: jsii.String("modelArn"),
//
//   			// the properties below are optional
//   			BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   				BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   				BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   			},
//   			BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   			BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   		},
//   	},
//   	SampleUtteranceGenerationSpecification: &SampleUtteranceGenerationSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   			ModelArn: jsii.String("modelArn"),
//
//   			// the properties below are optional
//   			BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   				BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   				BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   			},
//   			BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   			BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-buildtimesettings.html
//
type CfnBot_BuildtimeSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-buildtimesettings.html#cfn-lex-bot-buildtimesettings-descriptivebotbuilderspecification
	//
	DescriptiveBotBuilderSpecification interface{} `field:"optional" json:"descriptiveBotBuilderSpecification" yaml:"descriptiveBotBuilderSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-buildtimesettings.html#cfn-lex-bot-buildtimesettings-sampleutterancegenerationspecification
	//
	SampleUtteranceGenerationSpecification interface{} `field:"optional" json:"sampleUtteranceGenerationSpecification" yaml:"sampleUtteranceGenerationSpecification"`
}

