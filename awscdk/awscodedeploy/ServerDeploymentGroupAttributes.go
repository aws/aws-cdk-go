package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
)

// Properties of a reference to a CodeDeploy EC2/on-premise Deployment Group.
//
// Example:
//   var application ServerApplication
//
//   deploymentGroup := codedeploy.ServerDeploymentGroup_FromServerDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &ServerDeploymentGroupAttributes{
//   	Application: Application,
//   	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
//   })
//
// See: ServerDeploymentGroup# import.
//
type ServerDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy EC2/on-premise Application that this Deployment Group belongs to.
	Application interfacesawscodedeploy.IApplicationRef `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy EC2/on-premise Deployment Group that we are referencing.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Default: ServerDeploymentConfig#OneAtATime.
	//
	DeploymentConfig interfacesawscodedeploy.IDeploymentConfigRef `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

