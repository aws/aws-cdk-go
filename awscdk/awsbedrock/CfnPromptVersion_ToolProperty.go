package awsbedrock


// Tool details.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-tool.html
//
type CfnPromptVersion_ToolProperty struct {
	// Tool specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-tool.html#cfn-bedrock-promptversion-tool-toolspec
	//
	ToolSpec interface{} `field:"required" json:"toolSpec" yaml:"toolSpec"`
}

