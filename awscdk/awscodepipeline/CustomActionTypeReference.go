package awscodepipeline


// A reference to a CustomActionType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionTypeReference := &CustomActionTypeReference{
//   	Category: jsii.String("category"),
//   	Provider: jsii.String("provider"),
//   	Version: jsii.String("version"),
//   }
//
type CustomActionTypeReference struct {
	// The Category of the CustomActionType resource.
	Category *string `field:"required" json:"category" yaml:"category"`
	// The Provider of the CustomActionType resource.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The Version of the CustomActionType resource.
	Version *string `field:"required" json:"version" yaml:"version"`
}

