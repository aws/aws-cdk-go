package previewawswisdommixins


// Properties for CfnAIAgentVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAIAgentVersionMixinProps := &CfnAIAgentVersionMixinProps{
//   	AiAgentId: jsii.String("aiAgentId"),
//   	AssistantId: jsii.String("assistantId"),
//   	ModifiedTimeSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html
//
type CfnAIAgentVersionMixinProps struct {
	// The identifier of the AI Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html#cfn-wisdom-aiagentversion-aiagentid
	//
	AiAgentId *string `field:"optional" json:"aiAgentId" yaml:"aiAgentId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html#cfn-wisdom-aiagentversion-assistantid
	//
	AssistantId *string `field:"optional" json:"assistantId" yaml:"assistantId"`
	// The time the AI Agent version was last modified in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aiagentversion.html#cfn-wisdom-aiagentversion-modifiedtimeseconds
	//
	ModifiedTimeSeconds *float64 `field:"optional" json:"modifiedTimeSeconds" yaml:"modifiedTimeSeconds"`
}

