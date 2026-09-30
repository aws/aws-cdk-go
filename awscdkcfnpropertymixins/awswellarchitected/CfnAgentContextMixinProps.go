package awswellarchitected


// Properties for CfnAgentContextPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var content interface{}
//
//   cfnAgentContextMixinProps := &CfnAgentContextMixinProps{
//   	Content: content,
//   	ContextType: jsii.String("contextType"),
//   	ProfileArn: jsii.String("profileArn"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentcontext.html
//
type CfnAgentContextMixinProps struct {
	// The free-form content of the Agent Context, supplied as an arbitrary JSON object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentcontext.html#cfn-wellarchitected-agentcontext-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// The type of the Agent Context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentcontext.html#cfn-wellarchitected-agentcontext-contexttype
	//
	ContextType *string `field:"optional" json:"contextType" yaml:"contextType"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile that owns this context.
	//
	// Pass `!Ref` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentcontext.html#cfn-wellarchitected-agentcontext-profilearn
	//
	ProfileArn *string `field:"optional" json:"profileArn" yaml:"profileArn"`
	// The title of the Agent Context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentcontext.html#cfn-wellarchitected-agentcontext-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

