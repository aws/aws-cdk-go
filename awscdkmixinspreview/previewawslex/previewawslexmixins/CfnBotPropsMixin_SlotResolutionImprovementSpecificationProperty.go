package previewawslexmixins


// Contains specifications for the assisted slot resolution feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slotResolutionImprovementSpecificationProperty := &SlotResolutionImprovementSpecificationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotresolutionimprovementspecification.html
//
type CfnBotPropsMixin_SlotResolutionImprovementSpecificationProperty struct {
	// An object containing information about the Amazon Bedrock model used to assist slot resolution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotresolutionimprovementspecification.html#cfn-lex-bot-slotresolutionimprovementspecification-bedrockmodelspecification
	//
	BedrockModelSpecification interface{} `field:"optional" json:"bedrockModelSpecification" yaml:"bedrockModelSpecification"`
	// Specifies whether assisted slot resolution is turned on or off.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotresolutionimprovementspecification.html#cfn-lex-bot-slotresolutionimprovementspecification-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

