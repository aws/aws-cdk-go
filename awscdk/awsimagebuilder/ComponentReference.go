package awsimagebuilder


// A reference to a Component resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentReference := &ComponentReference{
//   	ComponentArn: jsii.String("componentArn"),
//   }
//
type ComponentReference struct {
	// The Arn of the Component resource.
	ComponentArn *string `field:"required" json:"componentArn" yaml:"componentArn"`
}

