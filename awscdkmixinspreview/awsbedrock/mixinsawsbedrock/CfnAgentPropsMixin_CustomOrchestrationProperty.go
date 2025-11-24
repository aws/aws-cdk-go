package mixinsawsbedrock


// Contains details of the custom orchestration configured for the agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customOrchestrationProperty := &CustomOrchestrationProperty{
//   	Executor: &OrchestrationExecutorProperty{
//   		Lambda: jsii.String("lambda"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-customorchestration.html
//
type CfnAgentPropsMixin_CustomOrchestrationProperty struct {
	// The structure of the executor invoking the actions in custom orchestration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-customorchestration.html#cfn-bedrock-agent-customorchestration-executor
	//
	Executor interface{} `field:"optional" json:"executor" yaml:"executor"`
}

