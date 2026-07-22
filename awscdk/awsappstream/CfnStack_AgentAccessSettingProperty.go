package awsappstream


// A permission setting for an agent action.
//
// Each setting specifies an agent action and whether it is enabled or disabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentAccessSettingProperty := &AgentAccessSettingProperty{
//   	AgentAction: jsii.String("agentAction"),
//   	Permission: jsii.String("permission"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccesssetting.html
//
type CfnStack_AgentAccessSettingProperty struct {
	// The agent action to configure.
	//
	// Valid values are COMPUTER_VISION, COMPUTER_INPUT, and FORWARD_MCP_TOOLS. COMPUTER_VISION allows agents to take screenshots of the desktop. COMPUTER_INPUT allows agents to click, type, and scroll on the desktop and requires COMPUTER_VISION to also be enabled. FORWARD_MCP_TOOLS allows agents to interact with applications and the desktop operating system through direct MCP calls rather than using computer use tools. Forwards MCP tools configured on the WorkSpaces application session to the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccesssetting.html#cfn-appstream-stack-agentaccesssetting-agentaction
	//
	AgentAction *string `field:"required" json:"agentAction" yaml:"agentAction"`
	// Whether the agent action is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-agentaccesssetting.html#cfn-appstream-stack-agentaccesssetting-permission
	//
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

