package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   toolInstructionProperty := &ToolInstructionProperty{
//   	Examples: []*string{
//   		jsii.String("examples"),
//   	},
//   	Instruction: jsii.String("instruction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolinstruction.html
//
type CfnAIAgent_ToolInstructionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolinstruction.html#cfn-wisdom-aiagent-toolinstruction-examples
	//
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolinstruction.html#cfn-wisdom-aiagent-toolinstruction-instruction
	//
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
}

