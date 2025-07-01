package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   descriptiveBotBuilderSpecificationProperty := &DescriptiveBotBuilderSpecificationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	BedrockModelSpecification: &BedrockModelSpecificationProperty{
//   		ModelArn: jsii.String("modelArn"),
//
//   		// the properties below are optional
//   		BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   			BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   			BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   		},
//   		BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   		BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-descriptivebotbuilderspecification.html
//
type CfnBot_DescriptiveBotBuilderSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-descriptivebotbuilderspecification.html#cfn-lex-bot-descriptivebotbuilderspecification-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-descriptivebotbuilderspecification.html#cfn-lex-bot-descriptivebotbuilderspecification-bedrockmodelspecification
	//
	BedrockModelSpecification interface{} `field:"optional" json:"bedrockModelSpecification" yaml:"bedrockModelSpecification"`
}

