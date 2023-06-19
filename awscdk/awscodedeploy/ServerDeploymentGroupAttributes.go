package awscodedeploy


// Properties of a reference to a CodeDeploy EC2/on-premise Deployment Group.
//
// Example:
//   var application serverApplication
//
//   deploymentGroup := codedeploy.ServerDeploymentGroup_FromServerDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &ServerDeploymentGroupAttributes{
//   	Application: Application,
//   	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
//   })
//
// See: ServerDeploymentGroup#import.
//
// Experimental.
type ServerDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy EC2/on-premise Application that this Deployment Group belongs to.
	// Experimental.
	Application IServerApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy EC2/on-premise Deployment Group that we are referencing.
	// Experimental.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	// Experimental.
	DeploymentConfig IServerDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

