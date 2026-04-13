package interfacesawssagemaker


// A reference to a Model resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelReference := &ModelReference{
//   	ModelArn: jsii.String("modelArn"),
//   }
//
type ModelReference struct {
	// The ModelArn of the Model resource.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
}

