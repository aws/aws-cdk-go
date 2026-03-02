package interfacesawscodedeploy


// A reference to a DeploymentGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentGroupReference := &DeploymentGroupReference{
//   	ApplicationName: jsii.String("applicationName"),
//   	DeploymentGroupName: jsii.String("deploymentGroupName"),
//   }
//
type DeploymentGroupReference struct {
	// The ApplicationName of the DeploymentGroup resource.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The DeploymentGroupName of the DeploymentGroup resource.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
}

