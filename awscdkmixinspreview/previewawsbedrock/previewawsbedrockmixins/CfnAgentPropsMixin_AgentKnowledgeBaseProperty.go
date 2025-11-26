package previewawsbedrockmixins


// Contains details about a knowledge base that is associated with an agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   agentKnowledgeBaseProperty := &AgentKnowledgeBaseProperty{
//   	Description: jsii.String("description"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	KnowledgeBaseState: jsii.String("knowledgeBaseState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html
//
type CfnAgentPropsMixin_AgentKnowledgeBaseProperty struct {
	// The description of the association between the agent and the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html#cfn-bedrock-agent-agentknowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the association between the agent and the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html#cfn-bedrock-agent-agentknowledgebase-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Specifies whether to use the knowledge base or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentknowledgebase.html#cfn-bedrock-agent-agentknowledgebase-knowledgebasestate
	//
	KnowledgeBaseState *string `field:"optional" json:"knowledgeBaseState" yaml:"knowledgeBaseState"`
}

