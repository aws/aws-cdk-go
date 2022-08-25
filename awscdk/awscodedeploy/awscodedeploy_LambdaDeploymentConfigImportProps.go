package awscodedeploy


// Properties of a reference to a CodeDeploy Lambda Deployment Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaDeploymentConfigImportProps := &lambdaDeploymentConfigImportProps{
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   }
//
// See: LambdaDeploymentConfig#import.
//
type LambdaDeploymentConfigImportProps struct {
	// The physical, human-readable name of the custom CodeDeploy Lambda Deployment Configuration that we are referencing.
	DeploymentConfigName *string `field:"required" json:"deploymentConfigName" yaml:"deploymentConfigName"`
}

