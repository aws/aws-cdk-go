package interfacesawscodedeploy


// A reference to a DeploymentConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigReference := &DeploymentConfigReference{
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   }
//
type DeploymentConfigReference struct {
	// The DeploymentConfigName of the DeploymentConfig resource.
	DeploymentConfigName *string `field:"required" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

