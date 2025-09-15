package awsworkspacesthinclient


// A reference to a Environment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentReference := &EnvironmentReference{
//   	EnvironmentArn: jsii.String("environmentArn"),
//   	EnvironmentId: jsii.String("environmentId"),
//   }
//
type EnvironmentReference struct {
	// The ARN of the Environment resource.
	EnvironmentArn *string `field:"required" json:"environmentArn" yaml:"environmentArn"`
	// The Id of the Environment resource.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}

