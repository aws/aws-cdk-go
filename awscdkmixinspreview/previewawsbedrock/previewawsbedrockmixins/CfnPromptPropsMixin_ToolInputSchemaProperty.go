package previewawsbedrockmixins


// The schema for the tool.
//
// The top level schema type must be `object` . For more information, see [Call a tool with the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var json interface{}
//
//   toolInputSchemaProperty := &ToolInputSchemaProperty{
//   	Json: json,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolinputschema.html
//
type CfnPromptPropsMixin_ToolInputSchemaProperty struct {
	// The JSON schema for the tool.
	//
	// For more information, see [JSON Schema Reference](https://docs.aws.amazon.com/https://json-schema.org/understanding-json-schema/reference) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolinputschema.html#cfn-bedrock-prompt-toolinputschema-json
	//
	Json interface{} `field:"optional" json:"json" yaml:"json"`
}

