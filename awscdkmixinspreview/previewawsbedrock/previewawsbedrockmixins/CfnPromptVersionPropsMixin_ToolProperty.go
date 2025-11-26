package previewawsbedrockmixins


// Information about a tool that you can use with the Converse API.
//
// For more information, see [Call a tool with the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var json interface{}
//
//   toolProperty := &ToolProperty{
//   	CachePoint: &CachePointBlockProperty{
//   		Type: jsii.String("type"),
//   	},
//   	ToolSpec: &ToolSpecificationProperty{
//   		Description: jsii.String("description"),
//   		InputSchema: &ToolInputSchemaProperty{
//   			Json: json,
//   		},
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-tool.html
//
type CfnPromptVersionPropsMixin_ToolProperty struct {
	// CachePoint to include in the tool configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-tool.html#cfn-bedrock-promptversion-tool-cachepoint
	//
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// The specfication for the tool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-tool.html#cfn-bedrock-promptversion-tool-toolspec
	//
	ToolSpec interface{} `field:"optional" json:"toolSpec" yaml:"toolSpec"`
}

