package awsbedrock


// Determines which tools the model should request in a call to `Converse` or `ConverseStream` .
//
// For more information, see [Call a tool with the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
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
	// The model must request at least one tool (no text is generated).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html#cfn-bedrock-promptversion-toolchoice-any
	//
	Any interface{} `field:"optional" json:"any" yaml:"any"`
	// (Default).
	//
	// The Model automatically decides if a tool should be called or whether to generate text instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html#cfn-bedrock-promptversion-toolchoice-auto
	//
	Auto interface{} `field:"optional" json:"auto" yaml:"auto"`
	// The Model must request the specified tool.
	//
	// Only supported by Anthropic Claude 3 models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolchoice.html#cfn-bedrock-promptversion-toolchoice-tool
	//
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
}

