package awssagemaker


// A reference to a ModelPackageGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageGroupReference := &ModelPackageGroupReference{
//   	ModelPackageGroupArn: jsii.String("modelPackageGroupArn"),
//   }
//
type ModelPackageGroupReference struct {
	// The ModelPackageGroupArn of the ModelPackageGroup resource.
	ModelPackageGroupArn *string `field:"required" json:"modelPackageGroupArn" yaml:"modelPackageGroupArn"`
}

