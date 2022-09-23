package awscodedeploy


// Properties of a reference to a CodeDeploy Lambda Deployment Group.
//
// Example:
//   var application lambdaApplication
//
//   deploymentGroup := codedeploy.lambdaDeploymentGroup.fromLambdaDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &lambdaDeploymentGroupAttributes{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
//   })
//
// See: LambdaDeploymentGroup#fromLambdaDeploymentGroupAttributes.
//
type LambdaDeploymentGroupAttributes struct {
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	Application ILambdaApplication `field:"required" json:"application" yaml:"application"`
	// The physical, human-readable name of the CodeDeploy Lambda Deployment Group that we are referencing.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// The Deployment Configuration this Deployment Group uses.
	DeploymentConfig ILambdaDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
}

