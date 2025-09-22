package awsappconfig


// A reference to a Environment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentReference := &EnvironmentReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	EnvironmentId: jsii.String("environmentId"),
//   }
//
type EnvironmentReference struct {
	// The ApplicationId of the Environment resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The EnvironmentId of the Environment resource.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}

