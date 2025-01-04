package awsbedrock


// The schema for the tool.
//
// The top level schema type must be `object` .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolinputschema.html
//
type CfnPrompt_ToolInputSchemaProperty struct {
	// The JSON schema for the tool.
	//
	// For more information, see [JSON Schema Reference](https://docs.aws.amazon.com/https://json-schema.org/understanding-json-schema/reference) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolinputschema.html#cfn-bedrock-prompt-toolinputschema-json
	//
	Json interface{} `field:"required" json:"json" yaml:"json"`
}

