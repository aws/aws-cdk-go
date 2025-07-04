package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeSettingsProperty := &RuntimeSettingsProperty{
//   	NluImprovementSpecification: &NluImprovementSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	SlotResolutionImprovementSpecification: &SlotResolutionImprovementSpecificationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-runtimesettings.html
//
type CfnBot_RuntimeSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-runtimesettings.html#cfn-lex-bot-runtimesettings-nluimprovementspecification
	//
	NluImprovementSpecification interface{} `field:"optional" json:"nluImprovementSpecification" yaml:"nluImprovementSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-runtimesettings.html#cfn-lex-bot-runtimesettings-slotresolutionimprovementspecification
	//
	SlotResolutionImprovementSpecification interface{} `field:"optional" json:"slotResolutionImprovementSpecification" yaml:"slotResolutionImprovementSpecification"`
}

