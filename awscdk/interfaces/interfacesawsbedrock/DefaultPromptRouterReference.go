package interfacesawsbedrock


// A reference to a DefaultPromptRouter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultPromptRouterReference := &DefaultPromptRouterReference{
//   	PromptRouterArn: jsii.String("promptRouterArn"),
//   }
//
type DefaultPromptRouterReference struct {
	// The PromptRouterArn of the DefaultPromptRouter resource.
	PromptRouterArn *string `field:"required" json:"promptRouterArn" yaml:"promptRouterArn"`
}

