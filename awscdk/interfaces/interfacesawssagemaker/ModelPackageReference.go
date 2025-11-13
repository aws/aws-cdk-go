package interfacesawssagemaker


// A reference to a ModelPackage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageReference := &ModelPackageReference{
//   	ModelPackageArn: jsii.String("modelPackageArn"),
//   }
//
type ModelPackageReference struct {
	// The ModelPackageArn of the ModelPackage resource.
	ModelPackageArn *string `field:"required" json:"modelPackageArn" yaml:"modelPackageArn"`
}

