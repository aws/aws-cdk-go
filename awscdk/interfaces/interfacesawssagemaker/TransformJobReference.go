package interfacesawssagemaker


// A reference to a TransformJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformJobReference := &TransformJobReference{
//   	TransformJobArn: jsii.String("transformJobArn"),
//   }
//
type TransformJobReference struct {
	// The TransformJobArn of the TransformJob resource.
	TransformJobArn *string `field:"required" json:"transformJobArn" yaml:"transformJobArn"`
}

