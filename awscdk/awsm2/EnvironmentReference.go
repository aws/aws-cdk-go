package awsm2


// A reference to a Environment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentReference := &EnvironmentReference{
//   	EnvironmentArn: jsii.String("environmentArn"),
//   }
//
type EnvironmentReference struct {
	// The EnvironmentArn of the Environment resource.
	EnvironmentArn *string `field:"required" json:"environmentArn" yaml:"environmentArn"`
}

