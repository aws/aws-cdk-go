package awsbedrock


// The schema for the tool.
//
// The top level schema type must be `object` . For more information, see [Call a tool with the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   toolInputSchemaProperty := &ToolInputSchemaProperty{
//   	Json: json,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolinputschema.html
//
type CfnPromptVersion_ToolInputSchemaProperty struct {
	// The JSON schema for the tool.
	//
	// For more information, see [JSON Schema Reference](https://docs.aws.amazon.com/https://json-schema.org/understanding-json-schema/reference) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolinputschema.html#cfn-bedrock-promptversion-toolinputschema-json
	//
	Json interface{} `field:"required" json:"json" yaml:"json"`
}

