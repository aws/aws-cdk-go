package awsbedrock


// Tool specification.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolspecification.html
//
type CfnPromptVersion_ToolSpecificationProperty struct {
	// Tool input schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolspecification.html#cfn-bedrock-promptversion-toolspecification-inputschema
	//
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Tool name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolspecification.html#cfn-bedrock-promptversion-toolspecification-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolspecification.html#cfn-bedrock-promptversion-toolspecification-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

