package awsbedrock


// Configuration for chat prompt template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentBlockProperty := &ContentBlockProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-contentblock.html
//
type CfnPromptVersion_ContentBlockProperty struct {
	// Configuration for chat prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-contentblock.html#cfn-bedrock-promptversion-contentblock-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}

