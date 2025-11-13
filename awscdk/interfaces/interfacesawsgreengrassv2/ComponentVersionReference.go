package interfacesawsgreengrassv2


// A reference to a ComponentVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentVersionReference := &ComponentVersionReference{
//   	ComponentVersionArn: jsii.String("componentVersionArn"),
//   }
//
type ComponentVersionReference struct {
	// The Arn of the ComponentVersion resource.
	ComponentVersionArn *string `field:"required" json:"componentVersionArn" yaml:"componentVersionArn"`
}

