package interfacesawsrefactorspaces


// A reference to a Environment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentReference := &EnvironmentReference{
//   	EnvironmentArn: jsii.String("environmentArn"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   }
//
type EnvironmentReference struct {
	// The ARN of the Environment resource.
	EnvironmentArn *string `field:"required" json:"environmentArn" yaml:"environmentArn"`
	// The EnvironmentIdentifier of the Environment resource.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
}

