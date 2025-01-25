package awsbedrock


// Specific Tool choice.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   specificToolChoiceProperty := &SpecificToolChoiceProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-specifictoolchoice.html
//
type CfnPromptVersion_SpecificToolChoiceProperty struct {
	// Tool name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-specifictoolchoice.html#cfn-bedrock-promptversion-specifictoolchoice-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

