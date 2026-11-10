package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiAgentInstructionProperty := &MultiAgentInstructionProperty{
//   	Examples: []*string{
//   		jsii.String("examples"),
//   	},
//   	Instruction: jsii.String("instruction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-multiagentinstruction.html
//
type CfnAIAgent_MultiAgentInstructionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-multiagentinstruction.html#cfn-wisdom-aiagent-multiagentinstruction-examples
	//
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-multiagentinstruction.html#cfn-wisdom-aiagent-multiagentinstruction-instruction
	//
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
}

