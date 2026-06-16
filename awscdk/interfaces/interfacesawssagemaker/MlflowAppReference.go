package interfacesawssagemaker


// A reference to a MlflowApp resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mlflowAppReference := &MlflowAppReference{
//   	MlflowAppArn: jsii.String("mlflowAppArn"),
//   }
//
type MlflowAppReference struct {
	// The Arn of the MlflowApp resource.
	MlflowAppArn *string `field:"required" json:"mlflowAppArn" yaml:"mlflowAppArn"`
}

