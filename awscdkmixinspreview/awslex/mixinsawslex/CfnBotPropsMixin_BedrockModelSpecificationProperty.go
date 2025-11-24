package mixinsawslex


// Contains information about the Amazon Bedrock model used to interpret the prompt used in descriptive bot building.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockModelSpecificationProperty := &BedrockModelSpecificationProperty{
//   	BedrockGuardrailConfiguration: &BedrockGuardrailConfigurationProperty{
//   		BedrockGuardrailIdentifier: jsii.String("bedrockGuardrailIdentifier"),
//   		BedrockGuardrailVersion: jsii.String("bedrockGuardrailVersion"),
//   	},
//   	BedrockModelCustomPrompt: jsii.String("bedrockModelCustomPrompt"),
//   	BedrockTraceStatus: jsii.String("bedrockTraceStatus"),
//   	ModelArn: jsii.String("modelArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockmodelspecification.html
//
type CfnBotPropsMixin_BedrockModelSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockmodelspecification.html#cfn-lex-bot-bedrockmodelspecification-bedrockguardrailconfiguration
	//
	BedrockGuardrailConfiguration interface{} `field:"optional" json:"bedrockGuardrailConfiguration" yaml:"bedrockGuardrailConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockmodelspecification.html#cfn-lex-bot-bedrockmodelspecification-bedrockmodelcustomprompt
	//
	BedrockModelCustomPrompt *string `field:"optional" json:"bedrockModelCustomPrompt" yaml:"bedrockModelCustomPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockmodelspecification.html#cfn-lex-bot-bedrockmodelspecification-bedrocktracestatus
	//
	BedrockTraceStatus *string `field:"optional" json:"bedrockTraceStatus" yaml:"bedrockTraceStatus"`
	// The ARN of the foundation model used in descriptive bot building.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-bedrockmodelspecification.html#cfn-lex-bot-bedrockmodelspecification-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
}

