package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
)

// Properties of a reference to a CodeDeploy ECS Deployment Group.
//
// Example:
//   var application EcsApplication
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
	Application interfacesawscodedeploy.IApplicationRef `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy ECS Deployment Group that we are referencing.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Default: EcsDeploymentConfig.ALL_AT_ONCE
	//
	DeploymentConfig interfacesawscodedeploy.IDeploymentConfigRef `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

