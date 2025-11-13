package interfacesawsm2


// A reference to a Deployment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReference := &DeploymentReference{
//   	ApplicationId: jsii.String("applicationId"),
//   }
//
type DeploymentReference struct {
	// The ApplicationId of the Deployment resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
}

