package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   delegateAgentConfigurationProperty := &DelegateAgentConfigurationProperty{
//   	AgentTarget: &AgentTargetProperty{
//   		AiAgentId: jsii.String("aiAgentId"),
//   		ApplicationId: jsii.String("applicationId"),
//   	},
//   	Instruction: &MultiAgentInstructionProperty{
//   		Examples: []*string{
//   			jsii.String("examples"),
//   		},
//   		Instruction: jsii.String("instruction"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-delegateagentconfiguration.html
//
type CfnAIAgentPropsMixin_DelegateAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-delegateagentconfiguration.html#cfn-wisdom-aiagent-delegateagentconfiguration-agenttarget
	//
	AgentTarget interface{} `field:"optional" json:"agentTarget" yaml:"agentTarget"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-delegateagentconfiguration.html#cfn-wisdom-aiagent-delegateagentconfiguration-instruction
	//
	Instruction interface{} `field:"optional" json:"instruction" yaml:"instruction"`
}

