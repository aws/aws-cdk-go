package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiAgentConfigurationProperty := &MultiAgentConfigurationProperty{
//   	DelegateAgentConfiguration: &DelegateAgentConfigurationProperty{
//   		AgentTarget: &AgentTargetProperty{
//   			AiAgentId: jsii.String("aiAgentId"),
//   			ApplicationId: jsii.String("applicationId"),
//   		},
//
//   		// the properties below are optional
//   		Instruction: &MultiAgentInstructionProperty{
//   			Examples: []*string{
//   				jsii.String("examples"),
//   			},
//   			Instruction: jsii.String("instruction"),
//   		},
//   	},
//   	HandoffAgentConfiguration: &HandoffAgentConfigurationProperty{
//   		AgentTarget: &AgentTargetProperty{
//   			AiAgentId: jsii.String("aiAgentId"),
//   			ApplicationId: jsii.String("applicationId"),
//   		},
//
//   		// the properties below are optional
//   		AudioStreamingEnabled: jsii.Boolean(false),
//   		ImmediateHandoff: jsii.Boolean(false),
//   		Instruction: &MultiAgentInstructionProperty{
//   			Examples: []*string{
//   				jsii.String("examples"),
//   			},
//   			Instruction: jsii.String("instruction"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-multiagentconfiguration.html
//
type CfnAIAgent_MultiAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-multiagentconfiguration.html#cfn-wisdom-aiagent-multiagentconfiguration-delegateagentconfiguration
	//
	DelegateAgentConfiguration interface{} `field:"optional" json:"delegateAgentConfiguration" yaml:"delegateAgentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-multiagentconfiguration.html#cfn-wisdom-aiagent-multiagentconfiguration-handoffagentconfiguration
	//
	HandoffAgentConfiguration interface{} `field:"optional" json:"handoffAgentConfiguration" yaml:"handoffAgentConfiguration"`
}

