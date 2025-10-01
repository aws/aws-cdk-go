package awsbedrock


// A reference to a IntelligentPromptRouter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intelligentPromptRouterReference := &IntelligentPromptRouterReference{
//   	PromptRouterArn: jsii.String("promptRouterArn"),
//   }
//
type IntelligentPromptRouterReference struct {
	// The PromptRouterArn of the IntelligentPromptRouter resource.
	PromptRouterArn *string `field:"required" json:"promptRouterArn" yaml:"promptRouterArn"`
}

