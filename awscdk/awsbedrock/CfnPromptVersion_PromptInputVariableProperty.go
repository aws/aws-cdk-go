package awsbedrock


// Contains information about a variable in the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptInputVariableProperty := &PromptInputVariableProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptinputvariable.html
//
type CfnPromptVersion_PromptInputVariableProperty struct {
	// The name of the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptinputvariable.html#cfn-bedrock-promptversion-promptinputvariable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

