package awsbedrock


// Selective content guarding controls for enforced guardrails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectiveContentGuardingProperty := &SelectiveContentGuardingProperty{
//   	Messages: jsii.String("messages"),
//   	System: jsii.String("system"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-enforcedguardrailconfiguration-selectivecontentguarding.html
//
type CfnEnforcedGuardrailConfiguration_SelectiveContentGuardingProperty struct {
	// Selective guarding mode for user messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-enforcedguardrailconfiguration-selectivecontentguarding.html#cfn-bedrock-enforcedguardrailconfiguration-selectivecontentguarding-messages
	//
	Messages *string `field:"optional" json:"messages" yaml:"messages"`
	// Selective guarding mode for system prompts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-enforcedguardrailconfiguration-selectivecontentguarding.html#cfn-bedrock-enforcedguardrailconfiguration-selectivecontentguarding-system
	//
	System *string `field:"optional" json:"system" yaml:"system"`
}

