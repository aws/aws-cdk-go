package awsbedrock


// Details about topics for the guardrail to identify and deny.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicConfigProperty := &TopicConfigProperty{
//   	Definition: jsii.String("definition"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Examples: []*string{
//   		jsii.String("examples"),
//   	},
//   	InputAction: jsii.String("inputAction"),
//   	InputEnabled: jsii.Boolean(false),
//   	OutputAction: jsii.String("outputAction"),
//   	OutputEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html
//
type CfnGuardrail_TopicConfigProperty struct {
	// A definition of the topic to deny.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-definition
	//
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The name of the topic to deny.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies to deny the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of prompts, each of which is an example of a prompt that can be categorized as belonging to the topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-examples
	//
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-inputaction
	//
	InputAction *string `field:"optional" json:"inputAction" yaml:"inputAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-inputenabled
	//
	InputEnabled interface{} `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-outputaction
	//
	OutputAction *string `field:"optional" json:"outputAction" yaml:"outputAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-outputenabled
	//
	OutputEnabled interface{} `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}

