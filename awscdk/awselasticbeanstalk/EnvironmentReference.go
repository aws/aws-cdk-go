package awselasticbeanstalk


// A reference to a Environment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentReference := &EnvironmentReference{
//   	EnvironmentName: jsii.String("environmentName"),
//   }
//
type EnvironmentReference struct {
	// The EnvironmentName of the Environment resource.
	EnvironmentName *string `field:"required" json:"environmentName" yaml:"environmentName"`
}

