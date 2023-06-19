package awscodedeploy


// Properties of a reference to a CodeDeploy ECS Deployment Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecsApplication ecsApplication
//   var ecsDeploymentConfig iEcsDeploymentConfig
//
//   ecsDeploymentGroupAttributes := &EcsDeploymentGroupAttributes{
//   	Application: ecsApplication,
//   	DeploymentGroupName: jsii.String("deploymentGroupName"),
//
//   	// the properties below are optional
//   	DeploymentConfig: ecsDeploymentConfig,
//   }
//
// See: EcsDeploymentGroup#fromEcsDeploymentGroupAttributes.
//
// Experimental.
type EcsDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	// Experimental.
	Application IEcsApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

