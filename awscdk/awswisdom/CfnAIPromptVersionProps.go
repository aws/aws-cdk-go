package awswisdom


// Properties for defining a `CfnAIPromptVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAIPromptVersionProps := &CfnAIPromptVersionProps{
//   	AiPromptId: jsii.String("aiPromptId"),
//   	AssistantId: jsii.String("assistantId"),
//
//   	// the properties below are optional
//   	ModifiedTimeSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aipromptversion.html
//
type CfnAIPromptVersionProps struct {
	// The identifier of the Amazon Q in Connect AI prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aipromptversion.html#cfn-wisdom-aipromptversion-aipromptid
	//
	AiPromptId *string `field:"required" json:"aiPromptId" yaml:"aiPromptId"`
	// The identifier of the Amazon Q in Connect assistant.
	//
	// Can be either the ID or the ARN. URLs cannot contain the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aipromptversion.html#cfn-wisdom-aipromptversion-assistantid
	//
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The time the AI Prompt version was last modified in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aipromptversion.html#cfn-wisdom-aipromptversion-modifiedtimeseconds
	//
	ModifiedTimeSeconds *float64 `field:"optional" json:"modifiedTimeSeconds" yaml:"modifiedTimeSeconds"`
}

