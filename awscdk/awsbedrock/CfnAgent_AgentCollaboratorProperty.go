package awsbedrock


// An agent collaborator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentCollaboratorProperty := &AgentCollaboratorProperty{
//   	AgentDescriptor: &AgentDescriptorProperty{
//   		AliasArn: jsii.String("aliasArn"),
//   	},
//   	CollaborationInstruction: jsii.String("collaborationInstruction"),
//   	CollaboratorName: jsii.String("collaboratorName"),
//
//   	// the properties below are optional
//   	RelayConversationHistory: jsii.String("relayConversationHistory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentcollaborator.html
//
type CfnAgent_AgentCollaboratorProperty struct {
	// The collaborator's agent descriptor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentcollaborator.html#cfn-bedrock-agent-agentcollaborator-agentdescriptor
	//
	AgentDescriptor interface{} `field:"required" json:"agentDescriptor" yaml:"agentDescriptor"`
	// The collaborator's instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentcollaborator.html#cfn-bedrock-agent-agentcollaborator-collaborationinstruction
	//
	CollaborationInstruction *string `field:"required" json:"collaborationInstruction" yaml:"collaborationInstruction"`
	// The collaborator's collaborator name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentcollaborator.html#cfn-bedrock-agent-agentcollaborator-collaboratorname
	//
	CollaboratorName *string `field:"required" json:"collaboratorName" yaml:"collaboratorName"`
	// The collaborator's relay conversation history.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentcollaborator.html#cfn-bedrock-agent-agentcollaborator-relayconversationhistory
	//
	RelayConversationHistory *string `field:"optional" json:"relayConversationHistory" yaml:"relayConversationHistory"`
}

