package awscodedeploy


// Properties of a reference to a CodeDeploy ECS Deployment Group.
//
// Example:
//   var application ecsApplication
//
//   deploymentGroup := codedeploy.EcsDeploymentGroup_FromEcsDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &EcsDeploymentGroupAttributes{
//   	Application: Application,
//   	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
//   })
//
// See: EcsDeploymentGroup#fromEcsDeploymentGroupAttributes.
//
type EcsDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	Application IEcsApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Default: EcsDeploymentConfig.ALL_AT_ONCE
	//
	DeploymentConfig IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

