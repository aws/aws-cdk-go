package awswisdom


// Properties for defining a `CfnAIGuardrailVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAIGuardrailVersionProps := &CfnAIGuardrailVersionProps{
//   	AiGuardrailId: jsii.String("aiGuardrailId"),
//   	AssistantId: jsii.String("assistantId"),
//
//   	// the properties below are optional
//   	ModifiedTimeSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrailversion.html
//
type CfnAIGuardrailVersionProps struct {
	// The ID of the AI guardrail version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrailversion.html#cfn-wisdom-aiguardrailversion-aiguardrailid
	//
	AiGuardrailId *string `field:"required" json:"aiGuardrailId" yaml:"aiGuardrailId"`
	// The ID of the AI guardrail version assistant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrailversion.html#cfn-wisdom-aiguardrailversion-assistantid
	//
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The modified time of the AI guardrail version in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiguardrailversion.html#cfn-wisdom-aiguardrailversion-modifiedtimeseconds
	//
	ModifiedTimeSeconds *float64 `field:"optional" json:"modifiedTimeSeconds" yaml:"modifiedTimeSeconds"`
}

