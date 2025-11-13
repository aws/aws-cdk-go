package interfacesawsrefactorspaces


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReference := &ApplicationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   }
//
type ApplicationReference struct {
	// The ARN of the Application resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The ApplicationIdentifier of the Application resource.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The EnvironmentIdentifier of the Application resource.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
}

