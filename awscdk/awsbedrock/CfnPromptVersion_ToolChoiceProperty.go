package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var any interface{}
//   var auto interface{}
//
//   toolChoiceProperty := &ToolChoiceProperty{
//   	Any: any,
//   	Auto: auto,
//   	Tool: &SpecificToolChoiceProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html
//
type CfnPromptVersion_ToolChoiceProperty struct {
	// Any Tool choice.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html#cfn-bedrock-promptversion-toolchoice-any
	//
	Any interface{} `field:"optional" json:"any" yaml:"any"`
	// Auto Tool choice.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html#cfn-bedrock-promptversion-toolchoice-auto
	//
	Auto interface{} `field:"optional" json:"auto" yaml:"auto"`
	// Specific Tool choice.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html#cfn-bedrock-promptversion-toolchoice-tool
	//
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
}

