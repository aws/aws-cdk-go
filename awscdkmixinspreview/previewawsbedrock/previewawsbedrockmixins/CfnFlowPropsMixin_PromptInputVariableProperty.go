package previewawsbedrockmixins


// Contains information about a variable in the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptInputVariableProperty := &PromptInputVariableProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinputvariable.html
//
type CfnFlowPropsMixin_PromptInputVariableProperty struct {
	// The name of the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinputvariable.html#cfn-bedrock-flow-promptinputvariable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

