package awsbedrock


// The specification for the tool.
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
//   toolSpecificationProperty := &ToolSpecificationProperty{
//   	InputSchema: &ToolInputSchemaProperty{
//   		Json: json,
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolspecification.html
//
type CfnPrompt_ToolSpecificationProperty struct {
	// The input schema for the tool in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolspecification.html#cfn-bedrock-prompt-toolspecification-inputschema
	//
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The name for the tool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolspecification.html#cfn-bedrock-prompt-toolspecification-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description for the tool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolspecification.html#cfn-bedrock-prompt-toolspecification-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

