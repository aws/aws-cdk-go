package previewawsbedrockmixins


// The model must request a specific tool.
//
// For example, `{"tool" : {"name" : "Your tool name"}}` . For more information, see [Call a tool with the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide
//
// > This field is only supported by Anthropic Claude 3 models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   specificToolChoiceProperty := &SpecificToolChoiceProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-specifictoolchoice.html
//
type CfnPromptVersionPropsMixin_SpecificToolChoiceProperty struct {
	// The name of the tool that the model must request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-specifictoolchoice.html#cfn-bedrock-promptversion-specifictoolchoice-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

