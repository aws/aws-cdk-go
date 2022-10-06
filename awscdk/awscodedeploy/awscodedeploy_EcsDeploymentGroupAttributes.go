package awscodedeploy


// Properties of a reference to a CodeDeploy ECS Deployment Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecsApplication ecsApplication
//   var ecsDeploymentConfig ecsDeploymentConfig
//
//   ecsDeploymentGroupAttributes := &ecsDeploymentGroupAttributes{
//   	application: ecsApplication,
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//
//   	// the properties below are optional
//   	deploymentConfig: ecsDeploymentConfig,
//   }
//
// See: EcsDeploymentGroup#fromEcsDeploymentGroupAttributes.
//
type EcsDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	Application IEcsApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	DeploymentConfig IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

