package previewawslexmixins


// Contains specifications for the sample utterance generation feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sampleUtteranceGenerationSpecificationProperty := &SampleUtteranceGenerationSpecificationProperty{
//   	BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   		BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   			BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   			BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   		},
//   		BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   		BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   		ModelArn: jsii.String("modelArn"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sampleutterancegenerationspecification.html
//
type CfnBotPropsMixin_SampleUtteranceGenerationSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sampleutterancegenerationspecification.html#cfn-lex-bot-sampleutterancegenerationspecification-bedrockmodelspecification
	//
	BedrockModelSpecification interface{} `field:"optional" json:"bedrockModelSpecification" yaml:"bedrockModelSpecification"`
	// Specifies whether to enable sample utterance generation or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sampleutterancegenerationspecification.html#cfn-lex-bot-sampleutterancegenerationspecification-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

