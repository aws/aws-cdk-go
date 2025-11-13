package interfacesawsbedrock


// A reference to a PromptVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptVersionReference := &PromptVersionReference{
//   	PromptVersionArn: jsii.String("promptVersionArn"),
//   }
//
type PromptVersionReference struct {
	// The Arn of the PromptVersion resource.
	PromptVersionArn *string `field:"required" json:"promptVersionArn" yaml:"promptVersionArn"`
}

