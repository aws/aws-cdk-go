package interfacesawsappconfig


// A reference to a DeploymentStrategy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentStrategyReference := &DeploymentStrategyReference{
//   	DeploymentStrategyId: jsii.String("deploymentStrategyId"),
//   }
//
type DeploymentStrategyReference struct {
	// The Id of the DeploymentStrategy resource.
	DeploymentStrategyId *string `field:"required" json:"deploymentStrategyId" yaml:"deploymentStrategyId"`
}

