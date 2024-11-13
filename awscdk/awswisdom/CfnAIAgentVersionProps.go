package awswisdom


// Properties for defining a `CfnAIAgentVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAIAgentVersionProps := &CfnAIAgentVersionProps{
//   	AiAgentId: jsii.String("aiAgentId"),
//   	AssistantId: jsii.String("assistantId"),
//
//   	// the properties below are optional
//   	ModifiedTimeSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html
//
type CfnAIAgentVersionProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html#cfn-wisdom-aiagentversion-aiagentid
	//
	AiAgentId *string `field:"required" json:"aiAgentId" yaml:"aiAgentId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html#cfn-wisdom-aiagentversion-assistantid
	//
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html#cfn-wisdom-aiagentversion-modifiedtimeseconds
	//
	ModifiedTimeSeconds *float64 `field:"optional" json:"modifiedTimeSeconds" yaml:"modifiedTimeSeconds"`
}

