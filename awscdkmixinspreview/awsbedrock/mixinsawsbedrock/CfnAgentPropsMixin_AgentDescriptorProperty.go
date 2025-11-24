package mixinsawsbedrock


// An agent descriptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   agentDescriptorProperty := &AgentDescriptorProperty{
//   	AliasArn: jsii.String("aliasArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentdescriptor.html
//
type CfnAgentPropsMixin_AgentDescriptorProperty struct {
	// The agent's alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentdescriptor.html#cfn-bedrock-agent-agentdescriptor-aliasarn
	//
	AliasArn *string `field:"optional" json:"aliasArn" yaml:"aliasArn"`
}

