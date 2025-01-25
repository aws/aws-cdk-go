package awsbedrock


// Tool input schema.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolinputschema.html#cfn-bedrock-promptversion-toolinputschema-json
	//
	Json interface{} `field:"required" json:"json" yaml:"json"`
}

