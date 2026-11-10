package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentTargetProperty := &AgentTargetProperty{
//   	AiAgentId: jsii.String("aiAgentId"),
//   	ApplicationId: jsii.String("applicationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-agenttarget.html
//
type CfnAIAgent_AgentTargetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-agenttarget.html#cfn-wisdom-aiagent-agenttarget-aiagentid
	//
	AiAgentId *string `field:"optional" json:"aiAgentId" yaml:"aiAgentId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-agenttarget.html#cfn-wisdom-aiagent-agenttarget-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
}

