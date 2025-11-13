package interfacesawsappconfig


// A reference to a Deployment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReference := &DeploymentReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	DeploymentNumber: jsii.String("deploymentNumber"),
//   	EnvironmentId: jsii.String("environmentId"),
//   }
//
type DeploymentReference struct {
	// The ApplicationId of the Deployment resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The DeploymentNumber of the Deployment resource.
	DeploymentNumber *string `field:"required" json:"deploymentNumber" yaml:"deploymentNumber"`
	// The EnvironmentId of the Deployment resource.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}

