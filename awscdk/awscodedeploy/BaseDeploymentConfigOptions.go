package awscodedeploy


// Construction properties of `BaseDeploymentConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseDeploymentConfigOptions := &BaseDeploymentConfigOptions{
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   }
//
type BaseDeploymentConfigOptions struct {
	// The physical, human-readable name of the Deployment Configuration.
	// Default: - automatically generated name.
	//
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

