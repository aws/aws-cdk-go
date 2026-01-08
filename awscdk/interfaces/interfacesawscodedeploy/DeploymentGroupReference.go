package interfacesawscodedeploy


// A reference to a DeploymentGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentGroupReference := &DeploymentGroupReference{
//   	DeploymentGroupName: jsii.String("deploymentGroupName"),
//   }
//
type DeploymentGroupReference struct {
	// The DeploymentGroupName of the DeploymentGroup resource.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
}

