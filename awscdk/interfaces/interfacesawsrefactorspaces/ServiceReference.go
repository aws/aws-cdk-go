package interfacesawsrefactorspaces


// A reference to a Service resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceReference := &ServiceReference{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	ServiceArn: jsii.String("serviceArn"),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//   }
//
type ServiceReference struct {
	// The ApplicationIdentifier of the Service resource.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The EnvironmentIdentifier of the Service resource.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The ARN of the Service resource.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// The ServiceIdentifier of the Service resource.
	ServiceIdentifier *string `field:"required" json:"serviceIdentifier" yaml:"serviceIdentifier"`
}

