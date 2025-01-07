package awsbedrock


// Information about a tool that you can use with the Converse API.
//
// For more information, see [Call a tool with the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   toolProperty := &ToolProperty{
//   	ToolSpec: &ToolSpecificationProperty{
//   		InputSchema: &ToolInputSchemaProperty{
//   			Json: json,
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-tool.html
//
type CfnPrompt_ToolProperty struct {
	// The specfication for the tool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-tool.html#cfn-bedrock-prompt-tool-toolspec
	//
	ToolSpec interface{} `field:"required" json:"toolSpec" yaml:"toolSpec"`
}

