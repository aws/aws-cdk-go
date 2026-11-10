package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   handoffAgentConfigurationProperty := &HandoffAgentConfigurationProperty{
//   	AgentTarget: &AgentTargetProperty{
//   		AiAgentId: jsii.String("aiAgentId"),
//   		ApplicationId: jsii.String("applicationId"),
//   	},
//   	AudioStreamingEnabled: jsii.Boolean(false),
//   	ImmediateHandoff: jsii.Boolean(false),
//   	Instruction: &MultiAgentInstructionProperty{
//   		Examples: []*string{
//   			jsii.String("examples"),
//   		},
//   		Instruction: jsii.String("instruction"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-handoffagentconfiguration.html
//
type CfnAIAgentPropsMixin_HandoffAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-handoffagentconfiguration.html#cfn-wisdom-aiagent-handoffagentconfiguration-agenttarget
	//
	AgentTarget interface{} `field:"optional" json:"agentTarget" yaml:"agentTarget"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-handoffagentconfiguration.html#cfn-wisdom-aiagent-handoffagentconfiguration-audiostreamingenabled
	//
	AudioStreamingEnabled interface{} `field:"optional" json:"audioStreamingEnabled" yaml:"audioStreamingEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-handoffagentconfiguration.html#cfn-wisdom-aiagent-handoffagentconfiguration-immediatehandoff
	//
	ImmediateHandoff interface{} `field:"optional" json:"immediateHandoff" yaml:"immediateHandoff"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-handoffagentconfiguration.html#cfn-wisdom-aiagent-handoffagentconfiguration-instruction
	//
	Instruction interface{} `field:"optional" json:"instruction" yaml:"instruction"`
}

